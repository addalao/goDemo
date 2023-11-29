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
		panic("Content")
	}

	return &ContentService{
		db:  db,
		rds: rds,
	}

}

func (s *ContentService) GetContentItem(page int, pageSize int) (items []*models.Content, total int64, err error) {

	page = (page - 1) * pageSize

	query := s.db.Model(&models.Content{})

	err = query.Error

	if err != nil {
		return
	}

	err = query.Count(&total).Error

	if err != nil {
		return
	}

	err = query.Offset(page).Limit(pageSize).Scan(&items).Error

	return
}

// Create 创建
func (s *ContentService) Create(userId string, title string, content string, image string) error {

	return s.db.Create(&models.Content{
		UserId:  userId,
		Title:   title,
		Content: content,
		Image:   image,
	}).Error

}
