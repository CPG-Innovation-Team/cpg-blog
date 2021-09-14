package qo

import "cpg-blog/global/cpgConst"

type PermissionQO struct {
	Name    string `json:"name"`
	Uri     string `json:"uri"`
	Operate string `json:"-"`
}

type AddGroupQO struct {
	RName string `json:"RName"` //role name
}

type GroupAddPermissionQO struct {
	RName string `json:"RName"` //role name
	PName string `json:"pName"` //policy name
}

type AddUserIntoRoleQO struct {
	Uid   int    `json:"uid"`
	RName string `json:"RName"` //role name
}

func GetNewPermission() (p PermissionQO) {
	return PermissionQO{Operate: cpgConst.Operate}
}
