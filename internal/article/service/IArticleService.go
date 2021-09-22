package service

import "github.com/gin-gonic/gin"

type IArticle interface {
	//Info 查询文章详情
	Info(ctx *gin.Context)

	// List 搜索文章
	List(ctx *gin.Context)

	// Add 新增文章
	Add(ctx *gin.Context)

	// Delete 删除文章
	Delete(ctx *gin.Context)

	// Update 更新文章
	Update(ctx *gin.Context)
}
