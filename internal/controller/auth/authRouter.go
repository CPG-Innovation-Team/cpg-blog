package auth

import (
	"cpg-blog/internal/auth/service/impl"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

var auth = &impl.Auth{}

// RegisterRoute 添加article服务路由
func (u Controller) RegisterRoute(g *gin.RouterGroup) {
	authGroup := g.Group("/auth")

	// AddPermission 系统添加单个权限
	authGroup.POST("/add/permission", auth.AddPermission)

	// AddGroup 添加用户组
	authGroup.POST("/add/group", auth.AddGroup)

	// AddPermissionsForGroup 用户组添加权限
	authGroup.POST("/group/add/permission", auth.AddPermissionsForGroup)

	// AddUserIntoGroup 添加用户-用户组关联
	authGroup.POST("/add/permission", auth.AddUserIntoGroup)
}