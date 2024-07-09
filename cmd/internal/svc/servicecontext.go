package svc

import (
	"air-drop/cmd/internal/config"
	"air-drop/cmd/internal/data/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Db     *gorm.DB
	Config config.Config
	Redis  *redis.Redis
	*model.UserModel
	*model.AmountChangeTypeModel
	*model.AmountLogModel
	*model.LinkModel
	*model.LinkOrderModel
	*model.LinkReceiveModel
	*model.PackageChildModel
	*model.PackageModel
	*model.PackageOrderModel
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
		UserModel:             model.NewUserModel(db),
		AmountChangeTypeModel: model.NewAmountChangeTypeModel(db),
		AmountLogModel:        model.NewAmountLogModel(db),
		LinkModel:             model.NewLinkModel(db),
		LinkOrderModel:        model.NewLinkOrderModel(db),
		LinkReceiveModel:      model.NewLinkReceiveModel(db),
		PackageChildModel:     model.NewPackageChildModel(db),
		PackageModel:          model.NewPackageModel(db),
		PackageOrderModel:     model.NewPackageOrderModel(db),
	}

	return SCtx
}
