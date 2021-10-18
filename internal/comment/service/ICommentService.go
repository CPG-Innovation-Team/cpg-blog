package service

import "github.com/gin-gonic/gin"

type IComment interface {
	// List 文章所有评论及关系
	List(ctx *gin.Context)

	// Add 用户填写评论
	Add(ctx *gin.Context)

	// Delete 删除评论
	Delete(ctx *gin.Context)

	// AddReply 回复评论
	AddReply(ctx *gin.Context)

	// DeleteReply 删除回复
	DeleteReply(ctx *gin.Context)

	// UpdateCommentZan 服务间更新点赞信息
	//UpdateCommentZan(cid int, isAdd bool) (err error)
}
