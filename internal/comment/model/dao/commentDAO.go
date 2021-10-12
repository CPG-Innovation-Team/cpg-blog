package dao

import (
	"cpg-blog/global/common"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/comment/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
  @author: ethan.chen@cpgroup.cn
  @date:2021/10/11
  @description:评论表
**/

type CommentDao model.Comment

var Db = &(globalInit.Db)

func (c CommentDao) BeforeCreate(tx *gorm.DB) (err error) {
	result := tx.Find(&c)
	if result.RowsAffected != 0 {
		return result.Error
	}
	return
}

func (c CommentDao) BeforeUpdate(tx *gorm.DB) (err error) {
	result := tx.Find(&c)
	if result.RowsAffected != 0 {
		return result.Error
	}
	return
}

func (c CommentDao) CreateComment(ctx *gin.Context) (cid uint, err error) {
	tx := globalInit.Transaction()
	err = func(db *gorm.DB) error {
		e := common.ErrDatabase
		tx.Create(c)
		if tx.Error != nil {
			e.Message = tx.Error.Error()
			return e
		}

		tx.Commit()
		if tx.Error != nil {
			e.Message = tx.Error.Error()
			return e
		}
		return nil
	}(tx)
	return c.Cid, err
}

func (c CommentDao) UpdateComment(ctx *gin.Context) (err error) {
	tx := globalInit.Transaction()
	err = func(db *gorm.DB) error {
		e := common.ErrDatabase
		tx.Select("content", "zan_num", "state").Updates(c)
		if tx.Error != nil {
			e.Message = tx.Error.Error()
			return e
		}

		tx.Commit()
		if tx.Error != nil {
			e.Message = tx.Error.Error()
			return e
		}
		return nil
	}(tx)
	return
}
