package service

import "github.com/gin-gonic/gin"

/**
  @author: chenxi@cpgroup.cn
  @date:2022/5/12
  @description:
**/

// IObjectProcessing Service interface for object processing
type IObjectProcessing interface {
	UploadAvatar(ctx *gin.Context)
	UpdateAvatar(ctx *gin.Context)

	Upload(ctx *gin.Context)
	Update(ctx *gin.Context)

	DeleteFile(ctx *gin.Context)
}
