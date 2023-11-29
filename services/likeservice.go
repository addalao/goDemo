package services

import (
	"goDemo/models"
	"gorm.io/gorm"
)

type LikeService struct {
	db *gorm.DB
}

func NewLikeService(db *gorm.DB) *LikeService {

	err := db.AutoMigrate(&models.Likes{})

	if err != nil {
		panic("LikeModels")
	}

	return &LikeService{
		db: db,
	}
}

func (s *LikeService) GetCount(contentId string) (count int64, err error) {

	q := s.db.Model(&models.Likes{})

	err = q.Error

	if err != nil {
		return
	}

	q.Where("contentId = ?", contentId).Count(&count)

	return
}
