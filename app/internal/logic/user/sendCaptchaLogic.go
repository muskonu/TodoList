package user

import (
	"TodoList/app/internal/logic/email"
	"TodoList/common/errorz"
	"context"
	"github.com/pkg/errors"

	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCaptchaLogic {
	return &SendCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCaptchaLogic) SendCaptcha(req *types.CaptchaRequest) error {
	sendCaptcha := email.NewSendCaptcha(l.svcCtx.Redis)
	err := sendCaptcha(req.Email)
	if err != nil {
		return errors.Wrapf(errorz.NewErrCode(errorz.EMAIL_SEND_ERROR), "email:%s,err:%v", req.Email, err)
	}
	return nil
}
