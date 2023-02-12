package svc

import (
	"butane-netdisk/service/user/api/internal/config"
	"butane-netdisk/service/user/model"
	model2 "butane-netdisk/service/user_repository/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	UserBasicModel      model.UserBasicModel
	UserRepositoryModel model2.UserRepositoryModel
	RedisClient         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		UserBasicModel:      model.NewUserBasicModel(conn, c.CacheRedis),
		UserRepositoryModel: model2.NewUserRepositoryModel(conn, c.CacheRedis),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
