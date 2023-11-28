package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Postgres struct {
	Dsn         string
	IsMigration bool
}

type Config struct {
	rest.RestConf
	Author    string
	Postgres  Postgres
	RedisConf redis.RedisConf
	Auth      struct {
		AccessSecret string
		AccessExpire int64
	}
}
