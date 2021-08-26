package qo

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

type ArticleInfoQO struct {
}
