package qo

/**
  @author: chenxi@cpgroup.cn
  @date:2021/11/8
  @description: 审核接口请求参数
**/

//ReviewArticleQO 审核文章请求参数
type ReviewArticleQO struct {
	/*
		文章sn号
	*/
	Sn string `binding:"required"`

	/*
		审核状态（通过/不通过 => true/false）
	 */
	State bool
}

//ReviewCommentQO 审核评论请求参数
type ReviewCommentQO struct {
	//评论ID
	CommentId int `binding:"required"`

	/*
		审核状态（通过/不通过 => true/false）
	*/
	State bool
}

//ReviewReplyQO 审核评论回复请求参数
type ReviewReplyQO struct {
	//评论回复ID
	ReplyId int `binding:"required"`

	/*
		审核状态（通过/不通过 => true/false）
	*/
	State bool
}