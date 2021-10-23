package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/user/model"
	"cpg-blog/internal/user/model/dao"
	"cpg-blog/internal/user/qo"
	"cpg-blog/internal/user/vo"
	jwt2 "cpg-blog/middleware/jwt"
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type Users struct{}

type tokenInfo struct {
	uid   int
	name  string
	email string
	root  int
}

//用户状态
const (
	one   int = iota + 1 //正常
	two                  //禁发文
	there                //冻结
)

// 生成token
func genToken(info tokenInfo) (token string, err error) {
	j := jwt2.NewJWT()

	// 构造用户claims信息(负荷)
	// 过期时间
	expiredTime := time.Now().Add(time.Duration(viper.GetInt("token.expires")) * time.Hour)
	claims := jwt2.CustomClaims{
		Uid:   strconv.Itoa(info.uid),
		Name:  info.name,
		Email: info.email,
		Root:  info.root,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),               // 过期时间
			IssuedAt:  time.Now().Unix(),                // 颁发时间
			Issuer:    viper.GetString("token.issuer"),  // 颁发者
			NotBefore: time.Now().Unix(),                // 生效时间
			Subject:   viper.GetString("token.subject"), // token主题
		},
	}
	token, err = j.CreateToken(claims)
	return token, err
}

//token添加进入redis
func addTokenToRedis(ctx *gin.Context, token string) (err error) {
	_, err = globalInit.RedisClient.SAdd(ctx, "token", token).Result()
	return err
}

//redis移除对应token
func removeTokenFromRedis(ctx *gin.Context, token string) (err error) {
	_, err = globalInit.RedisClient.SRem(ctx, "token", token).Result()
	return err
}

//加密
func (u Users) encryption(passwd string) (string, error) {
	store, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return string(store), err
	}
	return string(store), nil
}

//校验密码
func (u Users) comparePwd(storePasswd, passwd string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(storePasswd), []byte(passwd))
}

func (u *Users) Login(ctx *gin.Context) {
	loginQo := qo.LoginQO{}
	util.JsonConvert(ctx, &loginQo)
	users := new(dao.UserDAO).SelectByName(ctx, loginQo.Username)
	if 0 == len(users) {
		common.SendResponse(ctx, common.ErrUserNotFound, "")
		return
	}

	if u.comparePwd(users[0].Passwd, loginQo.Passwd) != nil {
		common.SendResponse(ctx, common.ErrPasswordIncorrect, "")
		return
	}
	//生成token
	tokenInfo := tokenInfo{
		int(users[0].UID),
		users[0].UserName,
		users[0].Email,
		users[0].IsRoot,
	}
	token, err := genToken(tokenInfo)
	if err != nil {
		common.SendResponse(ctx, common.ErrGenerateToken, err.Error())
		return
	}

	if ok := addTokenToRedis(ctx, token); ok != nil {
		common.SendResponse(ctx, common.ErrRedis, "")
		return
	}
	loginVo := vo.LoginVo{
		Token: token,
	}
	common.SendResponse(ctx, common.OK, loginVo)
}

func (u *Users) Logout(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	//清除redis对应token缓存
	err := removeTokenFromRedis(ctx, token)
	if err != nil {
		common.SendResponse(ctx, common.ErrRedis, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")
}

func (u Users) Register(ctx *gin.Context) {
	registerQO := qo.RegisterQO{}

	//校验必填请求参数
	util.JsonConvert(ctx, &registerQO)

	//校验唯一参数username、email
	users := new(dao.UserDAO).SelectByNameAndEmail(ctx, &model.User{UserName: registerQO.UserName, Email: registerQO.Email})
	if 0 < len(users) {
		common.SendResponse(ctx, common.ErrUserExisted, "")
		return
	} else {
		registerQO.State = cpgConst.ONE
		registerQO.IsRoot = cpgConst.ZERO
		storePasswd, err := u.encryption(registerQO.Passwd)
		if err != nil {
			common.SendResponse(ctx, common.ErrEncryption, err.Error())
			return
		}
		registerQO.Passwd = storePasswd
		user := model.User{}
		err = copier.Copy(&user, registerQO)
		err = new(dao.UserDAO).Create(ctx, &user)
		if err != nil {
			common.SendResponse(ctx, common.ErrDatabase, err.Error())
			return
		}

		user1 := new(dao.UserDAO).SelectByName(ctx, registerQO.UserName)
		//生成token
		tokenInfo := tokenInfo{
			int(user1[0].UID),
			user1[0].UserName,
			user1[0].Email,
			user1[0].IsRoot,
		}
		token, err := genToken(tokenInfo)

		if err != nil {
			common.SendResponse(ctx, common.ErrGenerateToken, err.Error())
			return
		}
		loginVo := vo.LoginVo{
			Token: token,
		}
		common.SendResponse(ctx, common.OK, loginVo)
	}

}

func (u Users) Info(ctx *gin.Context) {
	infoQO := qo.UserInfoQO{}
	util.JsonConvert(ctx, &infoQO)
	var user []model.User

	if infoQO.Email == "" && infoQO.Username == "" {
		common.SendResponse(ctx, common.ErrValidation, "")
		return
	} else if infoQO.Email == "" {
		user = new(dao.UserDAO).SelectByName(ctx, infoQO.Username)
	} else {
		user = new(dao.UserDAO).SelectByEmail(ctx, infoQO.Email)
	}
	common.SendResponse(ctx, common.OK, user[0])
}

func (u Users) List(ctx *gin.Context) {
	//TODO 分页
	var users []model.User
	listQO := qo.UserListQO{}
	util.JsonConvert(ctx, &listQO)
	if listQO.State != one && listQO.State != two && listQO.State != there {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}
	globalInit.Db.Where("state", listQO.State).Find(&users)
	var userList []vo.UserListVO
	copier.Copy(&userList, &users)
	common.SendResponse(ctx, common.OK, userList)
}

func (u Users) Modify(ctx *gin.Context) {
	modifyQO := qo.ModifyQO{}
	util.JsonConvert(ctx, &modifyQO)
	user := model.User{}
	if err := copier.Copy(&user, modifyQO); err != nil {
		common.SendResponse(ctx, common.ErrBind, err.Error())
		return
	}

	//查询数据库中该username或email存在的用户信息
	userList := new(dao.UserDAO).SelectByNameAndEmail(ctx, &user)

	if len(userList) > 1 || (len(userList) == 1 && userList[0].UID != modifyQO.UID) {
		common.SendResponse(ctx, common.ErrUserExisted, "")
		return
	}

	if err := new(dao.UserDAO).UpdateUserInfo(ctx, &user); err != nil {
		common.SendResponse(ctx, common.ErrDatabase, err.Error())
		return
	}
}
