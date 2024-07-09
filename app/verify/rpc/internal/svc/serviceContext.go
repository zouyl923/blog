package svc

import (
	"blog/app/verify/rpc/internal/config"
	"blog/common/helper"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Cache  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Cache:  GetRedis(c),
	}
}

func GetRedis(c config.Config) *redis.Redis {
	var conf redis.RedisConf
	helper.ExchangeStruct(c.Redis, &conf)
	rds := redis.MustNewRedis(conf)
	return rds
}
