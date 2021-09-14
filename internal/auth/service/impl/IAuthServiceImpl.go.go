package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	"cpg-blog/internal/auth"
	"cpg-blog/internal/auth/qo"
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
)

type Auth struct{}

//[][]string转换成map[string][]string
func sliceToMap(s [][]string) map[string][]string {
	m := map[string][]string{}
	if len(s) == 0 {
		return m
	}
	if len(s[0]) == 2 {
		for _, v := range s {
			sv, ok := m[strings.TrimPrefix(v[1], cpgConst.RolePrefix)]
			if ok {
				sv = append(sv, v[0])
			}
			m[strings.TrimPrefix(v[1], cpgConst.RolePrefix)] = []string{v[0]}
		}
	} else if len(s[0]) == 3 {
		for _, v := range s {
			m[v[0]] = []string{v[1]}
		}
	}
	return m
}

// AllPolicies 查询所有权限
func (a Auth) AllPolicies(ctx *gin.Context) {
	e, _ := auth.GetE(ctx)
	policies := e.GetAllRoles()
	common.SendResponse(ctx, common.OK, policies)
}

// AllRoles 查询所有角色及其权限
func (a Auth) AllRoles(ctx *gin.Context) {
	e, _ := auth.GetE(ctx)
	roleAndPermission := make(map[string]map[string]string)

	//角色-权限关系
	roleAndPermissionRelationship := sliceToMap(e.GetNamedGroupingPolicy("g2"))

	//权限
	permission := sliceToMap(e.GetNamedPolicy("p"))

	for _,v:= range roleAndPermissionRelationship{
		for _, v1:=range v{
			v1 = permission[v1][0]
		}
	}

	//角色继承关系
	inheritanceRelationship := e.GetNamedGroupingPolicy("g")
	for _, v := range inheritanceRelationship {
		//存在继承关系
		if strings.Contains(v[0], cpgConst.RolePrefix) {

		}
	}

	//enforce, err := e.GetImplicitUsersForPermission("权限1")
	//if err != nil {
	//	return
	//}
	log.Println(inheritanceRelationship)
	log.Println(roleAndPermissionRelationship)
	log.Println(permission)

	log.Println(roleAndPermission)
	//var groups map[string]string
	//for _,v:=range group{
	//	if strings.Contains(v[0],"group::"){
	//		groups = append(groups,v[0])
	//	}
	//}
	common.SendResponse(ctx, common.OK, e.GetNamedGroupingPolicy("g"))
}

// AddPermission 系统添加单个权限
func (a Auth) AddPermission(ctx *gin.Context) {
	p := qo.GetNewPermission()
	e, _ := auth.GetE(ctx)
	util.JsonConvert(ctx, &p)

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

// AddRole 添加角色
func (a Auth) AddRole(ctx *gin.Context) {
	name := new(qo.AddGroupQO)
	util.JsonConvert(ctx, name)
	e, _ := auth.GetE(ctx)
	result, err := e.AddNamedGroupingPolicy("g", cpgConst.RolePrefix+name.RName, cpgConst.AdminRole)
	if err != nil {
		common.SendResponse(ctx, common.ErrDatabase, err)
		return
	}
	if !result {
		common.SendResponse(ctx, common.ErrRoleExisted, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")
}

// AddPermissionsForRole 角色添加权限
func (a Auth) AddPermissionsForRole(ctx *gin.Context) {
	e, _ := auth.GetE(ctx)
	gap := new(qo.GroupAddPermissionQO)
	util.JsonConvert(ctx, gap)

	//校验角色
	hasGroup := e.GetFilteredNamedGroupingPolicy("g", 0, gap.RName)
	if len(hasGroup) == 0 {
		common.SendResponse(ctx, common.ErrGroupNotExisted, "")
		return
	}
	//根据PName查询策略字段
	hasPermission := e.GetFilteredNamedPolicy("p", 0, gap.PName)
	if len(hasPermission) > 0 {
		hasPolicy, err := e.AddNamedGroupingPolicy("g2", gap.PName, gap.RName)

		if err != nil {
			common.SendResponse(ctx, common.ErrDatabase, "添加失败"+err.Error())
			return
		}
		if !hasPolicy {
			common.SendResponse(ctx, common.OK, "该角色已存在该权限")
			return
		}
		common.SendResponse(ctx, nil, "添加成功")
		return
	} else {
		common.SendResponse(ctx, common.ErrPermissionNotExisted, "")
		return
	}
}

// AddUserIntoRole 添加用户-角色关联
func (a Auth) AddUserIntoRole(ctx *gin.Context) {
	userIntoGroup := new(qo.AddUserIntoRoleQO)
	util.JsonConvert(ctx, userIntoGroup)
	uid := strconv.Itoa(userIntoGroup.Uid)
	e, _ := auth.GetE(ctx)
	result, err := e.AddRoleForUser(cpgConst.UserPrefix+uid, userIntoGroup.RName)
	if err != nil {
		common.SendResponse(ctx, common.ErrDatabase, err)
		return
	}
	if !result {
		common.SendResponse(ctx, common.ErrUserExistedInRole, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")
}
