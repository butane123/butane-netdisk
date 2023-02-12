package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf
	Redis      redis.RedisConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
	Salt string
}
