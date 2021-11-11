package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/internal/review/qo"
	"cpg-blog/internal/review/vo"
	"cpg-blog/pkg/commonFunc/articleCommonFunc"
	"cpg-blog/pkg/commonFunc/commentCommonFunc"
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

/**
  @author: ethan.chen@cpgroup.cn
  @date:2021/11/8
  @description: 审核服务实现
**/

type Review struct{}

//文章状态
const (
	//0-未审核
	unreviewed int = iota

	//1-已上线
	published

	//2-下线(审核失败)
	removed

	//3-用户删除
	deleted
)

func (v Review) ReviewArticleList(ctx *gin.Context) {
	data := articleCommonFunc.ArticleCommonFunc{}.FindArticlesByState(unreviewed)
	reviewArticleVo := new(vo.ReviewArticleVO)
	_ = copier.Copy(reviewArticleVo, &data)
	common.SendResponse(ctx, common.OK, reviewArticleVo)
}

func (v Review) ArticleReviewFailedList(ctx *gin.Context) {
	data := articleCommonFunc.ArticleCommonFunc{}.FindArticlesByState(removed)
	ArticleReviewFailedVo := new(vo.ReviewArticleVO)
	_ = copier.Copy(ArticleReviewFailedVo, &data)
	common.SendResponse(ctx, common.OK, ArticleReviewFailedVo)
}

func (v Review) ReviewCommentList(ctx *gin.Context) {
	data := commentCommonFunc.CommentCommonFunc{}.SelectCommentByState(unreviewed)
	ReviewCommentVo := new(vo.ReviewCommentVO)
	_ = copier.Copy(ReviewCommentVo, &data)
	common.SendResponse(ctx, common.OK, ReviewCommentVo)
}

func (v Review) CommentReviewFailedList(ctx *gin.Context) {
	data := commentCommonFunc.CommentCommonFunc{}.SelectCommentByState(removed)
	ReviewCommentVo := new(vo.ReviewCommentVO)
	_ = copier.Copy(ReviewCommentVo, &data)
	common.SendResponse(ctx, common.OK, ReviewCommentVo)
}

func (v Review) ReviewReplyList(ctx *gin.Context) {
	data := commentCommonFunc.CommentCommonFunc{}.SelectReplyByState(unreviewed)
	ReplyVo := new(vo.ReviewReplyVO)
	_ = copier.Copy(ReplyVo, &data)
	common.SendResponse(ctx, common.OK, ReplyVo)
}

func (v Review) ReplyReviewFailedList(ctx *gin.Context) {
	data := commentCommonFunc.CommentCommonFunc{}.SelectReplyByState(removed)
	ReplyVo := new(vo.ReviewReplyVO)
	_ = copier.Copy(ReplyVo, &data)
	common.SendResponse(ctx, common.OK, ReplyVo)
}

func (v Review) ReviewArticle(ctx *gin.Context, qo qo.ReviewArticleQO) {
	util.JsonConvert(ctx, qo)

	//根据sn查询相关文章
	articleMap := articleCommonFunc.ArticleCommonFunc{}.FindArticlesBySn(ctx, []int64{qo.Sn})
	if _, ok := articleMap[qo.Sn]; !ok {
		common.SendResponse(ctx, common.ErrArticleNotExisted, "")
		return
	}

	state := removed
	if qo.State {
		state = published
	}
	err := articleCommonFunc.ArticleCommonFunc{}.UpdateArticle(qo.Sn, state)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")

}

func (v Review) ReviewComment(ctx *gin.Context, qo qo.ReviewCommentQO) {
	util.JsonConvert(ctx, qo)

	//根据commentID查询相关评论信息
	err, _ := commentCommonFunc.CommentCommonFunc{}.SelectComment(qo.CommentId)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	state := removed
	if qo.State {
		state = published
	}
	err = commentCommonFunc.CommentCommonFunc{}.UpdateCommentState(qo.CommentId, state)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")
}

func (v Review) ReviewReply(ctx *gin.Context, qo qo.ReviewReplyQO) {
	util.JsonConvert(ctx, qo)

	//根据id查询相关回复信息
	err, _ := commentCommonFunc.CommentCommonFunc{}.SelectReply(qo.ReplyId)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	state := removed
	if qo.State {
		state = published
	}
	err = commentCommonFunc.CommentCommonFunc{}.UpdateReplyState(qo.ReplyId, state)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")
}
