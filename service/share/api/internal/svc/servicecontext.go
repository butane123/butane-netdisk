package svc

import (
	"cloud-disk/service/repository/rpc/repositoryclient"
	"cloud-disk/service/share/api/internal/config"
	"cloud-disk/service/share/model"
	"cloud-disk/service/user_repository/rpc/userrepositoryclient"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	ShareBasicModel   model.ShareBasicModel
	UserRepositoryRpc userrepositoryclient.UserRepository
	RepositoryRpc     repositoryclient.Repository
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:            c,
		ShareBasicModel:   model.NewShareBasicModel(conn, c.CacheRedis),
		UserRepositoryRpc: userrepositoryclient.NewUserRepository(zrpc.MustNewClient(c.UserRepositoryRpc)),
		RepositoryRpc:     repositoryclient.NewRepository(zrpc.MustNewClient(c.RepositoryRpc)),
	}
}
