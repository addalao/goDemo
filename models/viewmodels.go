package models

import "gorm.io/gorm"

type Views struct {
	gorm.Model
	ContentId string `json:"contentId" gorm:"type: varchar(32);not null;"` // 内容id
	UserId    string `json:"userId" gorm:"type varchar(32); not null"`
}
