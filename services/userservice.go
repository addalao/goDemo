package services

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type UserService struct {
	db  *gorm.DB
	rds *redis.Redis
}

func NewUserService(db *gorm.DB, rds *redis.Redis) *UserService {
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
