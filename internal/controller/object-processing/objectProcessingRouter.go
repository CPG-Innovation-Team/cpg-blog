package object_processing

import (
	"cpg-blog/internal/object-processing/service/impl"
	"github.com/gin-gonic/gin"
)

/**
  @author: chenxi@cpgroup.cn
  @date:2022/5/17
  @description:
**/

type Controller struct{}

var processor = &impl.ObjectProcessing{}

func (c Controller) RegisterRoute(g *gin.RouterGroup) {
	//query := g.Group("/object/query")
	update := g.Group("/object")

	//上传头像
	update.POST("/upload/avatar", processor.UploadAvatar)

	//更新头像
	update.POST("/update/avatar", processor.UpdateAvatar)

	//上传文件
	update.POST("/upload", processor.Upload)

	//更新文件
	update.POST("/update", processor.Update)

	//删除文件
	update.POST("/delete", processor.DeleteFile)

}
