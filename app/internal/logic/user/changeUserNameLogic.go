package user

import (
	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"
	"TodoList/common/ctxdata"
	"TodoList/common/errorz"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeUserNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeUserNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeUserNameLogic {
	return &ChangeUserNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeUserNameLogic) ChangeUserName(req *types.UserNameRequest) (resp *types.UserNameResponse, err error) {
	id := ctxdata.GetUidFromCtx(l.ctx)
	u, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
	if err != nil {
		return nil, errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "userid: %v,err: %+v", id, err)
	}

	u.Name = req.Name
	err = l.svcCtx.UserModel.Update(l.ctx, nil, u)
	if err != nil {
		return nil, errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "userid: %v,err: %+v", id, err)
	}

	return &types.UserNameResponse{u.Name}, nil
}
