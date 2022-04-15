package service

import "github.com/gin-gonic/gin"

type IAuth interface {
	// AllPolicies 查询所有权限
	AllPolicies(ctx *gin.Context)

	// AllRoles 查询所有角色及其权限
	AllRoles(ctx *gin.Context)

	// AddPermission 系统添加单个权限
	AddPermission(ctx *gin.Context)

	// AddRole 添加角色
	AddRole(ctx *gin.Context)

	// AddPermissionsForRole 角色添加权限
	AddPermissionsForRole(ctx *gin.Context)

	// RemovePermissionsFromRole 角色移除权限
	RemovePermissionsFromRole(ctx *gin.Context)

	// DeletePermission 移除权限，且解除权限-角色关联
	DeletePermission(ctx *gin.Context)

	// DeleteRole 删除角色，且解除角色与权限关联及角色与用户关联
	DeleteRole(ctx *gin.Context)

	//GetUserRoles 查询用户角色
	GetUserRoles(ctx *gin.Context)

	// AddUserIntoRole 添加用户-角色关联
	AddUserIntoRole(ctx *gin.Context)

	// UserAddRolesInBatches 用户批量添加角色
	UserAddRolesInBatches(ctx *gin.Context)

	// RoleRemoveUser 用户移除角色
	RoleRemoveUser(ctx *gin.Context)

	// UserDeleteRolesInBatches 批量删除用户的角色
	UserDeleteRolesInBatches(ctx *gin.Context)
}
