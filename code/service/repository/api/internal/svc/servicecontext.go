package svc

import (
	"butane-netdisk/service/repository/api/internal/config"
	"butane-netdisk/service/repository/model"
	"butane-netdisk/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	RepositoryPoolModel model.RepositoryPoolModel
	UserRpc             userclient.User
	RedisClient         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		RepositoryPoolModel: model.NewRepositoryPoolModel(conn, c.CacheRedis),
		UserRpc:             userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
