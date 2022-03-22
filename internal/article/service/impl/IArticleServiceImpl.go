package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/article/model"
	"cpg-blog/internal/article/model/dao"
	"cpg-blog/internal/article/qo"
	"cpg-blog/internal/article/vo"
	"cpg-blog/internal/user/service"
	"cpg-blog/middleware/jwt"
	"cpg-blog/pkg/commonFunc/likeCommonFunc"
	"cpg-blog/pkg/commonFunc/userCommonFunc"
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"log"
	"reflect"
	"strconv"
)

type Article struct{}

//文章状态
const (
	//0-未审核
	unreviewed int = iota

	//1-已上线
	published

	//2-下线
	removed

	//3-用户删除
	deleted
)

var userService service.IUser

func tokenInfo(ctx *gin.Context) (Info *jwt.CustomClaims, err error) {
	Info, err = jwt.NewJWT().ParseToken(ctx.Request.Header.Get("token"))
	return
}

// Info 根据sn查询
func Info(ctx *gin.Context) (*gin.Context, error, interface{}) {
	infoQO := new(qo.ArticleInfoQO)
	util.JsonConvert(ctx, infoQO)
	article := new(model.Article)

	if err := copier.Copy(article, infoQO); err != nil {
		return ctx, common.ErrBind, err.Error()

	}
	article.Sn, _ = strconv.ParseInt(infoQO.Sn, 10, 64)
	article = new(dao.ArticleDAO).SelectBySn(ctx, article)

	if article.Aid == 0 {
		return ctx, common.ErrArticleNotExisted, ""

	}

	articleVO := vo.ArticleInfoVO{}
	if err := copier.Copy(&articleVO, article); err != nil {
		return ctx, common.ErrBind, err.Error()
	}

	articleVO.Sn = strconv.FormatInt(article.Sn, 10)
	//userMap := userService.FindUser(ctx, []int{article.Uid}, "", "")
	userMap := userCommonFunc.IUser(userCommonFunc.UserCommonFunc{}).FindUser(ctx, []int{article.Uid}, "", "")
	articleVO.Author = userMap[uint(article.Uid)].UserName
	articleVO.CreateAt = article.CreatedAt.Unix()
	articleVO.UpdatedAt = article.UpdatedAt.Unix()
	return ctx, common.OK, articleVO
}

func (a Article) UnlistedQueryArticleInfo(ctx *gin.Context) {
	ctx, err, data := Info(ctx)
	common.SendResponse(ctx, err, data)
	return
}

func (a Article) LoginAndQueryArticleInfo(ctx *gin.Context) {
	ctx, err, data := Info(ctx)
	//查询是否点赞
	token, err := tokenInfo(ctx)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	uid, _ := strconv.Atoi(token.Uid)
	articleVO := vo.ArticleInfoVO{}
	if reflect.TypeOf(data) == reflect.TypeOf(vo.ArticleInfoVO{}) && data != nil {
		_ = copier.Copy(&articleVO, data)
		sn, _ := strconv.ParseInt(articleVO.Sn, 10, 64)
		err, zanInfo := likeCommonFunc.LikeCommonFunc{}.CheckUserZanState(ctx, uid, cpgConst.ZERO, sn)
		if zanInfo.State != cpgConst.ONE && err == common.OK {
			articleVO.ZanState = true
		}
	}

	common.SendResponse(ctx, err, articleVO)
	return
}

