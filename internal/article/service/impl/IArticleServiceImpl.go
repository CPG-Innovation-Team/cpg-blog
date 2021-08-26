package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/internal/article/model"
	"cpg-blog/internal/article/model/dao"
	"cpg-blog/internal/article/qo"
	"cpg-blog/internal/article/vo"
	"cpg-blog/internal/oauth"
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"strconv"
)

type Article struct{}

func (a Article) Info(ctx *gin.Context) {

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
	token, err := oauth.NewJWT().ParseToken(ctx.Request.Header.Get("token"))
	article.Uid, err = strconv.Atoi(token.Uid)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}

	//新增文章的state为未审核0
	article.State = 0

	//TODO 生成sn规则，数据库唯一且不能重复
	article.Sn = 12334

	err = new(dao.ArticleDAO).CreatArticle(ctx, article)
	if err != nil {
		common.SendResponse(ctx, common.ErrDatabase, err.Error())
		return
	}
	resp := vo.AddArticleVO{Sn: article.Sn}
	common.SendResponse(ctx, common.OK, resp)
}
