package service

import "github.com/gin-gonic/gin"

type User interface {
	// Login 登录接口
	Login(ctx *gin.Context)

	// Register 注册接口
	Register(ctx *gin.Context)

	// Info 查询用户信息
	Info(ctx *gin.Context)

	//Modify 修改用户信息
	Modify(ctx *gin.Context)
}