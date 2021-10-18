package likeCommonFunc

import (
	"cpg-blog/global/common"
	"cpg-blog/internal/like/model/dao"
	"github.com/gin-gonic/gin"
)

/**
  @author: ethan.chen@cpgroup.cn
  @date:2021/10/15
  @description:
**/

type ILike interface {
	// UpdateZanSate 服务间更新点赞表state(0未删除，1删除),objId为文章Id或评论Id
	UpdateZanSate(ctx *gin.Context, objId int64, state int) error
}

type LikeCommonFunc struct {}

func (c LikeCommonFunc)UpdateZanSate(ctx *gin.Context, objId int64, state int) (err error ) {
	err = dao.LikeDAO{}.UpdateZanSate(ctx, objId, state)
	if err == nil {
		return common.OK
	}
	e := common.ErrDatabase
	e.Message = err.Error()
	return e
}
