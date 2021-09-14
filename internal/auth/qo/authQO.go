package qo

type PermissionQO struct {
	PName string `json:"PName"`
	Uri   string `json:"uri"`
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
