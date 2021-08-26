package article

import (
	"cpg-blog/internal/article/service/impl"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

var article = &impl.Article{}

// RegisterRoute 添加article服务路由
func (u Controller) RegisterRoute(g *gin.RouterGroup) {
	query := g.Group("/article/query")
	update := g.Group("/article/update")

	//查询用户文章信息
	query.POST("/info", article.Info)

	//用户新增文章
	update.POST("/add", article.Add)
}
