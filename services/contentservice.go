package services

import (
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

	if err != nil {
		panic("")
	}

	return &ContentService{
		db:  db,
		rds: rds,
	}

}

func (s *ContentService) GetContentItem() (err error) {

	_, err = s.rds.Zadd("sortName", 1, "liufei")

	if err != nil {
		return
	}

	zRange, err := s.rds.Zrange("sortName", 0, -1)

	if err != nil {
		return err
	}

	for _, s2 := range zRange {
		println(s2)
	}

	return
}

// Create 创建
func (s *ContentService) Create(userId string, title string, content string, image string, gender int16) error {

	return s.db.Create(&models.Content{
		UserId:  userId,
		Title:   title,
		Content: content,
		Image:   image,
		Gender:  gender,
	}).Error

}
