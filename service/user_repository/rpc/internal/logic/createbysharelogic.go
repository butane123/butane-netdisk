package logic

import (
	"cloud-disk/common/utils"
	"cloud-disk/service/user_repository/model"
	"context"
	"database/sql"

	"cloud-disk/service/user_repository/rpc/internal/svc"
	"cloud-disk/service/user_repository/rpc/userRepository"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateByShareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateByShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateByShareLogic {
	return &CreateByShareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateByShareLogic) CreateByShare(in *userRepository.CreateByShareReq) (*userRepository.CreateByShareReply, error) {
	newIdentity := utils.GenerateUUID()
	_, err := l.svcCtx.UserRepositoryModel.Insert(l.ctx, &model.UserRepository{
		Identity:           sql.NullString{newIdentity, true},
		UserIdentity:       sql.NullString{in.UserIdentity, true},
		ParentId:           sql.NullInt64{in.ParentId, true},
		RepositoryIdentity: sql.NullString{in.RepositoryIdentity, true},
		Name:               sql.NullString{in.Name, true},
	})
	if err != nil {
		return nil, err
	}
	return &userRepository.CreateByShareReply{Identity: newIdentity}, nil
}
