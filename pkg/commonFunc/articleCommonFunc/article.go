package articleCommonFunc

import (
	"cpg-blog/global/cpgConst"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/article/model"
	"cpg-blog/internal/article/model/dao"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
  @author: ethan.chen@cpgroup.cn
  @date:2021/10/15
  @description:
**/

type IArticle interface {
	// UpdateArticleEx 服务间更新文章扩展信息
	UpdateArticleEx(ctx *gin.Context, sn int64, view bool, cmt bool, zan bool, add bool) error

	// UpdateArticle 服务间更新文章状态
	UpdateArticle(sn int64, state int) (err error)

	// FindArticlesBySn 服务间根据sn查询文章信息，支持list
	FindPublishedArticlesBySn(ctx *gin.Context, sn []int64) (articlesMap map[int64]model.Article)

	// FindArticlesByState 服务间根据state查询文章
	FindArticlesByState(state int) map[int64]model.Article
}

type ArticleCommonFunc struct{}

func (ac ArticleCommonFunc) UpdateArticle(sn int64, state int) (err error) {
	tx := globalInit.Transaction().Model(&model.Article{}).Where("sn", sn)
	err = func(db *gorm.DB) error {
		if tx.Error != nil {
			return tx.Error
		}
		tx.Update("state", state)
		if tx.Error != nil {
			tx.Rollback()
			return tx.Error
		}
		return tx.Commit().Error
	}(tx)
	return err
}

func (ac ArticleCommonFunc) UpdateArticleEx(ctx *gin.Context, sn int64, view bool, cmt bool, zan bool, add bool) error{
	return dao.ArticleDAO{}.UpdateArticleEx(sn, view, cmt, zan, add)
}

func (ac ArticleCommonFunc) FindPublishedArticlesBySn(ctx *gin.Context, sn []int64) (articlesMap map[int64]model.Article){
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

func (ac ArticleCommonFunc) FindArticlesBySn(ctx *gin.Context, sn []int64) (articlesMap map[int64]model.Article){
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
	tx.Find(&articles)
	for _, v := range articles {
		articlesMap[v.Sn] = v
	}
	return articlesMap
}

func (ac ArticleCommonFunc) FindArticlesByState(state int) map[int64]model.Article {
	return dao.ArticleDAO{}.SelectArticleByState(uint(state))
}
