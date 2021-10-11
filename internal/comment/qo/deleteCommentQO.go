package qo

/**
  @author: ethan.chen@cpgroup.cn
  @date:2021/10/11
  @description:删除评论请求参数
**/

type DeleteCommentQO struct {
	//评论ID
	CommentId int `binding:"required"`
}
