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
	SecretKey string
	RedisConf redis.RedisConf
}
