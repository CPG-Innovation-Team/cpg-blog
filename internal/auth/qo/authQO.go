package qo

type PermissionQO struct {
	PName string `json:"pName" binding:"required"`
	Uri   string `json:"uri" binding:"required"`
}

type AddRoleQO struct {
	RName string `json:"rName" binding:"required"` //role name
}

type GroupAddPermissionQO struct {
	RName string   `json:"rName" binding:"required"` //role name
	PName []string `json:"pName" binding:"required"` //policy name
}

type AddUserIntoRoleQO struct {
	Uid   int    `json:"uid" binding:"required"`
	RName string `json:"rName" binding:"required"` //role name
}

type DeletePermissionQO struct {
	PName string `json:"pName" binding:"required"` //policy name
}

type DeleteRoleQO struct {
	RName []string `json:"rName" binding:"required"` //role name
}

type GetUserRolesQO struct {
	Uid []int `json:"uid" binding:"required"`
}

type DeleteUserRoleQO struct {
	Uid   int    `json:"uid" binding:"required"`   //user id
	RName string `json:"rName" binding:"required"` //role name
}
