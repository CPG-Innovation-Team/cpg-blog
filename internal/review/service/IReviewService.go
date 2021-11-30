package service

import (
	"cpg-blog/internal/review/qo"
	"github.com/gin-gonic/gin"
)

/**
  @author: chenxi@cpgroup.cn
  @date:2021/11/8
  @description: 审核服务
**/

type IVerify interface {
	//ReviewArticleList 未审核的文章列表
	ReviewArticleList(ctx *gin.Context)

	//ArticleReviewFailedList 审核未通过的文章列表
	ArticleReviewFailedList(ctx *gin.Context)

	//ReviewCommentList 未审核的评论列表
	ReviewCommentList(ctx *gin.Context)

	//CommentReviewFailedList 审核未通过的评论列表
	CommentReviewFailedList(ctx *gin.Context)

	//ReviewReplyList 未审核的评论回复列表
	ReviewReplyList(ctx *gin.Context)

	//ReplyReviewFailedList 审核未通过的评论回复列表
	ReplyReviewFailedList(ctx *gin.Context)

	// ReviewArticle 审核新增的文章
	ReviewArticle(ctx *gin.Context, qo *qo.ReviewArticleQO)

	//ReviewComment 审核新增的评论
	ReviewComment(ctx *gin.Context, qo *qo.ReviewCommentQO)

	//ReviewReply 审核新增的评论回复
	ReviewReply(ctx *gin.Context, qo *qo.ReviewReplyQO)
}
