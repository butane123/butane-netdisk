package svc

import (
	"butane-netdisk/service/user_repository/model"
	"butane-netdisk/service/user_repository/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	UserRepositoryModel model.UserRepositoryModel
	RedisClient         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		UserRepositoryModel: model.NewUserRepositoryModel(conn, c.CacheRedis),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
