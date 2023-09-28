package user

import (
	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"
	"TodoList/common/crypt"
	"TodoList/common/errorz"
	"TodoList/common/globalkey"
	"TodoList/model"
	"context"
	"github.com/pkg/errors"
	"html"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (*types.LoginResponse, error) {
	//转义过滤持久型XSS
	req.Name = html.EscapeString(req.Name)

	//查询内存中向该邮箱发送验证码的行为是否存活
	captcha, _ := l.svcCtx.Redis.Get(l.ctx, globalkey.Email(req.Email)).Result()
	if captcha != req.Captcha {
		return nil, errors.Wrapf(errorz.NewErrCode(errorz.CAPTCHA_VALIDATE_ERROR), "email: %v", req.Email)
	}

	//创建新用户
	pwd, err := crypt.EncryptBcrypt(req.Password)
	if err != nil {
		return nil, errorz.NewErrMsg("密码加密出错")
	}
	user := &model.User{Email: req.Email, Name: req.Name, Password: pwd}
	err = l.svcCtx.UserModel.Insert(l.ctx, nil, user)
	if err != nil {
		return nil, errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "email: %v,err: %+v", req.Email, err)
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	token, err := getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, user.Id)
	if err != nil {
		return nil, errors.Wrapf(errorz.NewErrMsg("token生成失败"), "GenerateToken userId : %d", user.Id)
	}

	return &types.LoginResponse{
		AccessToken:  token,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
		Name:         user.Name,
	}, nil
}
