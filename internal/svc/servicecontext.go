package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"goDemo/internal/config"
	"goDemo/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config               config.Config
	Db                   *gorm.DB
	ContentService       *services.ContentService
	UserService          *services.UserService
	AuthorizationService *services.AuthorizationService
	LikeService          *services.LikeService
	ViewService          *services.ViewService
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(postgres.Open(c.Postgres.Dsn))

	if err != nil {

		panic(err)
	}

	rds := redis.MustNewRedis(c.RedisConf)

	return &ServiceContext{
		Config:               c,
		Db:                   db,
		ContentService:       services.NewContentService(db, rds),
		UserService:          services.NewUserService(db, rds),
		AuthorizationService: services.NewAuthorizationService(c.Auth.AccessSecret),
		LikeService:          services.NewLikeService(db),
		ViewService:          services.NewViewService(db),
	}
}
