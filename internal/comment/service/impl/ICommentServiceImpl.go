package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	"cpg-blog/global/globalInit"
	Article "cpg-blog/internal/article/service"
	"cpg-blog/internal/comment/model/dao"
	"cpg-blog/internal/comment/qo"
	User "cpg-blog/internal/user/service"
	"cpg-blog/middleware/jwt"
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Comment struct{}

var (
	UserService User.IUser
	ArticleService Article.IArticle
)

func (c Comment) List(ctx *gin.Context) {}

func (c Comment) Add(ctx *gin.Context) {
	var comment dao.CommentDao
	token :=ctx.Request.Header.Get("token")
	claims,_ := new(jwt.JWT).ParseToken(token)
	uid,_ := strconv.Atoi(claims.Uid)
	comment.UID = uint(uid)

	addQO := new(qo.AddCommentQO)
	util.JsonConvert(ctx, addQO)

	if addQO.Content == ""{
		common.SendResponse(ctx, common.ErrParam,"")
		return
	}

	//查询用户是否存在
	user:= UserService.FindUser(ctx,[]int{int(comment.UID)},"","")
	if userInfo,ok := user[comment.UID]; !ok || int(userInfo.State) != cpgConst.ONE{
		common.SendResponse(ctx,common.ErrUserNotFound,"")
		return
	}

	//查询文章是否存在
	articleMap := ArticleService.FindArticles(ctx,[]int64{addQO.Sn})
	if article,ok := articleMap[addQO.Sn];!ok||article.State != cpgConst.ONE{
		common.SendResponse(ctx,common.ErrArticleNotExisted,"")
		return
	}

	//查询当前文章楼层数
	var floor int
	globalInit.Db.Model(&comment).
		Select("floor").
		Where("sn",addQO.Sn).
		Order("floor desc").
		First(floor)

	//插入评论
	//TODO 后续增加审核功能
	comment.State = cpgConst.ONE
	comment.Sn = addQO.Sn
	comment.Content = addQO.Content
	if floor == cpgConst.ZERO{
		comment.Floor = floor
	}
	if floor > cpgConst.ZERO{
		comment.Floor = floor+cpgConst.ONE
	}
	err := comment.AddComment(ctx)
	if err != nil {
		common.SendResponse(ctx,err,"")
		return
	}
	common.SendResponse(ctx,common.OK,"")
}

func (c Comment) Delete(ctx *gin.Context) {}

func (c Comment) Reply(ctx *gin.Context) {}

func (c Comment) DeleteReply(ctx *gin.Context) {}
