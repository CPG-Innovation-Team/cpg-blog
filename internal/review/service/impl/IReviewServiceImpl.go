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
	_ = copier.Copy(&reviewArticleVo.ArticleMap, &data)
	common.SendResponse(ctx, common.OK, reviewArticleVo)
}

func (v Review) ArticleReviewFailedList(ctx *gin.Context) {
	data := articleCommonFunc.ArticleCommonFunc{}.FindArticlesByState(removed)
	ArticleReviewFailedVo := new(vo.ReviewArticleVO)
	_ = copier.Copy(&ArticleReviewFailedVo.ArticleMap, &data)
	common.SendResponse(ctx, common.OK, ArticleReviewFailedVo)
}

func (v Review) ReviewCommentList(ctx *gin.Context) {
	data := commentCommonFunc.CommentCommonFunc{}.SelectCommentByState(unreviewed)
	ReviewCommentVo := new(vo.ReviewCommentVO)
	_ = copier.Copy(&ReviewCommentVo.CommentMap, &data)
	common.SendResponse(ctx, common.OK, ReviewCommentVo)
}

func (v Review) CommentReviewFailedList(ctx *gin.Context) {
	data := commentCommonFunc.CommentCommonFunc{}.SelectCommentByState(removed)
	ReviewCommentVo := new(vo.ReviewCommentVO)
	_ = copier.Copy(&ReviewCommentVo.CommentMap, &data)
	common.SendResponse(ctx, common.OK, ReviewCommentVo)
}

func (v Review) ReviewReplyList(ctx *gin.Context) {
	data := commentCommonFunc.CommentCommonFunc{}.SelectReplyByState(unreviewed)
	ReplyVo := new(vo.ReviewReplyVO)
	_ = copier.Copy(&ReplyVo.ReplyMap, &data)
	common.SendResponse(ctx, common.OK, ReplyVo)
}

func (v Review) ReplyReviewFailedList(ctx *gin.Context) {
	data := commentCommonFunc.CommentCommonFunc{}.SelectReplyByState(removed)
	ReplyVo := new(vo.ReviewReplyVO)
	_ = copier.Copy(&ReplyVo.ReplyMap, &data)
	common.SendResponse(ctx, common.OK, ReplyVo)
}

func (v Review) ReviewArticle(ctx *gin.Context, query *qo.ReviewArticleQO) {
	util.JsonConvert(ctx, query)

	//根据sn查询相关文章
	articleMap := articleCommonFunc.ArticleCommonFunc{}.FindArticlesBySn(ctx, []int64{query.Sn})
	if _, ok := articleMap[query.Sn]; !ok {
		common.SendResponse(ctx, common.ErrArticleNotExisted, "")
		return
	}

	state := removed
	if query.State {
		state = published
	}
	err := articleCommonFunc.ArticleCommonFunc{}.UpdateArticle(query.Sn, state)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")

}

func (v Review) ReviewComment(ctx *gin.Context, query *qo.ReviewCommentQO) {
	util.JsonConvert(ctx, query)

	//根据commentID查询相关评论信息
	err, _ := commentCommonFunc.CommentCommonFunc{}.SelectComment(query.CommentId)
	if err != nil && err != common.OK {
		common.SendResponse(ctx, err, "")
		return
	}
	state := removed
	if query.State {
		state = published
	}
	err = commentCommonFunc.CommentCommonFunc{}.UpdateCommentState(query.CommentId, state)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")
}

func (v Review) ReviewReply(ctx *gin.Context, query *qo.ReviewReplyQO) {
	util.JsonConvert(ctx, query)

	//根据id查询相关回复信息
	err, _ := commentCommonFunc.CommentCommonFunc{}.SelectReply(query.ReplyId)
	if err != nil && err != common.OK{
		common.SendResponse(ctx, err, "")
		return
	}
	state := removed
	if query.State {
		state = published
	}
	err = commentCommonFunc.CommentCommonFunc{}.UpdateReplyState(query.ReplyId, state)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")
}
