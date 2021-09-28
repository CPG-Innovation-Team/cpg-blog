package qo

// UnlikeQO 取消点赞
type UnlikeQO struct {
	//被取消点赞文章对象sn号
	Sn int64

	//被取消点赞blogger对象uid
	Uid uint
}
