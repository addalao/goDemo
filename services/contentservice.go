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

	fmt.Printf("len(zRange)---------%v\n", len(zRange))

	if err != nil {
		return err
	}

	for _, s2 := range zRange {
		println(s2)
	}

	return
}
