package model

import "time"

type Zan struct {
	// 自增ID
	Id int `gorm:"primaryKey; autoIncrement;"`

	// 点赞用户uid
	Uid uint `gorm:"column:uid"`

	// 被赞对象类型,0-文章;1-评论
	ObjType int

	// 被赞对象id，属主(对象为文章时为sn号)
	ObjId int64

	// 点赞状态(0-未删除，1-已删除)
	State int

	// 点赞时间
	CreatedAt time.Time

	// 点赞更新时间
	UpdatedAt time.Time
}
