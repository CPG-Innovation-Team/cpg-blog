package dao

import (
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/user/model"
	"github.com/gin-gonic/gin"
)

type UserDAO struct{}

func (u UserDAO) Create(ctx *gin.Context, user *model.User) (err error) {
	tx := globalInit.Db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u UserDAO) SelectByUid(ctx *gin.Context, param interface{}) (users []model.User) {
	globalInit.Db.Model(&model.User{}).Where("uid", param).Find(&users)
	return
}

func (u UserDAO) SelectByName(ctx *gin.Context, param interface{}) (users []model.User) {
	globalInit.Db.Model(&model.User{}).Where("username", param).Find(&users)
	return
}

func (u UserDAO) SelectByEmail(ctx *gin.Context, param interface{}) (users []model.User) {
	globalInit.Db.Model(&model.User{}).Where("email", param).Find(&users)
	return
}

func (u UserDAO) SelectByNameAndEmail(ctx *gin.Context, param *model.User) (users []model.User) {
	globalInit.Db.Where("username = ? or email = ?", &param.UserName, &param.Email).Find(&users)
	return
}

func (u UserDAO) UpdateUserInfo(ctx *gin.Context, param *model.User) (err error) {
	tx := globalInit.Db.Model(param).Updates(&param)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
