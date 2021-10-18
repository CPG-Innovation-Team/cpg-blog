package articleCommonFunc

import (
	"cpg-blog/global/cpgConst"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/article/model"
	"cpg-blog/internal/article/model/dao"
	"github.com/gin-gonic/gin"
)

/**
  @author: ethan.chen@cpgroup.cn
  @date:2021/10/15
  @description:
**/

type IArticle interface {
	// UpdateArticleEx 服务间更新文章扩展信息
	UpdateArticleEx(ctx *gin.Context, sn int64, view bool, cmt bool, zan bool, add bool) error

	// FindArticles 服务间查询文章信息，支持list
	FindArticles(ctx *gin.Context, sn []int64) (articlesMap map[int64]model.Article)
}

type ArticleCommonFunc struct{}

func (ac ArticleCommonFunc) UpdateArticleEx(ctx *gin.Context, sn int64, view bool, cmt bool, zan bool, add bool) error{
	return dao.ArticleDAO{}.UpdateArticleEx(sn, view, cmt, zan, add)
}

func (ac ArticleCommonFunc)FindArticles(ctx *gin.Context, sn []int64) (articlesMap map[int64]model.Article){
	var articles []model.Article
	articles = []model.Article{}
	articlesMap = map[int64]model.Article{}

	tx := globalInit.Db.WithContext(ctx).Model(&model.Article{})
	if len(sn) == cpgConst.ONE {
		tx.Where("sn", sn[0])
	}

	if len(sn) > cpgConst.ONE {
		tx.Where("sn In", sn)
	}
	tx.Where("state", cpgConst.ONE).Find(&articles)
	for _, v := range articles {
		articlesMap[v.Sn] = v
	}
	return articlesMap
}
