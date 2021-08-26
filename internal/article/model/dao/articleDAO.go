package dao

import (
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/article/model"
	"github.com/gin-gonic/gin"
)

type ArticleDAO struct{}

func (ad ArticleDAO) CreatArticle(ctx *gin.Context, article *model.Article) (err error) {
	tx := globalInit.Db.Create(article)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
