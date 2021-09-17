package common

import (
	"gorm.io/gorm"
)

// PageQO 分页查询请求结构体
type PageQO struct {
	PageNum  int `form:"pageNum" json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
	// Order 默认是desc, 可选: desc, asc
	Order string `form:"order" json:"order"`
}

// PageVO 分页查询返回结构体
type PageVO struct {
	//当前页
	pageNum int
	//总条数
	Total int64
	//总页数
	TotalPage int
}

// NewPageVO 生成PageVO
func (pq *PageQO)NewPageVO(db *gorm.DB) (*gorm.DB, *PageVO) {
	pv := &PageVO{}
	if pq.PageNum > 0 && pq.PageSize > 0 {
		pv.TotalPage = int(pv.Total) / (pq.PageSize)
		pv.pageNum = pq.PageNum
		db = db.Limit(pq.PageSize).
			Offset((pq.PageNum - 1) * pq.PageSize).
			Count(&(pv.Total))
		return db, pv
	}
	return db, pv
}
