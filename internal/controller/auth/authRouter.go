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

	//查询所有权限
	authGroup.POST("/query/permissions", auth.AllPolicies)

	//查询所有用户组
	authGroup.POST("/query/roles", auth.AllRoles)

	// AddPermission 系统添加单个权限
	authGroup.POST("/add/permission", auth.AddPermission)

	//删除单个权限

	// AddRole 添加角色
	authGroup.POST("/add/role", auth.AddRole)

	// AddPermissionsForGroup 角色添加权限
	authGroup.POST("/role/add/permission", auth.AddPermissionsForRole)

	// AddUserIntoGroup 添加用户-用户组关联
	authGroup.POST("/group/add/user", auth.AddUserIntoRole)

	//删除用户组

	//用户移除用户组
}