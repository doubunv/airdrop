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
	ChainInfo  ChainInfo
}

type Auth struct {
	AccessSecret string
	AccessExpire int64
}

type ChainInfo struct {
	ChainID          int64
	Rpc              string
	Name             string
	InitBlock        int64
	PerLimit         int64
	ConfirmByBlocks  int64
	PrivateKey       string
	ChainBootAddress string
}
