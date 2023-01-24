package svc

import (
	"cloud-disk/service/repository/rpc/repositoryclient"
	"cloud-disk/service/user/rpc/userclient"
	"cloud-disk/service/user_repository/api/internal/config"
	"cloud-disk/service/user_repository/model"

	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	UserRepositoryModel model.UserRepositoryModel
	RepositoryRpc       repositoryclient.Repository
	UserRpc             userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		UserRepositoryModel: model.NewUserRepositoryModel(conn, c.CacheRedis),
		RepositoryRpc:       repositoryclient.NewRepository(zrpc.MustNewClient(c.RepositoryRpc)),
		UserRpc:             userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
