package services

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"goDemo/models"
	"gorm.io/gorm"
)

type ContentService struct {
	db  *gorm.DB
	rds *redis.Redis
}

func NewContentService(db *gorm.DB, rds *redis.Redis) *ContentService {
	err := db.AutoMigrate(&models.Content{})
	fmt.Printf("err-------------%s\n", err)

	if err != nil {
		panic("")
	}

	return &ContentService{
		db:  db,
		rds: rds,
	}

}

func (s *ContentService) GetContentItem(id string) {
	println(id)
}
