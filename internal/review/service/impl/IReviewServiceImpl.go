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
	"strconv"
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
	reviewArticleVo := new(vo.ReviewArticleListVO)
	_ = copier.Copy(&reviewArticleVo.ArticleMap, &data)
	common.SendResponse(ctx, common.OK, reviewArticleVo)
}

func (v Review) ArticleReviewFailedList(ctx *gin.Context) {
	data := articleCommonFunc.ArticleCommonFunc{}.FindArticlesByState(removed)
	ArticleReviewFailedVo := new(vo.ReviewArticleListVO)
	_ = copier.Copy(&ArticleReviewFailedVo.ArticleMap, &data)
	common.SendResponse(ctx, common.OK, ArticleReviewFailedVo)
}

func (v Review) ReviewCommentList(ctx *gin.Context) {
	data := commentCommonFunc.CommentCommonFunc{}.SelectCommentByState(unreviewed)
	ReviewCommentVo := new(vo.ReviewCommentListVO)
	_ = copier.Copy(&ReviewCommentVo.CommentMap, &data)
	common.SendResponse(ctx, common.OK, ReviewCommentVo)
}

func (v Review) CommentReviewFailedList(ctx *gin.Context) {
	data := commentCommonFunc.CommentCommonFunc{}.SelectCommentByState(removed)
	ReviewCommentVo := new(vo.ReviewCommentListVO)
	_ = copier.Copy(&ReviewCommentVo.CommentMap, &data)
	common.SendResponse(ctx, common.OK, ReviewCommentVo)
}

func (v Review) ReviewReplyList(ctx *gin.Context) {
	data := commentCommonFunc.CommentCommonFunc{}.SelectReplyByState(unreviewed)
	ReplyVo := new(vo.ReviewReplyListVO)
	_ = copier.Copy(&ReplyVo.ReplyMap, &data)
	common.SendResponse(ctx, common.OK, ReplyVo)
}

func (v Review) ReplyReviewFailedList(ctx *gin.Context) {
	data := commentCommonFunc.CommentCommonFunc{}.SelectReplyByState(removed)
	ReplyVo := new(vo.ReviewReplyListVO)
	_ = copier.Copy(&ReplyVo.ReplyMap, &data)
	common.SendResponse(ctx, common.OK, ReplyVo)
}

func (v Review) ReviewArticle(ctx *gin.Context, query *qo.ReviewArticleQO) {
	util.JsonConvert(ctx, query)
	sn, _ := strconv.ParseInt(query.Sn, 10, 64)
	//根据sn查询相关文章
	articleMap := articleCommonFunc.ArticleCommonFunc{}.FindArticlesBySn(ctx, []int64{sn})
	if _, ok := articleMap[sn]; !ok {
		common.SendResponse(ctx, common.ErrArticleNotExisted, "")
		return
	}

	state := removed
	if query.State {
		state = published
	}
	err := articleCommonFunc.ArticleCommonFunc{}.UpdateArticle(sn, state)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")

}

func (v Review) ReviewComment(ctx *gin.Context, query *qo.ReviewCommentQO) {
	util.JsonConvert(ctx, query)

	//根据commentID查询相关评论信息
	err, commentList := commentCommonFunc.CommentCommonFunc{}.SelectComment(query.CommentId)
	if err != nil && err != common.OK {
		common.SendResponse(ctx, err, "")
		return
	}

	if len(commentList) != len(query.CommentId) {
		m := make(map[int]bool)

		for _, comment := range commentList {
			m[int(comment.Cid)] = true
		}

		commentVo := new(vo.ReviewCommentVO)
		for _, commentId := range query.CommentId {
			if _, ok := m[commentId]; !ok {
				commentVo.CommentList = append(commentVo.CommentList, commentId)
			}
		}
		e := common.ErrParam
		e.Message = "评论不存在!"
		common.SendResponse(ctx, e, commentVo)
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
	err, replyList := commentCommonFunc.CommentCommonFunc{}.SelectReply(query.ReplyId)
	if err != nil && err != common.OK {
		common.SendResponse(ctx, err, "")
		return
	}

	if len(replyList) != len(query.ReplyId) {
		m := make(map[int]bool)

		for _, reply := range replyList {
			m[int(reply.Id)] = true
		}

		replyVo := new(vo.ReviewReplyVO)
		for _, replyId := range query.ReplyId {
			if _, ok := m[replyId]; !ok {
				replyVo.ReplyList = append(replyVo.ReplyList, replyId)
			}
		}
		e := common.ErrParam
		e.Message = "回复不存在!"
		common.SendResponse(ctx, e, replyVo)
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
