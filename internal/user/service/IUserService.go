package service

import (
	"cpg-blog/internal/user/model"
	"github.com/gin-gonic/gin"
)

type IUser interface {
	// Login 登录接口
	Login(ctx *gin.Context)

	// Register 注册接口
	Register(ctx *gin.Context)

	// FindUser 服务间查询用户信息
	FindUser(ctx *gin.Context, uidList []int, name string, email string) (users map[uint]model.User)

	// Info 前端请求查询用户信息
	Info(ctx *gin.Context)

	// List 查询用户list
	List(ctx *gin.Context)

	//Modify 修改用户信息
	Modify(ctx *gin.Context)
}
