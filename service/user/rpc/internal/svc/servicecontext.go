package svc

import (
	"cloud-disk/service/user/model"
	"cloud-disk/service/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	UserBasicModel model.UserBasicModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		UserBasicModel: model.NewUserBasicModel(conn, c.CacheRedis),
	}
}
