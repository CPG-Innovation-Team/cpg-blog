package vo

import (
	"time"
)

/**
  @author: chenxi@cpgroup.cn
  @date:2021/11/10
  @description: 返回参数
**/

type ReviewArticle struct {
	Aid       int
	Sn        int64
	Title     string
	Uid       int
	Cover     string
	Content   string
	Tags      string
	State     int //0-未审核;1-已上线;2-下线(审核拒绝);3-用户删除
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	//自增ID
	Cid uint
	//文章sn号
	Sn int64
	//评论用户uid
	UID uint
	//评论内容
	Content string
	//点赞数
	ZanNum int
	//第几楼
	Floor int
	//状态：0-未审核;1-已上线;2-下线(审核拒绝);3-用户删除
	State     int
	CreatedAt time.Time
}

type Reply struct {
	//自增ID
	Id uint
	//评论cid
	Cid uint
	//回复用户uid
	UID uint
	//回复内容
	Content string
	//状态：0-未审核;1-已上线;2-下线(审核拒绝);3-用户删除
	State     int
	CreatedAt time.Time
}

type ReviewArticleListVO struct {
	ArticleMap map[int64]ReviewArticle
}

type ReviewCommentListVO struct {
	CommentMap map[uint]Comment
}

type ReviewReplyListVO struct {
	ReplyMap map[uint]Reply
}

type ReviewCommentVO struct {
	CommentList []int
}

type ReviewReplyVO struct {
	ReplyList []int
}
