package model

import "time"

type Comment struct {
	//自增ID
	Cid uint `gorm:"primaryKey; autoIncrement; column:cid"`

	//文章sn号
	Sn int64

	//评论用户uid
	UID uint `gorm:"column:uid"`

	//评论内容
	Content string

	//点赞数
	ZanNum int

	//第几楼
	Floor int

	//状态：0-未审核;1-已上线;2-下线(审核拒绝);3-用户删除
	State int

	CreatedAt time.Time
}
