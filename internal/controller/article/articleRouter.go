package article

import (
	"cpg-blog/internal/article/service/impl"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

var article = &impl.Article{}

// RegisterRoute 添加article服务路由
func (u Controller) RegisterRoute(g *gin.RouterGroup) {
	articleGroup := g.Group("/article")

	//查询文章详情
	articleGroup.POST("/info", article.Info)

	//查询文章列表
	articleGroup.POST("/list", article.List)

	//用户新增文章
	articleGroup.POST("/add", article.Add)
}
