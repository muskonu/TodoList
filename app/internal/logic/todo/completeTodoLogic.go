package todo

import (
	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"
	"TodoList/common/ctxdata"
	"TodoList/common/errorz"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CompleteTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompleteTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteTodoLogic {
	return &CompleteTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompleteTodoLogic) CompleteTodo(req *types.CompleteTodoRequest) error {
	id := ctxdata.GetUidFromCtx(l.ctx)
	err := l.svcCtx.TodoListModel.Transaction(l.ctx, func(db *gorm.DB) error {
		//更改数据库状态
		old, err := l.svcCtx.TodoListModel.FindOne(l.ctx, req.Id)
		if err != nil {
			return errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "todo: %v, err: %+v", req.Id, err)
		}
		err = l.svcCtx.TodoListModel.ChangeComplete(l.ctx, db, old)
		if err != nil {
			return errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "todo: %v, err: %+v", req.Id, err)
		}
		//更改提醒状态
		user, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
		if err != nil {
			return errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "user: %v, err: %+v", id, err)
		}
		if old.IsCompleted {
			l.svcCtx.Schedule.RemoveJob(old)
		} else {
			err = l.svcCtx.Schedule.AddJob(old, user.Email)
			if err != nil {
				return errors.Wrapf(errorz.NewErrCode(errorz.OPERATE_CRON_ERROR), "todo: %v, err: %+v", id, err)
			}
		}
		return nil
	})

	return err
}
