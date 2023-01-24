package svc

import (
	"cloud-disk/service/user_repository/model"
	"cloud-disk/service/user_repository/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	UserRepositoryModel model.UserRepositoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		UserRepositoryModel: model.NewUserRepositoryModel(conn, c.CacheRedis),
	}
}
