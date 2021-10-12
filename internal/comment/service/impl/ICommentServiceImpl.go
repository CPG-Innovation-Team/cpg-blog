package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	"cpg-blog/global/globalInit"
	Article "cpg-blog/internal/article/service"
	"cpg-blog/internal/comment/model"
	"cpg-blog/internal/comment/model/dao"
	"cpg-blog/internal/comment/qo"
	"cpg-blog/internal/comment/vo"
	Like "cpg-blog/internal/like/service"
	User "cpg-blog/internal/user/service"
	"cpg-blog/middleware/jwt"
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

type Comment struct{}

var (
	userService    User.IUser
	articleService Article.IArticle
	likeService    Like.ILike
)

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
		First(comment)
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

func (c Comment) List(ctx *gin.Context) {}

//Add
/**
* @Author: ethan.chen@cpgroup.cn
* @Date: 2021/10/11 11:29
* @Description: 添加评论
* @Params:
* @Return:
**/
func (c Comment) Add(ctx *gin.Context) {
	var comment dao.CommentDao
	claims, _ := c.tokenInfo(ctx)
	uid, _ := strconv.Atoi(claims.Uid)
	comment.UID = uint(uid)

	addQO := new(qo.AddCommentQO)
	util.JsonConvert(ctx, addQO)
	commentVO := vo.AddCommentVO{}

	if addQO.Content == "" {
		common.SendResponse(ctx, common.ErrParam, commentVO)
		return
	}

	//查询用户是否存在
	user := userService.FindUser(ctx, []int{int(comment.UID)}, "", "")
	if userInfo, ok := user[comment.UID]; !ok || int(userInfo.State) != cpgConst.ONE {
		common.SendResponse(ctx, common.ErrUserNotFound, commentVO)
		return
	}

	//查询文章是否存在
	articleMap := articleService.FindArticles(ctx, []int64{addQO.Sn})
	if article, ok := articleMap[addQO.Sn]; !ok || article.State != cpgConst.ONE {
		common.SendResponse(ctx, common.ErrArticleNotExisted, commentVO)
		return
	}

	//查询当前文章楼层数
	var floor int
	globalInit.Db.Model(&comment).
		Select("floor").
		Where("sn", addQO.Sn).
		Order("floor desc").
		First(floor)

	//插入评论
	//TODO 后续增加审核功能
	comment.State = cpgConst.ONE
	comment.Sn = addQO.Sn
	comment.Content = addQO.Content
	if floor == cpgConst.ZERO {
		comment.Floor = floor
	}
	if floor > cpgConst.ZERO {
		comment.Floor = floor + cpgConst.ONE
	}
	cid, err := comment.CreateComment(ctx)
	if err != nil {
		common.SendResponse(ctx, err, commentVO)
		return
	}

	//查询添加的评论，返回评论Id
	commentVO.CommentId = cid
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
	util.JsonConvert(ctx, deleteQO)

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
		Find(commentReply)

	if !reflect.DeepEqual([]model.CommentReply{}, commentReply) {
		//删除该Cid下所有回复
		err := dao.CommentReplyDao{Cid: comment.Cid, State: cpgConst.THREE}.UpdateCommentReply(ctx)
		if err != nil {
			common.SendResponse(ctx, err, "")
			return
		}
	}

	//更新评论状态为删除
	err := dao.CommentDao{Cid: comment.Cid, State: cpgConst.THREE}.UpdateComment(ctx)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}

	//文章扩展表文章评论数更改
	err = articleService.UpdateArticleEx(ctx, comment.Sn, false, true, false, false)
	if err != nil {
		e := common.ErrDatabase
		e.Message = err.Error()
		common.SendResponse(ctx, e, "")
		return
	}

	//点赞表更改评论点赞状态
	err = likeService.Update(ctx, comment.Sn, cpgConst.ONE)
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
	util.JsonConvert(ctx, replyQO)
	replyVO := vo.AddCommentReplyVO{}
	token, _ := c.tokenInfo(ctx)
	uid, _ := strconv.Atoi(token.Uid)

	//查询评论状态（如果非上线状态则不允许进行回复）
	comment := model.Comment{}
	globalInit.Db.Model(model.Comment{}).
		Where("id = ?, state = ?", replyQO.CommentId, cpgConst.ZERO).Find(&comment)
	if reflect.DeepEqual(model.Comment{}, comment) {
		common.SendResponse(ctx, common.ErrParam, replyVO)
	}

	//添加回复
	reply := dao.CommentReplyDao{}
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
* @Params:
* @Return:
**/
func (c Comment) DeleteReply(ctx *gin.Context) {}
