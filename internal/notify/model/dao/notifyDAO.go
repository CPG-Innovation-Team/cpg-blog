package dao

import (
	"cpg-blog/global/common"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/notify/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
  @author: chenxi@cpgroup.cn
  @date:2021/12/8
  @description: 仅对cpg_blog_notify表进行操作
**/

type NotifyDao model.Notify

func (n NotifyDao) BeforeCreate(tx *gorm.DB) (err error) {
	result := tx.Find(&n)
	if result.RowsAffected != 0 {
		return result.Error
	}
	return
}

func (n NotifyDao) BeforeUpdate(tx *gorm.DB) (err error) {
	result := tx.Find(&n)
	if result.RowsAffected != 0 {
		return result.Error
	}
	return
}

func (n NotifyDao) Creat(ctx *gin.Context) (id int, err error) {
	tx := globalInit.Transaction()
	err = func(db *gorm.DB) error {
		e := common.ErrDatabase
		tx.Create(&n).Commit()
		if tx.Error != nil {
			tx.Rollback()
			e.Message = tx.Error.Error()
			return e
		}
		return nil
	}(tx)
	return n.Id, err
}

func (n NotifyDao) Update(ctx *gin.Context) (err error) {
	tx := globalInit.Transaction()
	err = func(db *gorm.DB) error {
		e := common.ErrDatabase
		tx.Updates(&n).Commit()
		if tx.Error != nil {
			tx.Rollback()
			e.Message = tx.Error.Error()
			return e
		}
		return nil
	}(tx)
	return err
}
