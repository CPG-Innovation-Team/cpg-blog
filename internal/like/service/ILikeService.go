package service

import "github.com/gin-gonic/gin"

type ILike interface {
	// Like 点赞文章
	Like(ctx *gin.Context)
}
