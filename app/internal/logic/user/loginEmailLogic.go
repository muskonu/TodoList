package user

import (
	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"
	"TodoList/common/errorz"
	"TodoList/common/globalkey"
	"context"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginEmailLogic {
	return &LoginEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginEmailLogic) LoginEmail(req *types.EmailCaptchaLoginRequest) (resp *types.LoginResponse, err error) {
	//查询内存中向该邮箱发送验证码的行为是否存活
	captcha, _ := l.svcCtx.Redis.Get(l.ctx, globalkey.Email(req.Email)).Result()
	if captcha != req.Captcha {
		return nil, errors.Wrapf(errorz.NewErrCode(errorz.CAPTCHA_VALIDATE_ERROR), "email: %v,err: %+v", req.Email)
	}

	//查询userid
	userid, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil {
		return nil, errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "email: %v,err: %+v", req.Email, err)
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	token, err := getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, userid.Id)
	if err != nil {
		return nil, errors.Wrapf(errorz.NewErrMsg("token生成失败"), "GenerateToken userId : %d", userid.Id)
	}

	return &types.LoginResponse{
		AccessToken:  token,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
		Name:         userid.Name,
	}, nil

}
