package qo

type AddCommentQO struct {
	//文章sn号
	Sn int64 `binding:"required"`

	//评论内容
	Content string`binding:"required"`
}
