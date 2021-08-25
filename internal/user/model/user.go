package model

import (
	"time"
)

type User struct {
	UID       uint   `gorm:"primaryKey; autoIncrement; column:uid"`
	UserName  string `gorm:"column:username"`
	Email     string
	Passcode  string
	Passwd    string
	Nickname  string
	Avatar    string
	Gender    uint8
	Introduce string
	State     uint8
	IsRoot    bool `gorm:"column:is_root"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
