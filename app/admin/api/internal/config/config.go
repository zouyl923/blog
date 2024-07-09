package config

import (
	"github.com/zeromicro/go-zero/rest"
	"time"
)

type Config struct {
	rest.RestConf

	MySql struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
		Prefix   string
		Charset  string
	}
	Redis struct {
		Host     string
		Type     string `json:",default=node,options=node|cluster"`
		Pass     string `json:",optional"`
		Tls      bool   `json:",optional"`
		NonBlock bool   `json:",default=true"`
		// PingTimeout is the timeout for ping redis.
		PingTimeout time.Duration `json:",default=1s"`
	}
	JwtSecret string
}
