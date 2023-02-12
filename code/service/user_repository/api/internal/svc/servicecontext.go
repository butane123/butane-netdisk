package svc

import (
	"butane-netdisk/service/repository/rpc/repositoryclient"
	"butane-netdisk/service/user/rpc/userclient"
	"butane-netdisk/service/user_repository/api/internal/config"
	"butane-netdisk/service/user_repository/model"

	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	UserRepositoryModel model.UserRepositoryModel
	RepositoryRpc       repositoryclient.Repository
	UserRpc             userclient.User
	RedisClient         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		UserRepositoryModel: model.NewUserRepositoryModel(conn, c.CacheRedis),
		RepositoryRpc:       repositoryclient.NewRepository(zrpc.MustNewClient(c.RepositoryRpc)),
		UserRpc:             userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
