package dao

import (
	"cpg-blog/global/common"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/object-processing/model"
	"gorm.io/gorm"
)

/**
  @author: chenxi@cpgroup.cn
  @date:2022/6/14
  @description:
**/

type File model.File

var Db = &(globalInit.Db)

func (fd File) BeforeCreate(tx *gorm.DB) (err error) {
	result := tx.Find(&fd)
	if result.RowsAffected != 0 {
		return result.Error
	}
	return
}

func (fd File) BeforeUpdate(tx *gorm.DB) (err error) {
	result := tx.Find(&fd)
	if result.RowsAffected != 0 {
		return result.Error
	}
	return
}

func (fd File) Creat() (id int, err error) {
	tx := globalInit.Transaction()
	err = func(db *gorm.DB) error {
		e := common.ErrDatabase
		tx.Create(&fd).Commit()
		if tx.Error != nil {
			tx.Rollback()
			e.Message = tx.Error.Error()
			return e
		}
		return nil
	}(tx)
	return fd.Id, err
}

func (fd File) Update() (err error) {
	tx := globalInit.Transaction()
	err = func(db *gorm.DB) error {
		e := common.ErrDatabase
		tx.Updates(&fd).Commit()
		if tx.Error != nil {
			tx.Rollback()
			e.Message = tx.Error.Error()
			return e
		}
		return nil
	}(tx)
	return err
}
