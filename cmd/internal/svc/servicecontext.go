package svc

import (
	"air-drop/cmd/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Db     *gorm.DB
	Config config.Config
	Redis  *redis.Redis
}

var SCtx *ServiceContext

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := gorm.Open(mysql.Open(c.DataSource))
	if err != nil {
		panic("Database not connected\n")
	}
	config.ConfigData = c
	config.RedisClient = redis.MustNewRedis(c.RedisConf)
	SCtx = &ServiceContext{
		Db:     db,
		Config: c,
		//Redis:  redis.MustNewRedis(c.RedisConf),
	}

	return SCtx
}
