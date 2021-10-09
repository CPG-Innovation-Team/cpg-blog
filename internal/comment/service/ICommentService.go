package service

import "github.com/gin-gonic/gin"

type Comment interface {
	// List 文章所有评论及关系
	List(ctx *gin.Context)

	// Add 用户填写评论
	Add(ctx *gin.Context)

	// Delete 删除评论
	Delete(ctx *gin.Context)

	// Reply 回复评论
	Reply(ctx *gin.Context)

	// DeleteReply 删除回复
	DeleteReply(ctx *gin.Context)
}
