package service

import "github.com/gin-gonic/gin"

type IAuth interface {
	// AddPermission 系统添加单个权限
	AddPermission(ctx *gin.Context)

	// AddGroup 添加用户组
	AddGroup(ctx *gin.Context)

	// AddPermissionsForGroup 用户组添加权限
	AddPermissionsForGroup(ctx *gin.Context)

	// AddUserIntoGroup 添加用户-用户组关联
	AddUserIntoGroup(ctx *gin.Context)
}