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

	//查询所有角色
	authGroup.POST("/query/roles", auth.AllRoles)

	// AddPermission 系统添加单个权限
	authGroup.POST("/add/permission", auth.AddPermission)

	//删除单个权限,且解除与角色的关联
	authGroup.POST("/delete/permission", auth.DeletePermission)

	// AddRole 添加角色
	authGroup.POST("/add/role", auth.AddRole)

	// AddPermissionsForGroup 角色添加权限
	authGroup.POST("/role/add/permission", auth.AddPermissionsForRole)

	// RemovePermissionsFromRole 角色删除权限
	authGroup.POST("/role/remove/permission", auth.RemovePermissionsFromRole)

	// AddUserIntoGroup 添加用户-用户组关联
	authGroup.POST("/role/add/user", auth.AddUserIntoRole)

	//用户移除角色
	authGroup.POST("/role/remove/user", auth.RoleRemoveUser)

	//删除角色，且解除角色与权限关联
	authGroup.POST("/delete/role", auth.DeleteRole)

	//查询用户所有角色
	authGroup.POST("/query/user/roles", auth.GetUserRoles)

	//用户批量添加角色
	authGroup.POST("/user/add/roles", auth.UserAddRolesInBatches)

	//用户批量删除角色
	authGroup.POST("/user/delete/roles", auth.UserDeleteRolesInBatches)
}
