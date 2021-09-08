package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/internal/auth"
	"cpg-blog/internal/auth/qo"
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
)

type Auth struct{}

// AddPermission 系统添加单个权限
func (a Auth) AddPermission(ctx *gin.Context) {
	p := qo.GetNewPermission()
	e, _ := auth.GetE(ctx)
	util.JsonConvert(ctx, p)

	ok, err := e.AddPolicy(p.Name, p.Uri, p.Operate)
	if err != nil {
		common.SendResponse(ctx, common.ErrDatabase, err)
		return
	}
	//校验数据库是否存在该条权限规则
	if !ok {
		common.SendResponse(ctx, common.OK, "数据库存在该接口权限规则:"+p.Uri)
		return
	}
	common.SendResponse(ctx, common.OK, "接口权限添加成功！")
}

// AddGroup 添加用户组 TODO
func (a Auth) AddGroup(ctx *gin.Context)  {

}

// AddPermissionsForGroup 用户组添加权限
func (a Auth) AddPermissionsForGroup(ctx *gin.Context) {
	e, _ := auth.GetE(ctx)
	gap := new(qo.GroupAddPermission)
	util.JsonConvert(ctx, gap)

	//根据PName查询策略字段
	hasPermission := e.GetFilteredNamedPolicy("p", 0, gap.PName)
	if len(hasPermission) > 0 {
		hasPolicy, err := e.AddGroupingPolicy(gap.GName, gap.PName)

		if err != nil {
			common.SendResponse(ctx, common.ErrDatabase, "添加失败"+err.Error())
			return
		}
		if !hasPolicy {
			common.SendResponse(ctx, common.OK, "该群组已存在该权限")
			return
		}
		common.SendResponse(ctx, nil, "添加成功")
		return
	} else {
		common.SendResponse(ctx, common.ErrPermissionNotExisted, "")
		return
	}
}

// AddUserIntoGroup 添加用户-用户组关联 TODO
func (a Auth) AddUserIntoGroup(ctx *gin.Context)  {

}
