package main

import (
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/controller"
	"cpg-blog/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	//编译信息，默认unknown
	gitCommitLog = "unknown"
	buildTime    = "unknown"
	gitRelease   = "unknown"
)

func init() {
	globalInit.ViperInit()
	globalInit.DbInit()
	globalInit.RedisInit()
	globalInit.App.SetFrameMode(gin.ReleaseMode)
	globalInit.App.FillBuildInfo(gitCommitLog, buildTime, gitRelease)
	globalInit.App.SetLog(viper.GetBool("log.isStdout"))

}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// header
	r.Use(middleware.NoCache)
	r.Use(middleware.Secure)
	r.Use(middleware.Options)

	//TODO 服务访问权限，通过OAuth服务实现

	// 后端路由组
	//登录等特殊页
	special := r.Group("")
	controller.RegisterSpecialRoutes(special)

	//门户页
	portal := r.Group("")
	portal.Use(middleware.JwtAuth)
	controller.RegisterPortalRoutes(portal)

	//管理页
	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.JwtAuth)
	adminGroup.Use(middleware.PermissionAuth)
	controller.RegisterRoutes(adminGroup)

	//TODO 获取CA证书，并自动更新
	s := &http.Server{
		Addr:           viper.GetString("http.port"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//err := s.ListenAndServe()
	var err error
	if *globalInit.LocalDebug{
		err = s.ListenAndServe()
	} else {
		err = s.ListenAndServeTLS("/etc/ssl/certs/cpgroup.top_bundle.crt", "/etc/ssh/cpgroup.top.key")
	}

	if err != nil {
		return
	}
	// err := r.Run(viper.GetString("http.port"))
}
