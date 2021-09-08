package qo

import (
	"cpg-blog/global/common"
)

// AddArticleQO 新增文章
type AddArticleQO struct {
	/*
		文章标题
	*/
	Title string

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
}

// ArticleInfoQO 查询文章详情
type ArticleInfoQO struct {
	/*
		文章sn号
	*/
	Sn int `binding:"required"`
}

type Article struct {
	/*
	文章id，关联扩展表aid
	 */
	Aid     int

	/*
	文章sn号
	 */
	Sn      int

	/*
		文章标题
	*/
	Title   string

	/*
	作者uid
	 */
	Uid     int `json:"uid"`

	/*
		内容，markdown格式
	*/
	Content string

	/*
		文章 tag，逗号分隔
	*/
	Tags    string

	/*
	文章状态 0-未审核;1-已上线;2-下线;3-用户删除'
	 */
	State   int

	/*
	浏览量排序，默认asc
	 */
	ViewNum bool `json:"view_num"`

	/*
	评论数排序，默认asc
	 */
	CmtNum  bool `json:"cmt_num"`

	/*
	点赞数排序，默认asc
	 */
	ZanNum  bool `json:"zan_num"`
}

// ArticleListQO 根据条件搜索文章
type ArticleListQO struct {
	/*
		根据条件搜索所有的文章，否则查询自身所有文章
	*/
	IsAllMyselfArticles bool `json:"isAllMyselfArticles"`

	/*
		通过参数搜索文章
	*/
	Article

	/*
		分页
	*/
	page common.PageQO
}



