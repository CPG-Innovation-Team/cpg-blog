package dao

import (
	"cpg-blog/global/common"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/article/model"
	"cpg-blog/internal/article/vo"
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

type ArticleDAO struct {
	Aid     int
	Sn      int
	Title   string
	Uid     int `json:"uid"`
	Cover   string
	Content string
	Tags    string
	State   int
	ViewNum bool `json:"view_num"`
	CmtNum  bool `json:"cmt_num"`
	ZanNum  bool `json:"zan_num"`
	page    common.PageQO
}

var Db = &(globalInit.Db)

//增加前后百分号
func addPercent(s string) string {
	builder := strings.Builder{}
	builder.WriteString("%")
	builder.WriteString(s)
	builder.WriteString("%")
	return builder.String()
}

func (ad ArticleDAO) CreatArticle(ctx *gin.Context, article *model.Article) (err error) {
	tx := globalInit.Transaction()
	err = func(db *gorm.DB) error {
		if tx.Error != nil {
			return tx.Error
		}
		tx.Create(article)
		if tx.Error != nil {
			tx.Rollback()
			return tx.Error
		}
		return tx.Commit().Error
	}(tx)
	return err
}

func (ad ArticleDAO) SelectBySn(ctx *gin.Context, article *model.Article) *model.Article {
	(*Db).Model(&model.Article{}).Where("sn", article.Sn).First(&article)
	return article
}

func (ad ArticleDAO) FindArticles(ctx *gin.Context) (articlesVO vo.ArticleListVO) {
	tx := (*Db).WithContext(ctx)
	if ad.page.PageNum > 0 && ad.page.PageSize > 0 {
		tx.Limit(ad.page.PageSize).Offset((ad.page.PageNum - 1) * ad.page.PageSize)
	}
	if strings.Compare(ad.page.Order, "desc") == 0 {
		tx.Order(ad.page.Order)
	}

	if ad.Sn != 0 { //sn精确搜索
		tx = tx.Where("sn", ad.Sn)
	}
	if ad.Title != "" { //title模糊搜索
		tx = tx.Where("title Like ?", addPercent(ad.Title))
	}
	if ad.Uid != 0 { //uid精确搜索
		tx = tx.Where("uid", ad.Uid)
	}
	if ad.Content != "" { //模糊搜索文章内容
		tx = tx.Where("content Like ？", addPercent(ad.Content))
	}
	if ad.Tags != "" {
		tx = tx.Where("tags In ?", strings.Split(ad.Tags, ","))
	}
	if ad.State >= 0 {
		tx = tx.Where("state", ad.State)
	}
	if ad.ViewNum {
		tx = tx.Order("view_num desc")
	}
	if ad.CmtNum {
		tx = tx.Order("cmt_num desc")
	}
	if ad.ZanNum {
		tx = tx.Order("zan_num desc")
	}
	pageVO := new(common.PageVO)
	tx, pageVO = ad.page.NewPageVO(tx)
	tx.Model(&model.Article{}).Select("cpg_blog_article.aid,sn, title, uid, cover, content, tags, state, view_num, cmt_num, zan_num").
		Joins("LEFT JOIN cpg_blog_article_ex ON cpg_blog_article.aid = cpg_blog_article_ex.aid ")
	articlesVO.PageVO = *pageVO
	row, err := tx.Rows()

	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			common.SendResponse(ctx, common.ErrDatabase, err)
		}
	}(row)

	if err == nil {
		for row.Next() {
			article := &(vo.ArticleDetail{})
			err := tx.ScanRows(row, article)
			if err != nil {
				return vo.ArticleListVO{}
			}
			articlesVO.ArticleDetailList = append(articlesVO.ArticleDetailList, *article)
		}
	}
	return
}

func (ad ArticleDAO) FindArticleEx() {

}
