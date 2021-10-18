package service

import "github.com/gin-gonic/gin"

type ILike interface {
	// Like 点赞文章/评论
	Like(ctx *gin.Context)

	// CancelLike 取消点赞文章/评论
	CancelLike(ctx *gin.Context)

	// Update 服务间更新点赞表state(0未删除，1删除),objId为文章Id或评论Id
	//Update(ctx *gin.Context, objId int64, state int) error
}
