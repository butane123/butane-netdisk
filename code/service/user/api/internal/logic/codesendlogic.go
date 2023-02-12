package logic

import (
	"butane-netdisk/common/utils"
	"context"
	"net/smtp"

	"github.com/jordan-wright/email"

	"butane-netdisk/service/user/api/internal/svc"
	"butane-netdisk/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CodeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCodeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CodeSendLogic {
	return &CodeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CodeSendLogic) CodeSend(req *types.CodeSendRequest) (resp *types.CodeSendResponse, err error) {
	e := email.NewEmail()
	e.From = "VerificationCode by butane-netdisk <" + utils.ServerEmail + ">"
	e.To = []string{req.Email}
	e.Subject = "This is a VerificationCode!"
	verificationCode := utils.GenerateVerificationCode()
	e.Text = []byte(verificationCode)
	e.Send(utils.EmailSmtpAddr, smtp.PlainAuth("", utils.ServerEmail, utils.EmailAuthCode, utils.EmailSmtpHost))
	l.svcCtx.RedisClient.Setex(utils.CacheEmailCodeKey+req.Email, verificationCode, utils.EmailCodeExpireSeconds)
	return &types.CodeSendResponse{}, nil
}
