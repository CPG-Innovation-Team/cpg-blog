package service

import "github.com/gin-gonic/gin"

type ILike interface {
	// Like 点赞文章/评论
	Like(ctx *gin.Context)

	// CancelLike 取消点赞文章/评论
	CancelLike(ctx *gin.Context)
}
