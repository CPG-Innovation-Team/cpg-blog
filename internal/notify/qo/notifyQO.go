package qo

import "time"

/**
  @author: chenxi@cpgroup.cn
  @date:2021/12/8
  @description:通知相关接口请求参数
**/

// AddNotificationQO 添加通知请求参数
type AddNotificationQO struct {
	// 必传；通知类型：文章相关-1，点赞相关-2，评论相关-3，系统通知-4，其他-5
	Type int `binding:"required"`

	// 通知类型为4时默认填充0，其余情况需要绑定用户ID
	Uid []int

	// 通知内容
	Content string

	// 通知状态（默认为0）：关闭-0，开启-1
	State int

	// 通知开始时间
	BeginTime time.Time

	// 通知结束时间
	EndTime time.Time
}

// UpdateNotificationQO 更新通知请求参数
type UpdateNotificationQO struct {
	//必传
	Id int `binding:"required"`

	// 必传；通知类型：文章相关-1，点赞相关-2，评论相关-3，系统通知-4，其他-5
	Type int `binding:"required"`

	// 通知类型为4时默认填充0，其余情况需要绑定用户ID
	Uid []int

	// 通知内容
	Content string

	// 通知状态（默认为0）：关闭-0，开启-1
	State int

	// 通知开始时间
	BeginTime time.Time

	// 通知结束时间
	EndTime time.Time
}