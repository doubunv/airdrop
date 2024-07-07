package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

var ConfigData Config
var RedisClient *redis.Redis

type Config struct {
	Version string
	rest.RestConf
	RedisConf  redis.RedisConf
	DataSource string
	Auth       Auth
	Env        string
}

type Auth struct {
	AccessSecret string
	AccessExpire int64
}
