package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/article/model"
	"cpg-blog/internal/article/model/dao"
	"cpg-blog/internal/article/qo"
	"cpg-blog/internal/article/vo"
	"cpg-blog/internal/user/service/impl"
	"cpg-blog/middleware/jwt"
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"log"
	"strconv"
)

type Article struct{}

//文章状态
const (
	unreviewed int = iota
	published
	removed
	deleted
)

var userService = &impl.Users{}

// Info 根据sn查询
func (a Article) Info(ctx *gin.Context) {
	infoQO := new(qo.ArticleInfoQO)
	util.JsonConvert(ctx, infoQO)
	article := new(model.Article)

	if err := copier.Copy(article, infoQO); err != nil {
		common.SendResponse(ctx, common.ErrBind, err.Error())
	}
	article = new(dao.ArticleDAO).SelectBySn(ctx, article)

	if article.Aid == 0 {
		common.SendResponse(ctx, common.ErrArticleNotExisted, "")
	} else if article.State != published {
		common.SendResponse(ctx, common.OK, "当前文章已下线或已删除！")
	} else {
		articleVO := vo.ArticleInfoVO{}
		if err := copier.Copy(&articleVO, article); err != nil {
			common.SendResponse(ctx, common.ErrBind, err.Error())
		}
		userMap := userService.FindUser(ctx, []int{article.Uid}, "", "")
		articleVO.Author = userMap[uint(article.Uid)].UserName
		articleVO.CreateAt = article.CreatedAt.Unix()
		articleVO.UpdatedAt = article.UpdatedAt.Unix()
		common.SendResponse(ctx, common.OK, articleVO)
	}
}

func (a Article) List(ctx *gin.Context) {
	listQuery := new(qo.ArticleListQO)
	util.JsonConvert(ctx, listQuery)
	articleDAO := new(dao.ArticleDAO)
	copier.Copy(articleDAO, listQuery)
	copier.Copy(articleDAO, listQuery.Article)

	//默认查询已审核通过上线的文章
	articleDAO.State = 1

	log.Println("请求参数:", listQuery)
	log.Println("articleDAO:", articleDAO)

	//是否查询自身的所有文章
	if listQuery.IsAllMyselfArticles {
		token, err := jwt.NewJWT().ParseToken(ctx.Request.Header.Get("token"))
		if err != nil {
			common.SendResponse(ctx, err, "")
			return
		}
		articleDAO.Uid, err = strconv.Atoi(token.Uid)
		articleVO := articleDAO.FindArticles(ctx)

		//通过uid查询名称并填充
		userMap := userService.FindUser(ctx, []int{articleDAO.Uid}, "", "")
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
		userMap := userService.FindUser(ctx, []int{int(v.Uid)}, "", "")
		articleList[k].Author = userMap[v.Uid].UserName
	}
	common.SendResponse(ctx, common.OK, articleVO)
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
	token, err := jwt.NewJWT().ParseToken(ctx.Request.Header.Get("token"))
	if err != nil {
		common.SendResponse(ctx, common.ErrTokenInvalid, "")
		return
	}
	article.Uid, _ = strconv.Atoi(token.Uid)


	//新增文章的state为未审核1
	//TODO 后续需要增加审核功能，初始state应为0
	article.State = 1

	article.Sn = common.Snowflake.NextID()
	log.Println(article.Sn)

	err = new(dao.ArticleDAO).CreatArticle(ctx, article)
	if err != nil {
		common.SendResponse(ctx, common.ErrDatabase, err.Error())
		return
	}
	resp := vo.AddArticleVO{Sn: article.Sn}
	common.SendResponse(ctx, common.OK, resp)
}

func (a Article) Delete(ctx *gin.Context) {
	deleteQO := new(qo.ArticleInfoQO)
	util.JsonConvert(ctx, deleteQO)
	var articleList []model.Article
	tx := globalInit.Db.Where("sn", deleteQO.Sn).Find(&articleList)
	if len(articleList) == 0 {
		common.SendResponse(ctx, common.ErrArticleNotExisted, "")
		return
	}
	if articleList[0].State == 1{
		common.SendResponse(ctx, common.OK, "")
		return
	}
	tx.Update("state", 1).Commit()
	common.SendResponse(ctx, common.OK, "")
}
