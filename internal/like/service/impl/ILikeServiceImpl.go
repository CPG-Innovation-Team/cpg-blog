package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	articleDao "cpg-blog/internal/article/model/dao"
	"cpg-blog/internal/like/model"
	"cpg-blog/internal/like/model/dao"
	"cpg-blog/internal/like/qo"
	"cpg-blog/middleware/jwt"
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

var (
	zero   = cpgConst.ZERO
	zero64 = int64(zero)
)

type Like struct{}

func tokenInfo(ctx *gin.Context) (Info *jwt.CustomClaims, err error) {
	return jwt.NewJWT().ParseToken(ctx.Request.Header.Get("token"))
}

func isLike(ctx *gin.Context, isCancelLike bool) (e error) {
	likeQO := new(qo.LikeQO)
	util.JsonConvert(ctx, likeQO)

	if (likeQO.Sn == zero64 && likeQO.CommentId == zero) ||
		(likeQO.Sn != zero64 && likeQO.CommentId != zero) {
		common.SendResponse(ctx, common.ErrParam, "")
	}
	token, _ := tokenInfo(ctx)
	// 点赞/取消点赞用户uid
	uid, _ := strconv.Atoi(token.Uid)

	var err error

	// 点赞/取消点赞文章
	if likeQO.Sn != zero64 {
		if !isCancelLike{
			err = likeArticle(likeQO.Sn, uid)
		}else if isCancelLike {
			err = cancelLikeArticle(likeQO.Sn, uid)
		}

		if err != nil {
			e := common.ErrDatabase
			e.Message = err.Error()
			return e
		}
	} else if likeQO.CommentId != zero {
		// 点赞/取消点赞评论
		if !isCancelLike{
			err = likeComment(likeQO.CommentId, uid)
		}else if isCancelLike {
			err = cancelLikeComment(likeQO.CommentId, uid)
		}

		if err != nil {
			e := common.ErrDatabase
			e.Message = err.Error()
			return e
		}
	}
	return  common.OK
}

func likeArticle(sn int64, uid int) (err error) {
	//查询文章是否存在,且已上线
	article := articleDao.ArticleDAO{}.SelectArticleBySn(sn)
	if article == nil {
		e := common.ErrParam
		e.Message = "Invalid Param Or Article State is not Published."
		return e
	}

	//点赞表判断是否存在来增加/修改记录
	err = dao.LikeDAO{}.CreatOrUpdate(uid, cpgConst.ZERO, sn,false)
	if err != nil && err != common.OK {
		return err
	}

	if err == common.OK {
		return nil
	}

	//文章扩展表更新
	err = articleDao.ArticleDAO{}.UpdateArticleEx(sn, false, false, true, true)
	return
}

func cancelLikeArticle(sn int64, uid int) (err error) {
	//查询文章是否存在,且已上线
	article := articleDao.ArticleDAO{}.SelectArticleBySn(sn)
	if article == nil {
		e := common.ErrParam
		e.Message = "Invalid Param Or Article State Is Not Published."
		return e
	}

	zanRecord := dao.LikeDAO{}.SelectZan(uid, cpgConst.ZERO, sn)
	if reflect.DeepEqual(zanRecord,model.Zan{}){
		e := common.ErrDatabase
		e.Message = "Like Record Does Not Exist."
		return e
	}

	//点赞表判断是否存在来增加/修改记录
	err = dao.LikeDAO{}.CreatOrUpdate(uid, cpgConst.ZERO, sn,true)
	if err != nil && err != common.OK {
		return err
	}

	if err == common.OK {
		return nil
	}

	//文章扩展表更新
	err = articleDao.ArticleDAO{}.UpdateArticleEx(sn, false, false, true, false)
	return
}

func likeComment(commentId int, uid int) (err error) {
	//TODO 查询comment是否存在
	//TODO 点赞表增加记录
	//TODO 评论表增加点赞数
	return dao.LikeDAO{}.CreatOrUpdate(uid, cpgConst.ONE, int64(commentId),false)
}

func cancelLikeComment(commentId int, uid int) (err error) {
	//TODO 查询comment是否存在
	//TODO 点赞表更改记录
	//TODO 评论表减少点赞数
	return dao.LikeDAO{}.CreatOrUpdate(uid, cpgConst.ONE, int64(commentId),true)
}


func (l Like) Like(ctx *gin.Context) {
	//likeQO := new(qo.LikeQO)
	//util.JsonConvert(ctx, likeQO)
	//
	//if (likeQO.Sn == zero64 && likeQO.CommentId == zero) ||
	//	(likeQO.Sn != zero64 && likeQO.CommentId != zero) {
	//	common.SendResponse(ctx, common.ErrParam, "")
	//}
	//token, _ := tokenInfo(ctx)
	////点赞用户uid
	//uid, _ := strconv.Atoi(token.Uid)
	//var err error
	//
	////点赞文章
	//if likeQO.Sn != 0 {
	//	err = likeArticle(likeQO.Sn, uid)
	//	if err != nil {
	//		e := common.ErrDatabase
	//		e.Message = err.Error()
	//		common.SendResponse(ctx, e, "")
	//		return
	//	}
	//} else {
	//	//点赞评论
	//	err = likeComment(likeQO.CommentId, uid)
	//}
	//common.SendResponse(ctx, common.OK, "")
	common.SendResponse(ctx,isLike(ctx,false),"")
}

func (l Like) CancelLike(ctx *gin.Context)  {
	common.SendResponse(ctx,isLike(ctx,true),"")
}