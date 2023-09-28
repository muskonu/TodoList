package config

import (
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Host     string
		Password string
	}
	Mysql      gormc.Mysql
	CacheRedis cache.CacheConf
	CSRFKey    string
}
