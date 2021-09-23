package model

import "time"

type Article struct {
	Aid       int   `gorm:"primaryKey; autoIncrement;"`
	Sn        int64 `gorm:"primaryKey"`
	Title     string
	Uid       int `gorm:"column:uid"`
	Cover     string
	Content   string
	Tags      string
	State     int //0-未审核;1-已上线;2-下线(审核拒绝);3-用户删除
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ArticleEx struct {
	Sn        int64 `gorm:"primaryKey"`
	ViewNum   int
	CmtNum    int
	ZanNum    int
	UpdatedAt time.Time
}
