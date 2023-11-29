package models

import "gorm.io/gorm"

type Likes struct {
	gorm.Model
	ContentId string `json:"contentId" gorm:"type: varchar(32);not null;"` // 内容ID
	UserId    string `json:"userId" gorm:"type varchar(32); not null"`
}
