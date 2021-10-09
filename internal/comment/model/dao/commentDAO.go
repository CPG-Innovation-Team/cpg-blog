package dao

import (
	"cpg-blog/global/common"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/comment/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentDao model.Comment

var Db = &(globalInit.Db)

func (c CommentDao) AddComment(ctx *gin.Context) (err error) {
	tx := globalInit.Transaction()
	err = func(db *gorm.DB) error{
		e := common.ErrDatabase
		tx.Create(c)
		if tx.Error != nil{
			e.Message = tx.Error.Error()
			return e
		}

		tx.Commit()
		if tx.Error != nil{
			e.Message = tx.Error.Error()
			return e
		}
		return nil
	} (tx)
	return
}