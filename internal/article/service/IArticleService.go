package service

import "github.com/gin-gonic/gin"

type IArticle interface {
	//Info 查询文章
	Info(ctx *gin.Context)

	// Add 新增文章
	Add(ctx *gin.Context)
}
