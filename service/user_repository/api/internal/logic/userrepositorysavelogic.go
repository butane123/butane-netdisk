package logic

import (
	"cloud-disk/common/utils"
	"cloud-disk/service/repository/rpc/repository"
	"cloud-disk/service/user/rpc/user"
	"cloud-disk/service/user_repository/model"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"cloud-disk/service/user_repository/api/internal/svc"
	"cloud-disk/service/user_repository/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest) (resp *types.UserRepositorySaveResponse, err error) {
	// 判断文件是否超容量
	// //获取文件Size
	repositoryPoolInfo, err := l.svcCtx.RepositoryRpc.GetRepositoryPoolByRepositoryId(l.ctx, &repository.RepositoryReq{RepositoryId: req.RepositoryIdentity})
	if err != nil {
		return nil, err
	}
	// //更新个人total_Volume、nowVolume
	userIdentity := json.Number(fmt.Sprintf("%v", l.ctx.Value("userIdentity"))).String()
	userInfo, err := l.svcCtx.UserRpc.FindVolumeByIdentity(l.ctx, &user.FindVolumeReq{Identity: userIdentity})
	if err != nil {
		return nil, err
	}
	if repositoryPoolInfo.Size+userInfo.NowVolume > userInfo.TotalVolume {
		return nil, errors.New("文件超出容量限制！")
	}
	// 更新当前容量
	_, err = l.svcCtx.UserRpc.AddVolume(l.ctx, &user.AddVolumeReq{
		Identity: userIdentity,
		Size:     repositoryPoolInfo.Size,
	})
	if err != nil {
		return nil, err
	}
	// 新增关联记录
	newIdentity := utils.GenerateUUID()
	_, err = l.svcCtx.UserRepositoryModel.Insert(l.ctx, &model.UserRepository{
		Identity:           sql.NullString{String: newIdentity, Valid: true},
		UserIdentity:       sql.NullString{String: userIdentity, Valid: true},
		ParentId:           sql.NullInt64{Int64: req.ParentId, Valid: true},
		RepositoryIdentity: sql.NullString{String: req.RepositoryIdentity, Valid: true},
		Name:               sql.NullString{String: req.Name, Valid: true},
	})
	if err != nil {
		return nil, errors.New("存储失败！")
	}
	return &types.UserRepositorySaveResponse{}, nil
}
