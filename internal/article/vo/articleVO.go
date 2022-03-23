package vo

import "cpg-blog/global/common"

// AddArticleVO 新增文章返回参数
type AddArticleVO struct {
	Sn int64
}

// ArticleInfoVO 返回文章详情
type ArticleInfoVO struct {
	/*
		文章sn号
	*/
	Sn string

	/*
		文章标题
	*/
	Title string

	/*
		作者昵称
	*/
	Author string

	/*
		作者头像
	 */
	Avatar string

	/*
		文章封面图地址
	*/
	Cover string

	/*
		内容，markdown格式
	*/
	Content string

	/*
		文章 tag，逗号分隔
	*/
	Tags string

	/*
		状态 0-未审核;1-已上线;2-下线(审核拒绝);3-用户删除',
	*/
	State int

	/*
		点赞状态 true-已点赞， false-未点赞
	*/
	ZanState bool

	/*
		创建时间
	*/
	CreateAt int64
	/*
		更新时间
	*/
	UpdatedAt int64
}

// ArticleDetail 返回文章list详情
type ArticleDetail struct {
	/*
		文章sn号
	*/
	Sn string

	/*
		文章标题
	*/
	Title string

	/*
		uid
	*/
	Uid uint

	/*
		作者昵称
	*/
	Author string

	/*
		作者头像
	*/
	Avatar string

	/*
		文章封面图地址
	*/
	Cover string

	/*
		内容，markdown格式
	*/
	Content string

	/*
		文章 tag，逗号分隔
	*/
	Tags string

	/*
		状态 0-未审核;1-已上线;2-下线(审核拒绝);3-用户删除',
	*/
	State int

	/*
		浏览量
	*/
	ViewNum int

	/*
		评论数
	*/
	CmtNum int

	/*
		点赞数
	*/
	ZanNum int

	/*
		分页
	*/
	page common.PageVO
}

type ArticleListVO struct {
	ArticleDetailList []ArticleDetail
	common.PageVO
}
