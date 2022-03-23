package service

import (
	"github.com/gin-gonic/gin"
)

type IArticle interface {

	//UnlistedQueryArticleInfo 未登录查询文章详情
	UnlistedQueryArticleInfo(ctx *gin.Context)

	//LoginAndQueryArticleInfo 登陆查询文章详情
	LoginAndQueryArticleInfo(ctx *gin.Context)

	// List 搜索文章
	List(ctx *gin.Context)

	// Add 新增文章
	Add(ctx *gin.Context)

	// Delete 删除文章
	Delete(ctx *gin.Context)

	// Update 更新文章
	Update(ctx *gin.Context)

	//PopularArticlesList 热门文章，默认四篇
	PopularArticlesList(ctx *gin.Context)
}
