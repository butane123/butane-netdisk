package logic

import (
	"cloud-disk/common/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"cloud-disk/service/user/api/internal/svc"
	"cloud-disk/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(req *types.RefreshAuthRequest, Authorization string) (resp *types.RefreshAuthResponse, err error) {
	//获得原token的剩余信息
	restClaims := make(jwt.MapClaims)
	judgeValid, err := jwt.ParseWithClaims(Authorization, restClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(l.svcCtx.Config.Auth.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	//判断是否token有效
	if !judgeValid.Valid {
		return nil, errors.New("Token已失效！")
	}
	//利用过期token的其他值，生成新token等信息
	nowTime := time.Now().Unix()
	expireTime := l.svcCtx.Config.Auth.AccessExpire
	newToken, err := utils.GenerateJwtToken(l.svcCtx.Config.Auth.AccessSecret, nowTime, expireTime, json.Number(fmt.Sprintf("%v", restClaims["userIdentity"])).String())
	if err != nil {
		return nil, err
	}
	return &types.RefreshAuthResponse{
		AccesssToken: newToken,
		AccessExpire: nowTime + expireTime,
		RefreshAfter: nowTime + expireTime/2,
	}, nil
}
