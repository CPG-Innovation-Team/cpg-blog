package userCommonFunc

import (
	"cpg-blog/internal/user/model"
	"cpg-blog/internal/user/model/dao"
	"github.com/gin-gonic/gin"
)

/**
  @author: ethan.chen@cpgroup.cn
  @date:2021/10/15
  @description:
**/

type IUser interface {
	// FindUser 服务间查询用户信息
	FindUser(ctx *gin.Context, uidList []int, name string, email string) (users map[uint]model.User)

	//UpdateUserAvatar 更新用户头像
	UpdateUserAvatar(ctx *gin.Context, uid int, avatar string) (err error)
}

type UserCommonFunc struct{}

func (c UserCommonFunc) Get() *UserCommonFunc {
	return new(UserCommonFunc)
}

func (c UserCommonFunc) FindUser(ctx *gin.Context, uidList []int, name string, email string) (users map[uint]model.User) {
	findQO := &dao.UserDAO{
		UId:   uidList,
		Name:  name,
		Email: email,
	}
	userList := findQO.GetUser(ctx)
	users = map[uint]model.User{}
	for _, v := range *userList {
		users[v.UID] = v
	}
	return users
}

func (c UserCommonFunc) UpdateUserAvatar(ctx *gin.Context, uid int, avatar string) (err error) {
	updateQO := &model.User{
		UID:    uint(uid),
		Avatar: avatar,
	}

	return dao.UserDAO{}.UpdateUserAvatar(ctx, updateQO)
}
