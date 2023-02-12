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
	Redis             redis.RedisConf
	CacheRedis        cache.CacheConf
	UserRepositoryRpc zrpc.RpcClientConf
	RepositoryRpc     zrpc.RpcClientConf
	Auth              struct {
		AccessSecret string
		AccessExpire int64
	}
}
