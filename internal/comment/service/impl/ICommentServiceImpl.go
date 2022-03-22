package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/comment/model"
	"cpg-blog/internal/comment/model/dao"
	"cpg-blog/internal/comment/qo"
	"cpg-blog/internal/comment/vo"
	"cpg-blog/middleware/jwt"
	"cpg-blog/pkg/commonFunc/articleCommonFunc"
	"cpg-blog/pkg/commonFunc/likeCommonFunc"
	"cpg-blog/pkg/commonFunc/userCommonFunc"
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"reflect"
	"strconv"
)

type Comment struct{}

/**
* @Author: ethan.chen@cpgroup.cn
* @Date: 2021/10/12 14:02
* @Description: 查询未删除的评论信息
* @Params: cid
* @Return: model.Comment
**/
func (c Comment) commentInfo(cid int) (comment model.Comment) {
	globalInit.Db.Model(&model.Comment{}).
		Where("cid = ? and state = ?", cid, cpgConst.ONE).
		First(&comment)
	return
}

/**
* @Author: ethan.chen@cpgroup.cn
* @Date: 2021/10/12 14:42
* @Description: 查询token信息
* @Params: *gin.Context
* @Return: info *jwt.CustomClaims, err error
**/
func (c Comment) tokenInfo(ctx *gin.Context) (info *jwt.CustomClaims, err error) {
	return jwt.NewJWT().ParseToken(ctx.Request.Header.Get("token"))
}

//UpdateCommentZan
/**
* @Author: ethan.chen@cpgroup.cn
* @Date: 2021/10/13 15:22
* @Description: 其他服务更新评论点赞信息
* @Params: cid int, isAdd bool
* @Return: error
**/
//func (c Comment) UpdateCommentZan(cid int, isAdd bool) (err error) {
//	comment := model.Comment{}
//	globalInit.Db.Where("cid = ? and state = ?", cid, cpgConst.ONE).Find(&comment)
//
//	if reflect.DeepEqual(model.Comment{}, comment) {
//		e := common.ErrParam
//		e.Message = "Not Find Comment Or Comment Not Online"
//		return e
//	}
//	if !isAdd && comment.ZanNum == cpgConst.ZERO {
//		return nil
//	}
//
//	zanNum := comment.ZanNum
//	if isAdd {
//		zanNum += cpgConst.ONE
//	} else {
//		zanNum -= cpgConst.ONE
//	}
//
//	return dao.Comment{}.UpdateCommentZan(cid, zanNum)
//}

//List
/**
* @Author: ethan.chen@cpgroup.cn
* @Date: 2021/10/12 17:13
* @Description: 查询文章所有评论及回复
* @Params:
* @Return:
**/
func (c Comment) List(ctx *gin.Context) {
	listQO := qo.ListQO{}
	util.JsonConvert(ctx, &listQO)
	sn, _ := strconv.ParseInt(listQO.Sn, 10, 64)
	listMap := make(map[int]vo.CommentListVO)

	//通过sn查询文章所有评论，生成以floor为key,listVO为value
	var comments []model.Comment
	globalInit.Db.Model(model.Comment{}).Where("sn", sn).Find(&comments)
	if len(comments) == 0 {
		common.SendResponse(ctx, common.OK, listMap)
		return
	}

	uidList := make([]int, 5)

	for _, v := range comments {
		commentInfo := vo.CommentListVO{}
		_ = copier.Copy(&commentInfo, &v)
		uidList = append(uidList, int(v.UID))

		//查询该条评论下所有回复
		var reply []model.CommentReply
		globalInit.Db.Where("cid = ? and state = ?", v.Cid, cpgConst.ONE).Find(&reply)

		_ = copier.Copy(&commentInfo.ReplyList, &reply)

		for _, rv := range commentInfo.ReplyList {
			uidList = append(uidList, int(rv.UID))
		}
		listMap[v.Floor] = commentInfo
	}

	userMap := userCommonFunc.IUser(userCommonFunc.UserCommonFunc{}).FindUser(ctx, uidList, "", "")

	//查询登陆用户uid
	tokenInfo, err := jwt.NewJWT().ParseToken(ctx.Request.Header.Get("token"))
	var uid int
	if err == nil {
		//common.SendResponse(ctx, err, "")
		//return
		uid, _ = strconv.Atoi(tokenInfo.Uid)
	}
	//uid, _ := strconv.Atoi(tokenInfo.Uid)

	for k, v := range listMap {
		commentInfo := listMap[k]
		commentInfo.NickName = userMap[v.UID].Nickname

		if uid != cpgConst.ZERO{
			err, zanInfo := likeCommonFunc.LikeCommonFunc{}.CheckUserZanState(ctx, uid, cpgConst.ONE, int64(v.Cid))
			if zanInfo.State != cpgConst.ONE && err == common.OK {
				commentInfo.ZanState = true
			}
		}

		if v.ReplyList != nil {
			for rk := range commentInfo.ReplyList {
				commentInfo.ReplyList[rk].NickName = userMap[commentInfo.ReplyList[rk].UID].Nickname
			}
		}
		listMap[k] = commentInfo
	}

	common.SendResponse(ctx, common.OK, listMap)
}

