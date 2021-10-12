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
  @description:评论回复表
**/

type CommentReplyDao model.CommentReply

//UpdateCommentReply
/**
* @Author: ethan.chen@cpgroup.cn
* @Date: 2021/10/11 15:50
* @Description: 更新、删除回复
* @Params: model.CommentReply
* @Return: error
**/

func (c CommentReplyDao) BeforeCreate(tx *gorm.DB) (err error) {
	result := tx.Find(&c)
	if result.RowsAffected != 0 {
		return result.Error
	}
	return
}

func (c CommentReplyDao) BeforeUpdate(tx *gorm.DB) (err error) {
	result := tx.Find(&c)
	if result.RowsAffected != 0 {
		return result.Error
	}
	return
}

func (c CommentReplyDao) UpdateCommentReply(ctx *gin.Context) (err error) {
	tx := globalInit.Transaction()
	tx.Model(c)
	err = func(db *gorm.DB) error {
		e := common.ErrDatabase
		tx.Select("content", "state").Where("cid", c.Cid).Updates(c)

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

func (c CommentReplyDao) CreateCommentReply(ctx *gin.Context) (replyId uint, err error) {
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
	return c.Id, err
}
