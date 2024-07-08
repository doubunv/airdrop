package svc

import (
	"air-drop/cmd/internal/config"
	"air-drop/cmd/internal/data/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Db        *gorm.DB
	Config    config.Config
	Redis     *redis.Redis
	UserModel *model.UserModel
}

var SCtx *ServiceContext

func NewServiceContext(c config.Config) *ServiceContext {

	//db, err := gorm.Open(mysql.Open(c.DataSource))
	//if err != nil {
	//	panic("Database not connected\n")
	//}
	config.ConfigData = c
	config.RedisClient = redis.MustNewRedis(c.RedisConf)
	SCtx = &ServiceContext{
		//Db:     db,
		Config: c,
		//Redis:  redis.MustNewRedis(c.RedisConf),
		//UserModel: model.NewUserModel(db),
	}

	return SCtx
}
