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
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"log"
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

func (a Article) FindArticles(ctx *gin.Context, sn []int64) (articlesMap map[int64]model.Article) {
	var articles []model.Article
	articles = []model.Article{}
	articlesMap = map[int64]model.Article{}

	tx := globalInit.Db.WithContext(ctx).Model(&model.Article{})
	if len(sn) == cpgConst.ONE {
		tx.Where("sn", sn[0])
	}

	if len(sn) > cpgConst.ZERO {
		tx.Where("sn In", sn)
	}
	tx.Where("state", cpgConst.ONE).Find(articles)

	for _, v := range articles {
		articlesMap[v.Sn] = v
	}
	return articlesMap
}

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
		return
	}

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
	token, err := tokenInfo(ctx)
	if err != nil {
		common.SendResponse(ctx, common.ErrHandleToken, "")
		return
	}
	article.Uid, _ = strconv.Atoi(token.Uid)

	//新增文章的state为未审核1
	//TODO 后续需要增加审核功能，初始state应为0
	article.State = published

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
	copier.Copy(updateDAO, updateQO)
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
	updateDAO.State = state
	if state != unreviewed && state != published && state != removed && state != deleted {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}
	err := updateDAO.UpdateArticle(ctx)
	if err != nil {
		common.SendResponse(ctx, common.ErrDatabase, err)
		return
	}
	common.SendResponse(ctx, common.OK, "")
}

func (a Article) UpdateArticleEx(ctx *gin.Context, sn int64, view bool, cmt bool, zan bool, add bool) error {
	return dao.ArticleDAO{}.UpdateArticleEx(sn, view, cmt, zan, add)
}
