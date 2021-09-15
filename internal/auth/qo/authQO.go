package qo

type PermissionQO struct {
	PName string `json:"pName"`
	Uri   string `json:"uri"`
}

type AddRoleQO struct {
	RName string `json:"rName"` //role name
}

type GroupAddPermissionQO struct {
	RName string `json:"rName"` //role name
	PName string `json:"pName"` //policy name
}

type AddUserIntoRoleQO struct {
	Uid   int    `json:"uid"`
	RName string `json:"rName"` //role name
}

type DeletePermissionQO struct {
	PName string `json:"pName"` //policy name
}

type DeleteRoleQO struct {
	RName string `json:"rName"` //role name
}

type DeleteUserRoleQO struct {
	Uid int `json:"uid"` //user id
	RName string `json:"rName"` //role name
}
