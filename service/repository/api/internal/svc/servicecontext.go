package svc

import (
	"cloud-disk/service/repository/api/internal/config"
	"cloud-disk/service/repository/model"
	"cloud-disk/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	RepositoryPoolModel model.RepositoryPoolModel
	UserRpc             userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		RepositoryPoolModel: model.NewRepositoryPoolModel(conn, c.CacheRedis),
		UserRpc:             userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
