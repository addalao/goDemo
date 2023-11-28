package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"goDemo/internal/config"
	"goDemo/models"
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
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(postgres.Open(c.Postgres.Dsn))

	if err != nil {

		panic(err)
	}

	rds := redis.MustNewRedis(c.RedisConf)

	tables := []interface{}{
		&models.User{},
	}
	for _, table := range tables {
		err := db.AutoMigrate(table)

		if err != nil {
			panic(err)
		}
	}

	return &ServiceContext{
		Config:               c,
		Db:                   db,
		ContentService:       services.NewContentService(db, rds),
		UserService:          services.NewUserService(db, rds),
		AuthorizationService: services.NewAuthorizationService(c.Auth.AccessSecret),
	}
}
