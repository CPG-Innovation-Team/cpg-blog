package like

import (
	"cpg-blog/internal/like/service/impl"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

var like = &impl.Like{}

// RegisterRoute 添加Like服务路由
func (L Controller) RegisterRoute(g *gin.RouterGroup) {
	likeGroup := g.Group("/like")

	//点赞
	likeGroup.POST("", like.Like)

	//取消点赞
	likeGroup.POST("/cancel", like.CancelLike)
}
