package services

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"goDemo/models"
	"gorm.io/gorm"
)

type UserService struct {
	db  *gorm.DB
	rds *redis.Redis
}

func NewUserService(db *gorm.DB, rds *redis.Redis) *UserService {

	err := db.AutoMigrate(&models.User{})

	if err != nil {
		panic("User")
	}

	return &UserService{
		db:  db,
		rds: rds,
	}

}

func (s *UserService) GetId(ctx context.Context) string {

	id := ctx.Value("jti")

	if id != nil {
		return id.(string)
	}

	return ""
}

func (s *UserService) FindOne(userId string) (user *models.User, err error) {

	err = s.db.First(user, userId).Error

	return
}
