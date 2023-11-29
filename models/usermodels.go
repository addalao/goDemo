package models

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		Username string `json:"username" gorm:"type: varchar(32);not null;unique"` //用户名
		Password string `json:"password" gorm:"type: varchar(255);not null"`       //密码
		Gender   int16  `json:"gender" gorm:"default:0"`                           // 性别
	}
)
