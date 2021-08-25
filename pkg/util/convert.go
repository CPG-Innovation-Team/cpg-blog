package util

import (
	"cpg-blog/global/common"
	"github.com/gin-gonic/gin"
)

func JsonConvert(ctx *gin.Context, obj interface{}) {
		if err := ctx.ShouldBindJSON(obj); err != nil {
			common.SendResponse(ctx, common.ErrBind, err.Error())
			return
		}
}
