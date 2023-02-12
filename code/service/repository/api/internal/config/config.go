package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	UserRpc    zrpc.RpcClientConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis redis.RedisConf
}
