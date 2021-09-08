package util

import (
	"cpg-blog/global/common"
	"github.com/gin-gonic/gin"
)

func JsonConvert(ctx *gin.Context, obj interface{}) {
	err := ctx.ShouldBindJSON(obj)
	if err != nil {
		common.SendResponse(ctx, common.ErrBind, err.Error())
		panic(common.ErrBind)
	}
}
