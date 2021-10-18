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
  @author: ethan.chen@cpgroup.cn
  @date:2021/10/15
  @description:
**/

type IComment interface {
	// UpdateCommentZan 服务间更新点赞信息
	UpdateCommentZan(cid int, isAdd bool) (err error)
}

type CommentCommonFunc struct{}

func (c CommentCommonFunc) Get() *CommentCommonFunc {
	return new(CommentCommonFunc)
}

func (c CommentCommonFunc)UpdateCommentZan(cid int, isAdd bool) (err error) {
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