//Add
/**
* @Author: ethan.chen@cpgroup.cn
* @Date: 2021/10/11 11:29
* @Description: 添加评论
* @Params:
* @Return:
**/
func (c Comment) Add(ctx *gin.Context) {
	var comment dao.Comment
	claims, _ := c.tokenInfo(ctx)
	uid, _ := strconv.Atoi(claims.Uid)
	comment.UID = uint(uid)

	addQO := new(qo.AddCommentQO)
	util.JsonConvert(ctx, addQO)
	sn, _ := strconv.ParseInt(addQO.Sn, 10, 64)
	commentVO := vo.AddCommentVO{}

	if addQO.Content == "" {
		common.SendResponse(ctx, common.ErrParam, commentVO)
		return
	}

	//查询用户是否存在
	user := userCommonFunc.IUser(userCommonFunc.UserCommonFunc{}).FindUser(ctx, []int{int(comment.UID)}, "", "")
	if userInfo, ok := user[comment.UID]; !ok || int(userInfo.State) != cpgConst.ONE {
		common.SendResponse(ctx, common.ErrUserNotFound, commentVO)
		return
	}

	//查询文章是否存在
	articleMap := articleCommonFunc.IArticle(articleCommonFunc.ArticleCommonFunc{}).
		FindPublishedArticlesBySn(ctx, []int64{sn})
	if article, ok := articleMap[sn]; !ok || article.State != cpgConst.ONE {
		common.SendResponse(ctx, common.ErrArticleNotExisted, commentVO)
		return
	}

	//查询当前文章楼层数
	var floor int
	globalInit.Db.Model(&comment).
		Select("floor").
		Where("sn", sn).
		Order("floor desc").
		First(&floor)

	//插入评论
	//TODO 后续增加审核功能
	comment.State = cpgConst.ZERO
	comment.Sn = sn
	comment.Content = addQO.Content
	comment.Floor = floor + cpgConst.ONE

	cid, err := comment.CreateComment(ctx)
	if err != nil {
		common.SendResponse(ctx, err, commentVO)
		return
	}

	//查询添加的评论，返回评论Id
	commentVO.CommentId = cid

	//更新文章扩展表评论数
	err = articleCommonFunc.IArticle(articleCommonFunc.ArticleCommonFunc{}).
		UpdateArticleEx(ctx, sn, false, true, false, true)
	if err != nil {
		common.SendResponse(ctx, err, commentVO)
		return
	}

	common.SendResponse(ctx, common.OK, commentVO)
}

