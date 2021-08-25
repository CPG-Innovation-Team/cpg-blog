package controller

import (
	"cpg-blog/internal/controller/user"
	"github.com/gin-gonic/gin"
)

func RegisterSpecialRoutes(g *gin.RouterGroup)  {
	new(user.Controller).RegisterSpecialRoute(g)
}
// RegisterRoutes 统一注册路由
func RegisterRoutes(g *gin.RouterGroup)  {
	new(user.Controller).RegisterRoute(g)
}
