package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	"cpg-blog/internal/auth"
	"cpg-blog/internal/auth/qo"
	"cpg-blog/internal/auth/vo"
	"cpg-blog/pkg/commonFunc/userCommonFunc"
	"cpg-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
)

type Auth struct{}

// AllPolicies 查询所有权限
func (a Auth) AllPolicies(ctx *gin.Context) {
	e, _ := auth.GetE(ctx)
	permission := e.GetNamedGroupingPolicy("g2")
	log.Println(permission)
	permissionMap := map[string]string{}
	for _, v := range permission {
		_, ok := permissionMap[v[1]]
		if !ok {
			permissionMap[v[1]] = v[0]
		}
	}
	common.SendResponse(ctx, common.OK, permissionMap)
	return
}

// AllRoles 查询所有角色及其权限
func (a Auth) AllRoles(ctx *gin.Context) {
	e, _ := auth.GetE(ctx)

	//角色
	role := e.GetNamedGroupingPolicy("g")
	roleMap := map[string][]string{}
	for _, v := range role {
		_, ok := roleMap[strings.TrimPrefix(v[1], cpgConst.RolePrefix)]
		if !ok {
			roleMap[strings.TrimPrefix(v[1], cpgConst.RolePrefix)] = []string{}
		}
	}
	log.Println("角色map:", roleMap)

	//权限
	permission := e.GetNamedGroupingPolicy("g2")
	permissionMap := map[string]string{}
	for _, v := range permission {
		_, ok := permissionMap[v[1]]
		if !ok {
			permissionMap[v[1]] = v[0]
		}
	}
	log.Println("权限map:", permissionMap)

	//角色-权限表关系
	roleAndPermission := e.GetPolicy()
	log.Println(roleAndPermission)
	for _, v := range roleAndPermission {
		rName := strings.TrimPrefix(v[0], cpgConst.RolePrefix)
		roleMap[rName] = append(roleMap[rName], permissionMap[v[1]])
	}
	log.Println("角色与权限关系map:", roleMap)

	common.SendResponse(ctx, common.OK, roleMap)
	return
}

// AddPermission 系统添加单个权限
func (a Auth) AddPermission(ctx *gin.Context) {
	p := qo.PermissionQO{}
	e, _ := auth.GetE(ctx)
	util.JsonConvert(ctx, &p)

	ok, err := e.AddNamedGroupingPolicy("g2", p.Uri, p.PName)
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
	return
}

