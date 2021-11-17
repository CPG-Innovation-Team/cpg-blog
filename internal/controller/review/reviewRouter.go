package review

import (
	"cpg-blog/internal/review/qo"
	"cpg-blog/internal/review/service/impl"
	"github.com/gin-gonic/gin"
)

/**
  @author: chenxi@cpgroup.cn
  @date:2021/11/11
  @description:
**/

type Controller struct{}

var review = &impl.Review{}

// RegisterRoute 添加review服务路由
func (c Controller) RegisterRoute(g *gin.RouterGroup) {

	query := g.Group("/review/query")
	update := g.Group("/review")

	//查询未审核的文章列表
	query.POST("/article/list", review.ReviewArticleList)

	//查询审核未通过的文章列表
	query.POST("/article/failed/list", review.ArticleReviewFailedList)

	//查询未审核的评论列表
	query.POST("/comment/list", review.ReviewCommentList)

	//查询审核未通过的评论列表
	query.POST("/comment/failed/list", review.CommentReviewFailedList)

	//查询未审核的评论回复列表
	query.POST("/reply/list", review.ReviewReplyList)

	//查询审核未通过的评论回复列表
	query.POST("/reply/failed/list", review.ReplyReviewFailedList)

	//审核新增的文章
	update.POST("/article", func(context *gin.Context) {
		query := qo.ReviewArticleQO{}
		review.ReviewArticle(context, query)
	})

	//审核新增的评论
	update.POST("/comment", func(context *gin.Context) {
		query := qo.ReviewCommentQO{}
		review.ReviewComment(context, query)
	})

	//审核新增的评论回复
	update.POST("/reply", func(context *gin.Context) {
		query := qo.ReviewReplyQO{}
		review.ReviewReply(context, query)
	})
}
