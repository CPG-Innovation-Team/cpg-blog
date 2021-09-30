package comment

import (
	"cpg-blog/internal/comment/service/impl"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

var comment = &impl.Comment{}

// RegisterRoute 添加Comment服务路由
func (c Controller) RegisterRoute(g *gin.RouterGroup)  {
	commentGroup := g.Group("/comment")

	commentGroup.POST("")
}
