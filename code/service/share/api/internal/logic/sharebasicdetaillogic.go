package logic

import (
	"butane-netdisk/common/errorx"
	"butane-netdisk/common/utils"
	"butane-netdisk/service/repository/rpc/types/repository"
	"butane-netdisk/service/share/api/internal/svc"
	"butane-netdisk/service/share/api/internal/types"
	"butane-netdisk/service/user_repository/rpc/types/userRepository"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/redis"

	jsoniter "github.com/json-iterator/go"

	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	//先查询缓存中有没有该数据
	redisQueryKey := utils.CacheShareKey + strconv.FormatInt(req.Id, 10)
	ifExists, err := l.svcCtx.RedisClient.Exists(redisQueryKey)
	if err != nil {
		return nil, err
	}
	if ifExists == true {
		//有
		jsonStr, err := l.svcCtx.RedisClient.Get(redisQueryKey)
		if err != nil {
			return nil, err
		}
		//判断数据是否为空
		if jsonStr == "" {
			return nil, errorx.NewCodeError(100, "查无此分享信息")
		}
		var shareInfo types.DetailResponse
		err = json.UnmarshalFromString(jsonStr, &shareInfo)
		if err != nil {
			return nil, err
		}
		//增加点击数
		err = l.svcCtx.ShareBasicModel.AddOneClick(l.ctx, req.Id)
		if err != nil {
			return nil, errorx.NewDefaultError("增加点击数失败！")
		}
		return &shareInfo, nil
	}
	//从数据库查询数据
	//申请分布式锁，获取repositoryId和返回user库、repositoryPool库的相应值
	redisLockKey := redisQueryKey
	redisLock := redis.NewRedisLock(l.svcCtx.RedisClient, redisLockKey)
	redisLock.SetExpire(utils.RedisLockExpireSeconds)
	if ok, err := redisLock.Acquire(); !ok || err != nil {
		return nil, errorx.NewCodeError(100, "当前有其他用户正在进行操作，请稍后重试")
	}
	defer func() {
		recover()
		redisLock.Release()
	}()
	shareInfo, err := l.svcCtx.ShareBasicModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
		break
	case sqlc.ErrNotFound:
		//缓存空数据
		err = l.svcCtx.RedisClient.Setex(redisQueryKey, "", utils.RedisLockExpireSeconds)
		if err != nil {
			return nil, err
		}
		return nil, errorx.NewCodeError(100, "查无此分享信息")
	default:
		return nil, err
	}
	userRepositoryName, err := l.svcCtx.UserRepositoryRpc.GetUserRepositoryNameByRepositoryId(l.ctx, &userRepository.RepositoryIdReq{
		RepositoryId: shareInfo.RepositoryId,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("无法获得用户储存库的信息！")
	}
	RepositoryPool, err := l.svcCtx.RepositoryRpc.GetRepositoryPoolByRepositoryId(l.ctx, &repository.RepositoryReq{
		RepositoryId: shareInfo.RepositoryId,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("无法获得储存池的信息！")
	}
	//把数据存储到缓存中
	DetailInfo := types.DetailResponse{
		RepositoryId: shareInfo.RepositoryId,
		Name:         userRepositoryName.RepositoryName,
		Ext:          RepositoryPool.Ext,
		Size:         RepositoryPool.Size,
		Path:         RepositoryPool.Path,
	}
	jsonStr, err := json.MarshalToString(DetailInfo)
	if err != nil {
		return nil, err
	}
	l.svcCtx.RedisClient.Setex(redisQueryKey, jsonStr, utils.RedisLockExpireSeconds)
	//增加点击数
	err = l.svcCtx.ShareBasicModel.AddOneClick(l.ctx, req.Id)
	if err != nil {
		return nil, errorx.NewDefaultError("增加点击数失败！")
	}
	return &DetailInfo, nil
}
