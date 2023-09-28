package user

import (
	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"
	"TodoList/common/crypt"
	"TodoList/common/errorz"
	"context"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginPassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginPassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginPassLogic {
	return &LoginPassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginPassLogic) LoginPass(req *types.EmailPwdLoginRequest) (resp *types.LoginResponse, err error) {
	//查询userid
	userid, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil {
		return nil, errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "email: %v,err: %+v", req.Email, err)
	}
	if !crypt.ValidateBcrypt(userid.Password, req.Password) {
		return nil, errors.Wrapf(errorz.NewErrMsg("密码错误"), "Login userId : %d", userid.Id)
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
