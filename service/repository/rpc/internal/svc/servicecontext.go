package svc

import (
	"cloud-disk/service/repository/model"
	"cloud-disk/service/repository/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	RepositoryPoolModel model.RepositoryPoolModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn:=sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		RepositoryPoolModel: model.NewRepositoryPoolModel(conn,c.CacheRedis),
	}
}
