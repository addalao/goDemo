package services

import (
	"goDemo/models"
	"gorm.io/gorm"
)

type ViewService struct {
	db *gorm.DB
}

func NewViewService(db *gorm.DB) *ViewService {

	err := db.AutoMigrate(&models.Views{})

	if err != nil {
		panic("ViewModels")
	}

	return &ViewService{
		db: db,
	}
}

func (s *ViewService) GetCount(contentId string) (count int64, err error) {

	q := s.db.Model(&models.Views{})

	err = q.Error

	if err != nil {
		return
	}

	q.Where("contentId = ?", contentId).Count(&count)

	return
}
