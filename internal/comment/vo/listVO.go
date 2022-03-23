package vo

import "time"

/**
  @author: ethan.chen@cpgroup.cn
  @date:2021/10/12
  @description:所有评论及其回复返回参数
**/

type CommentListVO struct {
	//自增ID
	Cid uint

	//文章sn号
	Sn int64

	//评论用户uid
	UID uint

	//评论用户昵称
	NickName string

	//评论用户头像
	Avatar string

	//评论内容
	Content string

	//点赞数
	ZanNum int

	//第几楼
	Floor int

	//状态：0-未审核;1-已上线;2-下线(审核拒绝);3-用户删除
	State int

	//点赞状态 true-已点赞， false-未点赞
	ZanState bool

	CreatedAt time.Time

	ReplyList []ReplyVO
}

type ReplyVO struct {
	//自增ID
	Id uint

	//评论cid
	Cid uint

	//回复用户uid
	UID uint

	//回复用户昵称
	NickName string

	//回复用户头像
	Avatar string

	//回复内容
	Content string

	//状态：0-未审核;1-已上线;2-下线(审核拒绝);3-用户删除
	State int

	CreatedAt time.Time
}