// Info 根据sn查询
func (a Article) Info(ctx *gin.Context) {
	infoQO := new(qo.ArticleInfoQO)
	util.JsonConvert(ctx, infoQO)
	article := new(model.Article)

	if err := copier.Copy(article, infoQO); err != nil {
		common.SendResponse(ctx, common.ErrBind, err.Error())
		return
	}
	article.Sn, _ = strconv.ParseInt(infoQO.Sn, 10, 64)
	article = new(dao.ArticleDAO).SelectBySn(ctx, article)

	if article.Aid == 0 {
		common.SendResponse(ctx, common.ErrArticleNotExisted, "")
		return
	}

	articleVO := vo.ArticleInfoVO{}
	if err := copier.Copy(&articleVO, article); err != nil {
		common.SendResponse(ctx, common.ErrBind, err.Error())
		return
	}

	//查询是否点赞
	token, err := tokenInfo(ctx)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	uid, _ := strconv.Atoi(token.Uid)
	err, zanInfo := likeCommonFunc.LikeCommonFunc{}.CheckUserZanState(ctx, uid, cpgConst.ZERO, article.Sn)
	if zanInfo.State != cpgConst.ONE && err == common.OK {
		articleVO.ZanState = true
	}

	articleVO.Sn = strconv.FormatInt(article.Sn, 10)
	//userMap := userService.FindUser(ctx, []int{article.Uid}, "", "")
	userMap := userCommonFunc.IUser(userCommonFunc.UserCommonFunc{}).FindUser(ctx, []int{article.Uid}, "", "")
	articleVO.Author = userMap[uint(article.Uid)].UserName
	articleVO.CreateAt = article.CreatedAt.Unix()
	articleVO.UpdatedAt = article.UpdatedAt.Unix()
	common.SendResponse(ctx, common.OK, articleVO)
	return
}

func (a Article) List(ctx *gin.Context) {
	listQuery := new(qo.ArticleListQO)
	util.JsonConvert(ctx, listQuery)
	articleDAO := new(dao.ArticleDAO)
	copier.Copy(articleDAO, listQuery)
	copier.Copy(articleDAO, listQuery.Article)

	log.Println("请求参数:", listQuery)
	log.Println("articleDAO:", articleDAO)

	//是否查询自身的所有文章
	if listQuery.IsAllMyselfArticles {
		token, err := tokenInfo(ctx)
		if err != nil {
			common.SendResponse(ctx, err, "")
			return
		}
		articleDAO.Uid, err = strconv.Atoi(token.Uid)
		articleVO := articleDAO.FindArticles(ctx)

		//通过uid查询名称并填充
		//userMap := userService.FindUser(ctx, []int{articleDAO.Uid}, "", "")
		userMap := userCommonFunc.IUser(userCommonFunc.UserCommonFunc{}).FindUser(ctx, []int{articleDAO.Uid}, "", "")
		articleList := articleVO.ArticleDetailList
		for k, v := range articleVO.ArticleDetailList {
			articleList[k].Author = userMap[v.Uid].UserName
			//v.Author = userMap[uint(articleDAO.Uid)].UserName
		}
		common.SendResponse(ctx, common.OK, articleVO)
		return
	}
	articleVO := articleDAO.FindArticles(ctx)
	articleList := articleVO.ArticleDetailList
	for k, v := range articleList {
		//userMap := userService.FindUser(ctx, []int{int(v.Uid)}, "", "")
		userMap := userCommonFunc.IUser(userCommonFunc.UserCommonFunc{}).FindUser(ctx, []int{int(v.Uid)}, "", "")
		articleList[k].Author = userMap[v.Uid].UserName
	}
	common.SendResponse(ctx, common.OK, articleVO)
	return
}

func (a Article) Add(ctx *gin.Context) {
	addQO := new(qo.AddArticleQO)
	util.JsonConvert(ctx, addQO)
	article := new(model.Article)
	if err := copier.Copy(article, addQO); err != nil {
		common.SendResponse(ctx, common.ErrBind, err.Error())
		return
	}
	//用户UID从token中解析
	token, err := tokenInfo(ctx)
	if err != nil {
		common.SendResponse(ctx, common.ErrHandleToken, "")
		return
	}
	article.Uid, _ = strconv.Atoi(token.Uid)

	//新增文章的state为未审核1
	//TODO 后续需要增加审核功能，初始state应为0
	article.State = unreviewed

	article.Sn = common.Snowflake.NextID()
	log.Println(article.Sn)

	err = new(dao.ArticleDAO).CreatArticle(ctx, article)
	if err != nil {
		common.SendResponse(ctx, common.ErrDatabase, err.Error())
		return
	}
	resp := vo.AddArticleVO{Sn: article.Sn}
	common.SendResponse(ctx, common.OK, resp)
	return
}

