package service

import "github.com/gin-gonic/gin"

type IAuth interface {
	// AllPolicies 查询所有权限
	AllPolicies(ctx *gin.Context)

	// AllGroups 查询所有用户组
	AllRoles(ctx *gin.Context)

	// AddPermission 系统添加单个权限
	AddPermission(ctx *gin.Context)

	// AddGroup 添加用户组
	AddRole(ctx *gin.Context)

	// AddPermissionsForGroup 用户组添加权限
	AddPermissionsForGroup(ctx *gin.Context)

	// AddUserIntoGroup 添加用户-用户组关联
	AddUserIntoGroup(ctx *gin.Context)
}