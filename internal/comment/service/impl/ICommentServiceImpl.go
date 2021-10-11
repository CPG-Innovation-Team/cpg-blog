package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	"cpg-blog/global/globalInit"
	Article "cpg-blog/internal/article/service"
	"cpg-blog/internal/comment/model"
	"cpg-blog/internal/comment/model/dao"
	"cpg-blog/internal/comment/qo"
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
	token := ctx.Request.Header.Get("token")
	claims, _ := new(jwt.JWT).ParseToken(token)
	uid, _ := strconv.Atoi(claims.Uid)
	comment.UID = uint(uid)

	addQO := new(qo.AddCommentQO)
	util.JsonConvert(ctx, addQO)

	if addQO.Content == "" {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}

	//查询用户是否存在
	user := userService.FindUser(ctx, []int{int(comment.UID)}, "", "")
	if userInfo, ok := user[comment.UID]; !ok || int(userInfo.State) != cpgConst.ONE {
		common.SendResponse(ctx, common.ErrUserNotFound, "")
		return
	}

	//查询文章是否存在
	articleMap := articleService.FindArticles(ctx, []int64{addQO.Sn})
	if article, ok := articleMap[addQO.Sn]; !ok || article.State != cpgConst.ONE {
		common.SendResponse(ctx, common.ErrArticleNotExisted, "")
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
	err := comment.CreateComment(ctx)
	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")
}

//Delete
/**
* @Author: ethan.chen@cpgroup.cn
* @Date: 2021/10/11 15:10
* @Description: 删除评论（及附属回复）
* @Params:
* @Return:
**/
func (c Comment) Delete(ctx *gin.Context) {
	deleteQO := qo.DeleteCommentQO{}
	util.JsonConvert(ctx, deleteQO)

	comment := model.Comment{}
	var commentReply []model.CommentReply

	//查询评论状态，已删除/不存在则直接返回
	globalInit.Db.Model(&model.Comment{}).
		Where("cid = ? and state = ?", deleteQO.CommentId, cpgConst.ONE).
		First(comment)
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
func (c Comment) AddReply(ctx *gin.Context) {}

//DeleteReply
/**
* @Author: ethan.chen@cpgroup.cn
* @Date: 2021/10/11 15:10
* @Description: 删除回复
* @Params:
* @Return:
**/
func (c Comment) DeleteReply(ctx *gin.Context) {}
