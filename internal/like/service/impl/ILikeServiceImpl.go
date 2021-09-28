package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	articleDao "cpg-blog/internal/article/model/dao"
	"cpg-blog/internal/like/model/dao"
	"cpg-blog/internal/like/qo"
	"cpg-blog/middleware/jwt"
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
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

func (l Like) Like(ctx *gin.Context) {
	likeQO := new(qo.LikeQO)
	util.JsonConvert(ctx, likeQO)

	if (likeQO.Sn == zero64 && likeQO.CommentId == zero) ||
		(likeQO.Sn != zero64 && likeQO.CommentId != zero) {
		common.SendResponse(ctx, common.ErrParam, "")
	}
	token, _ := tokenInfo(ctx)
	//点赞用户uid
	uid, _ := strconv.Atoi(token.Uid)
	var err error

	//点赞文章
	if likeQO.Sn != 0 {
		err = likeArticle(likeQO.Sn, uid)
		if err != nil {
			e := common.ErrDatabase
			e.Message = err.Error()
			common.SendResponse(ctx, e, "")
			return
		}
	} else {
		//点赞评论
		err = likeComment(likeQO.CommentId, uid)
	}
	common.SendResponse(ctx, common.OK, "")
}

func likeArticle(sn int64, uid int) (err error) {
	//查询文章是否存在
	article := articleDao.ArticleDAO{}.SelectArticleBySn(sn)
	if article == nil {
		return common.ErrParam
	}

	err = dao.LikeDAO{}.Creat(uid, cpgConst.ZERO, sn)
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

func likeComment(commentId int, uid int) (err error) {
	//TODO 查询comment是否存在

	return dao.LikeDAO{}.Creat(uid, cpgConst.ONE, int64(commentId))
}