func (a Article) Delete(ctx *gin.Context) {
	deleteQO := new(qo.ArticleInfoQO)
	util.JsonConvert(ctx, deleteQO)
	sn, _ := strconv.ParseInt(deleteQO.Sn, 10, 64)
	var articleList []model.Article

	tx := globalInit.Db.Where("sn", sn).Find(&articleList)

	if len(articleList) == 0 {
		common.SendResponse(ctx, common.ErrArticleNotExisted, "")
		return
	}
	if articleList[0].State == deleted {
		common.SendResponse(ctx, common.OK, "")
		return
	}

	//tokenInfo, _ := tokenInfo(ctx)
	//tokenUid,_ := strconv.Atoi(tokenInfo.Uid)
	//if tokenInfo.Root != cpgConst.Root &&  tokenUid != articleList[0].Uid{
	//	common.SendResponse(ctx, common.ErrAccessDenied, "")
	//}

	tx.Update("state", deleted).Commit()
	common.SendResponse(ctx, common.OK, "")
	return
}

func (a Article) Update(ctx *gin.Context) {
	updateQO := new(qo.UpdateArticleQO)
	util.JsonConvert(ctx, updateQO)
	if updateQO.Tags == "" &&
		updateQO.Title == "" &&
		updateQO.Content == "" &&
		updateQO.Cover == "" &&
		updateQO.State == "" {
		ok := common.OK
		ok.Message = "请输入更新内容"
		common.SendResponse(ctx, ok, "")
		return
	}
	updateDAO := new(dao.ArticleDAO)
	_ = copier.Copy(updateDAO, updateQO)
	sn, _ := strconv.ParseInt(updateQO.Sn, 10, 64)
	updateDAO.Sn = sn
	oldArticle := &model.Article{}

	//校验文章是否存在
	number := globalInit.Db.Model(&model.Article{}).
		Where("sn", updateDAO.Sn).
		First(oldArticle).RowsAffected
	if number == 0 {
		common.SendResponse(ctx, common.ErrArticleNotExisted, "")
		return
	}
	tokenInfo, _ := tokenInfo(ctx)
	tokenUid, _ := strconv.Atoi(tokenInfo.Uid)
	if tokenInfo.Root != cpgConst.Root && tokenUid != oldArticle.Uid {
		common.SendResponse(ctx, common.ErrAccessDenied, "暂无权限修改该文章！")
		return
	}

	//校验state
	state, _ := strconv.Atoi(updateQO.State)
	if state != unreviewed && state != published && state != removed && state != deleted {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}
	updateDAO.State = state
	err := updateDAO.UpdateArticle(ctx)
	if err != nil {
		common.SendResponse(ctx, common.ErrDatabase, err)
		return
	}
	common.SendResponse(ctx, common.OK, "")
	return
}

func (a Article) PopularArticlesList(ctx *gin.Context) {
	listQO := qo.PopularArticleQO{}
	util.JsonConvert(ctx, &listQO)
	listQO.Page.PageNum = cpgConst.ONE
	listQO.Page.PageSize = cpgConst.FOUR

	articleDAO := new(dao.ArticleDAO)

	countBool := func(a, b, c bool) (count int) {
		if a {
			count++
		}
		if b {
			count++
		}
		if c {
			count++
		}
		return count
	}

	if countBool(listQO.CmtNum, listQO.ZanNum, listQO.ViewNum) != cpgConst.ONE {
		common.SendResponse(ctx, common.ErrParam, "顺序不能同时为true")
	}
	_ = copier.Copy(articleDAO, listQO)
	articleVO := articleDAO.FindArticles(ctx)
	articleList := articleVO.ArticleDetailList
	for k, v := range articleList {
		userMap := userCommonFunc.IUser(userCommonFunc.UserCommonFunc{}).FindUser(ctx, []int{int(v.Uid)}, "", "")
		articleList[k].Author = userMap[v.Uid].UserName
	}
	common.SendResponse(ctx, common.OK, articleVO)
	return
}
