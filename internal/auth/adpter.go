package auth

import (
	"cpg-blog/global/common"
	"cpg-blog/global/globalInit"
	"github.com/casbin/casbin/v2"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
)

func GetEnforcer() (e *casbin.Enforcer, err error) {
	a, _ := gormAdapter.NewAdapterByDB(globalInit.Db)

	e, _ = casbin.NewEnforcer(globalInit.App.RootDir+"/internal/auth/rbac_model.conf", a)
	err = e.LoadPolicy()
	e.EnableLog(true)

	// Save the policy back to DB.
	defer func(e *casbin.Enforcer) {
		err := e.SavePolicy()
		if err != nil {
			return
		}
	}(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func GetE(ctx *gin.Context) (e *casbin.Enforcer, err error) {
	e, err = GetEnforcer()
	if err != nil {
		common.SendResponse(ctx, common.ErrLoadPolicy, err)
		return
	}
	return e, nil
}