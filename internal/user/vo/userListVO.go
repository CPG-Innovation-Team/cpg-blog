package vo

// UserListVO 用户列表返回参数
type UserListVO struct {
	/*
		用户姓名（唯一）
	*/
	UserName string `json:"username" binding:"required"`

	/*
		用户邮箱（唯一）
	*/
	Email string `json:"email" binding:"required"`

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
