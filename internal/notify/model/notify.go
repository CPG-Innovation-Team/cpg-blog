package model

import "time"

/**
  @author: chenxi@cpgroup.cn
  @date:2021/12/8
  @description: 通知详情表
**/

type Notify struct {
	// 自增ID
	Id int `gorm:"primaryKey; autoIncrement;"`

	// 通知类型：文章相关-1，点赞相关-2，评论相关-3，系统通知-4，其他-5
	Type int

	// 通知类型为4时默认填充0，其余情况需要绑定用户ID列表
	Uid string `gorm:"column:uid"`

	// 通知内容
	Content string

	// 通知状态（默认为0）：关闭-0，开启-1
	State int

	// 通知开始时间
	BeginTime time.Time

	// 通知结束时间
	EndTime time.Time

	// 创建时间
	CreatedAt time.Time

	// 更新时间
	UpdatedAt time.Time
}
