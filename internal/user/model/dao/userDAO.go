package dao

import (
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/user/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserDAO struct {
	UId   []int
	Name  string
	Email string
}

var Db = &(globalInit.Db)

func (u UserDAO) Create(ctx *gin.Context, user *model.User) (err error) {
	tx := globalInit.Transaction()
	err = func(db *gorm.DB) error {
		if tx.Error != nil {
			return tx.Error
		}
		tx.Create(user)
		if tx.Error != nil {
			tx.Rollback()
			return tx.Error
		}
		return tx.Commit().Error
	}(tx)
	return err
}

func (u *UserDAO) GetUser(ctx *gin.Context) (users *[]model.User) {
	//tx := (*Db).Session(&gorm.Session{})
	tx := (*Db).WithContext(ctx)
	if len(u.UId) == 1 {
		tx = tx.Where("uid", u.UId[0])
	}
	if len(u.UId) > 1 {
		tx = tx.Where("uid In ?", u.UId)
	}
	if u.Name != "" {
		tx = tx.Where("username", u.Name)
	}
	if u.Email != "" {
		tx = tx.Where("email", u.Email)
	}
	tx.Model(&model.User{}).Find(&users)
	return
}

func (u UserDAO) SelectByUid(ctx *gin.Context, param interface{}) (users []model.User) {
	(*Db).Model(&model.User{}).Where("uid", param).Find(&users)
	return
}

func (u UserDAO) SelectByName(ctx *gin.Context, param interface{}) (users []model.User) {
	(*Db).Model(&model.User{}).Where("username", param).Find(&users)
	return
}

func (u UserDAO) SelectByEmail(ctx *gin.Context, param interface{}) (users []model.User) {
	(*Db).Model(&model.User{}).Where("email", param).Find(&users)
	return
}

func (u UserDAO) SelectByNameAndEmail(ctx *gin.Context, param *model.User) (users []model.User) {
	(*Db).Where("username = ? or email = ?", &param.UserName, &param.Email).Find(&users)
	return
}

func (u UserDAO) UpdateUserInfo(ctx *gin.Context, param *model.User) (err error) {
	//tx := (*Db).Model(param).Updates(&param)
	//if tx.Error != nil {
	//	return tx.Error
	//}
	//return nil

	tx := globalInit.Transaction()
	err = func(db *gorm.DB) error {
		if tx.Error != nil {
			return tx.Error
		}
		tx.Model(param).Updates(&param)
		if tx.Error != nil {
			tx.Rollback()
			return tx.Error
		}
		return tx.Commit().Error
	}(tx)
	return err
}

//UpdateUserAvatar 更新用户头像
func (u UserDAO) UpdateUserAvatar(ctx *gin.Context, param *model.User) (err error) {
	tx := globalInit.Transaction()
	err = func(db *gorm.DB) error {
		if tx.Error != nil {
			return tx.Error
		}
		tx.Model(param).Select("avatar").Updates(&param)
		if tx.Error != nil {
			tx.Rollback()
			return tx.Error
		}
		return tx.Commit().Error
	}(tx)
	return err
}
