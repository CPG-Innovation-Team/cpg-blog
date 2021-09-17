package qo

// LoginQO 登录请求参数
type LoginQO struct {
	Username string `json:"username" binding:"required"`
	Passwd   string `json:"passwd" binding:"required"`
}

// RegisterQO 用户注册请求参数
type RegisterQO struct {
	/*
		用户姓名（唯一）
	*/
	UserName string `json:"username" binding:"required"`

	/*
		用户邮箱（唯一）
	*/
	Email string `json:"email" binding:"required"`

	/*
		加密随机数
	*/
	PassCode string `json:"passCode" binding:"required"`

	/*
		用户md5密码
	*/
	Passwd string `json:"passwd" binding:"required"`

	/*
		用户昵称
	*/
	Nickname string `json:"nickname" binding:"required"`

	/*
		用户头像地址
	*/
	Avatar string `json:"avatar"`

	/*
		性别
	*/
	Gender int `json:"gender"`

	/*
		简介
	*/
	Introduce string `json:"introduce"`

	/*
		状态
	*/
	State int `json:"state"`
	/*
		是否为超级用户
	*/
	IsRoot int `json:"isRoot"`
}

// UserInfoQO 查询用户信息请求参数
type UserInfoQO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// ModifyQO 修改用户信息请求参数
type ModifyQO struct {
	/*
		用户ID（唯一且不允许修改）
	*/
	UID uint `json:"uid" binding:"required"`

	/*
		用户姓名（唯一）
	*/
	UserName string `json:"username" binding:"required"`

	/*
		用户邮箱（唯一）
	*/
	Email string `json:"email" binding:"required"`

	/*
		加密随机数
	*/
	PassCode string `json:"passCode" binding:"required"`

	/*
		用户md5密码
	*/
	Passwd string `json:"passwd" binding:"required"`

	/*
		用户昵称
	*/
	Nickname string `json:"nickname" binding:"required"`

	/*
		用户头像地址
	*/
	Avatar string `json:"avatar"`

	/*
		性别
	*/
	Gender int `json:"gender"`

	/*
		简介
	*/
	Introduce string `json:"introduce"`

	/*
		状态
	*/
	State int `json:"state"`
	/*
		是否为超级用户
	*/
	IsRoot int `json:"isRoot"`
}

type UserListQO struct {
	/*
		用户状态 1-正常;2-禁发文;3-冻结
	*/
	State int `json:"state"`
}
