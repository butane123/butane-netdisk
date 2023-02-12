package svc

import (
	"butane-netdisk/service/repository/rpc/repositoryclient"
	"butane-netdisk/service/share/api/internal/config"
	"butane-netdisk/service/share/model"
	"butane-netdisk/service/user_repository/rpc/userrepositoryclient"

	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	ShareBasicModel   model.ShareBasicModel
	UserRepositoryRpc userrepositoryclient.UserRepository
	RepositoryRpc     repositoryclient.Repository
	RedisClient       *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:            c,
		ShareBasicModel:   model.NewShareBasicModel(conn, c.CacheRedis),
		UserRepositoryRpc: userrepositoryclient.NewUserRepository(zrpc.MustNewClient(c.UserRepositoryRpc)),
		RepositoryRpc:     repositoryclient.NewRepository(zrpc.MustNewClient(c.RepositoryRpc)),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