// AddRole 添加角色
func (a Auth) AddRole(ctx *gin.Context) {
	name := new(qo.AddRoleQO)
	util.JsonConvert(ctx, name)
	e, _ := auth.GetE(ctx)
	result, err := e.AddNamedGroupingPolicy("g", "", cpgConst.RolePrefix+name.RName)
	if err != nil {
		common.SendResponse(ctx, common.ErrDatabase, err)
		return
	}
	if !result {
		common.SendResponse(ctx, common.ErrRoleExisted, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")
	return
}

// AddPermissionsForRole 角色添加权限
func (a Auth) AddPermissionsForRole(ctx *gin.Context) {
	e, _ := auth.GetE(ctx)
	gap := new(qo.GroupAddPermissionQO)
	util.JsonConvert(ctx, gap)

	if len(gap.PName) < 1 || gap.RName == "" {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}

	//policyList := strings.Join(gap.PName,", ")
	//log.Println(policyList)
	//校验角色
	hasGroup := e.GetFilteredNamedGroupingPolicy("g", 1, cpgConst.RolePrefix+gap.RName)
	if len(hasGroup) == 0 {
		common.SendResponse(ctx, common.ErrRoleNotExisted, "")
		return
	}
	//根据PName查询策略字段
	//hasPermission := e.GetFilteredNamedGroupingPolicy("g2", 1, policyList)
	//log.Println(hasPermission)

	var failureString string
	for _, v := range gap.PName {
		hasPolicy, _ := e.AddPolicy(cpgConst.RolePrefix+gap.RName, v, cpgConst.Operate)
		if !hasPolicy {
			var build strings.Builder
			build.WriteString(failureString)
			build.WriteString(v)
			build.WriteString(" ")
			failureString = build.String()
		}
	}
	if failureString != "" {
		common.SendResponse(ctx, common.ErrAddPermission, "添加失败的权限为："+failureString)
		return
	}
	common.SendResponse(ctx, common.OK, "权限添加成功")
	return
	//if len(hasPermission) > 0 {
	//	hasPolicy, err := e.AddPolicy(cpgConst.RolePrefix+gap.RName, gap.PName, cpgConst.Operate)
	//	if err != nil {
	//		common.SendResponse(ctx, common.ErrDatabase, "添加失败"+err.Error())
	//		return
	//	}
	//	if !hasPolicy {
	//		common.SendResponse(ctx, common.OK, "该角色已存在该权限")
	//		return
	//	}
	//	common.SendResponse(ctx, nil, "添加成功")
	//	return
	//} else {
	//	common.SendResponse(ctx, common.ErrAddPermission, "")
	//	return
	//}
}

// RemovePermissionsFromRole 角色移除权限
func (a Auth) RemovePermissionsFromRole(ctx *gin.Context) {
	query := new(qo.DeletePermissionFromRoleQO)
	util.JsonConvert(ctx, query)
	e, _ := auth.GetE(ctx)
	//校验角色是否存在
	role := e.GetFilteredNamedGroupingPolicy("g", 1, cpgConst.RolePrefix+query.RName)
	if len(role) < cpgConst.ONE {
		common.SendResponse(ctx, common.ErrRoleNotExisted, "")
		return
	}
	//校验权限是否为空、重复
	if len(query.PName) < cpgConst.ONE {
		common.SendResponse(ctx, common.ErrParam, "权限不能为空！")
		return
	}
	var permissionMap map[string]bool
	for _, v := range query.PName {
		if _, ok := permissionMap[v]; ok {
			common.SendResponse(ctx, common.ErrParam, "权限参数存在重复值！")
			return
		}
	}

	//解除权限-角色关联
	for _, v := range query.PName {
		_, _ = e.RemoveFilteredNamedPolicy("p", 0, cpgConst.RolePrefix+query.RName, v)
	}
	common.SendResponse(ctx, common.OK, "")
	return
}

// AddUserIntoRole 添加用户-角色关联
func (a Auth) AddUserIntoRole(ctx *gin.Context) {
	userIntoGroup := new(qo.AddUserIntoRoleQO)
	util.JsonConvert(ctx, userIntoGroup)
	uid := strconv.Itoa(userIntoGroup.Uid)
	e, _ := auth.GetE(ctx)

	//TODO 校验uid

	//校验角色是否存在
	hasGroup := e.GetFilteredNamedGroupingPolicy("g", 1, cpgConst.RolePrefix+userIntoGroup.RName)
	if len(hasGroup) == 0 {
		common.SendResponse(ctx, common.ErrRoleNotExisted, "")
		return
	}

	result, err := e.AddRoleForUser(cpgConst.UserPrefix+uid, cpgConst.RolePrefix+userIntoGroup.RName)
	if err != nil {
		common.SendResponse(ctx, common.ErrDatabase, err)
		return
	}
	if !result {
		common.SendResponse(ctx, common.ErrUserExistedInRole, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")
	return
}

// DeletePermission 移除权限，且解除权限-角色关联
func (a Auth) DeletePermission(ctx *gin.Context) {
	query := new(qo.DeletePermissionQO)
	util.JsonConvert(ctx, query)
	e, _ := auth.GetE(ctx)
	result, err := e.RemoveFilteredNamedGroupingPolicy("g2", 1, query.PName)
	if err != nil {
		common.SendResponse(ctx, common.ErrDatabase, err)
		return
	}
	if !result {
		common.SendResponse(ctx, common.ErrRemovePermission, "")
		return
	}
	_, err = e.RemoveFilteredNamedPolicy("p", 1, query.PName)
	common.SendResponse(ctx, common.OK, err)
	return
}

// DeleteRole 删除角色，且解除角色与权限关联及角色与用户关联
func (a Auth) DeleteRole(ctx *gin.Context) {
	query := new(qo.DeleteRoleQO)
	util.JsonConvert(ctx, query)
	e, _ := auth.GetE(ctx)
	//result, err := e.RemoveFilteredNamedGroupingPolicy("g", 1, cpgConst.RolePrefix+query.RName)
	//if !result {
	//	common.SendResponse(ctx, common.ErrRoleNotExisted, "")
	//	return
	//}
	//if err != nil {
	//	common.SendResponse(ctx, common.ErrDatabase, err)
	//	return
	//}
	//_, _ = e.RemoveFilteredNamedPolicy("p", 0, cpgConst.RolePrefix+query.RName)
	//common.SendResponse(ctx, common.OK, "")

	if len(query.RName) < cpgConst.ONE {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}
	for _, v := range query.RName {
		_, _ = e.RemoveFilteredNamedGroupingPolicy("g", 1, cpgConst.RolePrefix+v)
		_, _ = e.RemoveFilteredNamedPolicy("p", 0, cpgConst.RolePrefix+v)
	}
	common.SendResponse(ctx, common.OK, "")
	return
}

//GetUserRoles 查询用户角色
func (a Auth) GetUserRoles(ctx *gin.Context) {
	query := new(qo.GetUserRolesQO)
	util.JsonConvert(ctx, query)
	e, _ := auth.GetE(ctx)
	var userRolesVo []vo.UserRolesVO

	if len(query.Uid) < cpgConst.ZERO {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}

	uMap := make(map[int]bool)
	for _, v := range query.Uid {
		if _, ok := uMap[v]; ok {
			common.SendResponse(ctx, common.ErrParam, "uid存在重复!")
			return
		}
		uMap[v] = true
	}

	getRole := func(s [][]string) (res []string) {
		for _, v := range s {
			res = append(res, strings.TrimPrefix(v[cpgConst.ONE], cpgConst.RolePrefix))
		}
		return res
	}

	for _, v := range query.Uid {
		userRoles := e.GetFilteredNamedGroupingPolicy("g", cpgConst.ZERO, cpgConst.UserPrefix+strconv.Itoa(v))
		log.Println(userRoles)
		userRolesInfo := vo.UserRolesVO{
			UserId:    v,
			RoleNames: getRole(userRoles),
		}
		userRolesVo = append(userRolesVo, userRolesInfo)
	}
	userMap := userCommonFunc.UserCommonFunc{}.FindUser(ctx, query.Uid, "", "")

	for k, v := range userRolesVo {
		userRolesVo[k].UserName = userMap[uint(v.UserId)].Nickname
	}

	common.SendResponse(ctx, common.OK, userRolesVo)
	return
}

// RoleRemoveUser 用户移除角色
func (a Auth) RoleRemoveUser(ctx *gin.Context) {
	query := new(qo.DeleteUserRoleQO)
	util.JsonConvert(ctx, query)
	e, _ := auth.GetE(ctx)
	result, err := e.RemoveFilteredNamedGroupingPolicy("g", 0,
		cpgConst.UserPrefix+strconv.Itoa(query.Uid), cpgConst.RolePrefix+query.RName)

	if !result {
		common.SendResponse(ctx, common.ErrRelationshipNotExisted, "")
		return
	}
	if err != nil {
		common.SendResponse(ctx, common.ErrDatabase, err)
		return
	}
	common.SendResponse(ctx, common.OK, "")
	return
}
