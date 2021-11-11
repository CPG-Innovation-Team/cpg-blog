package commentCommonFunc

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/comment/model"
	"cpg-blog/internal/comment/model/dao"
	"reflect"
)

/**
  @author: chenxi@cpgroup.cn
  @date:2021/10/15
  @description:
**/

type IComment interface {
	// UpdateCommentZan 服务间更新点赞信息
	UpdateCommentZan(cid int, isAdd bool) (err error)

	//SelectComment 查询评论
	SelectComment(cid int) (err error, comment model.Comment)

	//SelectCommentByState 根据状态查询评论
	SelectCommentByState(state int) (comment map[uint]model.Comment)

	//SelectReplyByState 根据状态查询回复
	SelectReplyByState(state int) (comment map[uint]model.CommentReply)

	//UpdateCommentState 更新评论状态。state默认为0（未审核）
	UpdateCommentState(cid int, state int) (err error)

	//SelectReply 查询评论回复
	SelectReply(id int) (err error, reply model.CommentReply)

	// UpdateReplyState 更新评论回复状态。state默认为0（未审核）
	UpdateReplyState(id int, state int) (err error)
}

type CommentCommonFunc struct{}

func (c CommentCommonFunc) UpdateCommentZan(cid int, isAdd bool) (err error) {
	comment := model.Comment{}
	globalInit.Db.Where("cid = ? and state = ?", cid, cpgConst.ONE).Find(&comment)

	if reflect.DeepEqual(model.Comment{}, comment) {
		e := common.ErrParam
		e.Message = "Not Find Comment Or Comment Not Online"
		return e
	}
	if !isAdd && comment.ZanNum == cpgConst.ZERO {
		return nil
	}

	zanNum := comment.ZanNum
	if isAdd {
		zanNum += cpgConst.ONE
	} else {
		zanNum -= cpgConst.ONE
	}

	return dao.Comment{}.UpdateCommentZan(cid, zanNum)
}

func (c CommentCommonFunc) SelectComment(cid int) (err error, comment model.Comment) {
	comment = model.Comment{}
	globalInit.Db.Where("cid = ?", cid).Find(&comment)
	if reflect.DeepEqual(model.Comment{}, comment) {
		e := common.ErrParam
		e.Message = "Not Find Comment"
		return e, comment
	}
	return common.OK, comment
}

func (c CommentCommonFunc) SelectCommentByState(state int) (comment map[uint]model.Comment) {
	return dao.Comment{}.SelectByState(uint(state))
}

func (c CommentCommonFunc) SelectReplyByState(state int) (comment map[uint]model.CommentReply) {
	return dao.CommentReply{}.SelectByState(uint(state))
}

func (c CommentCommonFunc) UpdateCommentState(cid int, state int) (err error) {
	return dao.Comment{}.UpdateCommentState(cid, state)
}

func (c CommentCommonFunc) SelectReply(id int) (err error, reply model.CommentReply) {
	reply = model.CommentReply{}
	globalInit.Db.Where("id = ?", id).Find(&reply)
	if reflect.DeepEqual(model.CommentReply{}, reply) {
		e := common.ErrParam
		e.Message = "Not Find Reply"
		return e, reply
	}
	return common.OK, reply
}

func (c CommentCommonFunc) UpdateReplyState(id int, state int) (err error) {
	return dao.CommentReply{}.UpdateCommentReplyStateByCid(id, state)
}
