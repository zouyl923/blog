package svc

import (
	"blog/app/admin/api/internal/config"
	"blog/app/admin/api/internal/middleware"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	CorsMiddleware       rest.Middleware
	AuthMiddleware       rest.Middleware
	PermissionMiddleware rest.Middleware

	DB    *gorm.DB
	Cache *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := GetRedis(c)
	db := GetOrm(c)
	return &ServiceContext{
		Config: c,

		CorsMiddleware:       middleware.NewCorsMiddleware().Handle,
		AuthMiddleware:       middleware.NewAuthMiddleware(c, rds).Handle,
		PermissionMiddleware: middleware.NewPermissionMiddleware(db, rds).Handle,

		DB:    db,
		Cache: rds,
	}
}

func GetRedis(c config.Config) *redis.Redis {
	rds := redis.MustNewRedis(c.Redis)
	return rds
}

func GetOrm(c config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(GetDsn(c)))
	if err != nil {
		//中断程序并报错
		panic(err)
	}
	db.Logger.LogMode(4)
	return db
}

func GetDsn(c config.Config) string {
	conf := c.MySql
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.Charset,
	)
}
