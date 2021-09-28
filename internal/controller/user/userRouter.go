package user

import (
	"cpg-blog/internal/user/service/impl"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

var user = &impl.Users{}

func (u Controller) RegisterSpecialRoute(g *gin.RouterGroup){
	//登录
	g.POST("/login", user.Login)
	//注册
	g.POST("/register", user.Register)
}

// RegisterRoute 添加user服务路由
func (u Controller) RegisterRoute(g *gin.RouterGroup) {

	g.POST("/logout", user.Logout)

	query := g.Group("/user/query")
	update := g.Group("/user/update")

	//查询用户信息
	query.POST("/info", user.Info)

	//查询所有用户
	query.POST("/list", user.List)

	//修改用户信息
	update.POST("/info", user.Modify)
}
