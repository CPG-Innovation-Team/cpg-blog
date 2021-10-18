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

type Comment model.Comment

var Db = &(globalInit.Db)

func (c Comment) BeforeCreate(tx *gorm.DB) (err error) {
	result := tx.Find(&c)
	if result.RowsAffected != 0 {
		return result.Error
	}
	return
}

func (c Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	result := tx.Find(&c)
	if result.RowsAffected != 0 {
		return result.Error
	}
	return
}

func (c Comment) CreateComment(ctx *gin.Context) (cid uint, err error) {
	tx := globalInit.Transaction()
	err = func(db *gorm.DB) error {
		e := common.ErrDatabase
		tx.Create(&c)
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

func (c Comment) UpdateComment(ctx *gin.Context) (err error) {
	tx := globalInit.Transaction()
	err = func(db *gorm.DB) error {
		e := common.ErrDatabase
		tx.Select("content", "zan_num", "state").Updates(&c)
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

func (c Comment) UpdateCommentZan(cid int, add int) (err error) {
	c.Cid = uint(cid)
	tx := globalInit.Transaction()
	e := common.ErrDatabase
	err = func(db *gorm.DB) error {
		tx.Model(&c).Update("zan_num", add)
		tx.Commit()
		if tx.Error != nil {
			e.Message = tx.Error.Error()
			return e
		}
		return nil
	}(tx)
	return
}
