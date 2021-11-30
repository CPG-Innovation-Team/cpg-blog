package qo

/**
  @author: ethan.chen@cpgroup.cn
  @date:2021/10/12
  @description:查询文章下所有评论及其回复请求参数
**/

type ListQO struct {
	//文章sn号
	Sn string `binding:"required"`
}
