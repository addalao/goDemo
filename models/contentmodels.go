package models

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	UserId   string `json:"userId" gorm:"type:varchar(32);not null"`
	Title    string `json:"title" gorm:"type: varchar(255);" `
	Content  string `json:"content" gorm:"type: varchar(255);not null"`
	Image    string `json:"image" gorm:"type: varchar(255);not null"`
	Distance int64  `json:"distance" gorm:"type:int;default:0"` // 距离
}
