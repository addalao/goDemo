package models

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	UserId  string
	Title   string
	Content string
	Image   string
	Gender  int16 // 1男 2女
}
