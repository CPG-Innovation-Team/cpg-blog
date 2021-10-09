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

	//文章所有评论及关系
	commentGroup.POST("/list", comment.List)

	//填写评论
	commentGroup.POST("/add", comment.Add)

	//删除评论
	commentGroup.POST("/delete", comment.Delete)

	//回复评论
	commentGroup.POST("/reply", comment.Reply)

	//删除回复
	commentGroup.POST("/reply/delete", comment.DeleteReply)
}