//Delete
/**
* @Author: ethan.chen@cpgroup.cn
* @Date: 2021/10/11 15:10
* @Description: 删除评论（及附属回复）
* @Params: DeleteCommentQO
* @Return:
**/
func (c Comment) Delete(ctx *gin.Context) {
	deleteQO := qo.DeleteCommentQO{}
	util.JsonConvert(ctx, &deleteQO)

	comment := model.Comment{}
	var commentReply []model.CommentReply

	//查询评论状态，已删除/不存在则直接返回
	comment = c.commentInfo(deleteQO.CommentId)
	if reflect.DeepEqual(model.Comment{}, comment) {
		e := common.ErrParam
		e.Message = "comment not exist or was deleted"
		common.SendResponse(ctx, e, "")
		return
	}

	//查询评论是否存在回复，存在则先删除回复
	globalInit.Db.Model(&model.CommentReply{}).
		Where("cid = ? and state = ?", comment.Cid, cpgConst.ONE).
		Find(&commentReply)

	if !reflect.DeepEqual([]model.CommentReply{}, commentReply) {
		//删除该Cid下所有回复
		err := dao.CommentReply{Cid: comment.Cid, State: cpgConst.THREE}.
			UpdateCommentReplyByCid(ctx)
		if err != nil {
			common.SendResponse(ctx, err, "")
			return
		}
	}

	//更新评论状态为删除
	err := dao.Comment{Cid: comment.Cid, Content: comment.Content, State: cpgConst.THREE}.
		UpdateComment(ctx)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}

	//文章扩展表文章评论数更改
	err = articleCommonFunc.IArticle(articleCommonFunc.ArticleCommonFunc{}).
		UpdateArticleEx(ctx, comment.Sn, false, true, false, false)
	if err != nil {
		e := common.ErrDatabase
		e.Message = err.Error()
		common.SendResponse(ctx, e, "")
		return
	}

	//点赞表更改评论点赞状态
	err = likeCommonFunc.LikeCommonFunc{}.
		UpdateZanSate(ctx, comment.Sn, cpgConst.ONE)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")
}

//AddReply
/**
* @Author: ethan.chen@cpgroup.cn
* @Date: 2021/10/11 15:10
* @Description: 评论回复
* @Params:
* @Return:
**/
func (c Comment) AddReply(ctx *gin.Context) {
	replyQO := qo.AddCommentReplyQO{}
	util.JsonConvert(ctx, &replyQO)
	replyVO := vo.AddCommentReplyVO{}
	token, _ := c.tokenInfo(ctx)
	uid, _ := strconv.Atoi(token.Uid)

	//查询评论状态（如果非上线状态则不允许进行回复）
	comment := model.Comment{}
	globalInit.Db.Model(model.Comment{}).
		Where("cid = ? and state = ?", replyQO.CommentId, cpgConst.ONE).Find(&comment)
	if reflect.DeepEqual(model.Comment{}, comment) {
		common.SendResponse(ctx, common.ErrParam, replyVO)
		return
	}

	//添加回复
	reply := dao.CommentReply{}
	reply.Cid = uint(replyQO.CommentId)
	reply.UID = uint(uid)
	reply.Content = replyQO.Content
	//TODO 后续增加审核功能
	reply.State = cpgConst.ZERO
	replyId, err := reply.CreateCommentReply(ctx)
	if err != nil {
		common.SendResponse(ctx, err, replyVO)
		return
	}
	common.SendResponse(ctx, common.OK, replyId)
}

//DeleteReply
/**
* @Author: ethan.chen@cpgroup.cn
* @Date: 2021/10/11 15:10
* @Description: 删除回复
* @Params: DeleteCommentReplyQO
* @Return:
**/
func (c Comment) DeleteReply(ctx *gin.Context) {
	deleteQO := qo.DeleteCommentReplyQO{}
	util.JsonConvert(ctx, &deleteQO)

	err := dao.CommentReply{Id: uint(deleteQO.Id), State: cpgConst.THREE}.DeleteCommentReplyById(ctx)
	if err != nil {
		e := common.ErrDatabase
		e.Message = err.Error()
		common.SendResponse(ctx, e, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")
}
