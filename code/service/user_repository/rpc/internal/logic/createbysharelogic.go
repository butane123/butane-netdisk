package logic

import (
	"butane-netdisk/common/utils"
	"butane-netdisk/service/user_repository/model"
	"butane-netdisk/service/user_repository/rpc/internal/svc"
	"butane-netdisk/service/user_repository/rpc/types/userRepository"
	"context"

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
	newId := utils.GenerateNewId(l.svcCtx.RedisClient, "user_repository")
	_, err := l.svcCtx.UserRepositoryModel.InsertWithId(l.ctx, &model.UserRepository{
		Id:           newId,
		UserId:       in.UserId,
		ParentId:     in.ParentId,
		RepositoryId: in.RepositoryId,
		Name:         in.Name,
	})
	if err != nil {
		return nil, err
	}
	return &userRepository.CreateByShareReply{Id: newId}, nil
}
