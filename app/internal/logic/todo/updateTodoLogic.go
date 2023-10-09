package todo

import (
	"TodoList/common/ctxdata"
	"TodoList/common/errorz"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTodoLogic {
	return &UpdateTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTodoLogic) UpdateTodo(req *types.UpdateTodoRequest) error {

	id := ctxdata.GetUidFromCtx(l.ctx)
	err := l.svcCtx.TodoListModel.Transaction(l.ctx, func(db *gorm.DB) error {
		//处理数据库
		old, err := l.svcCtx.TodoListModel.FindOne(l.ctx, req.Id)
		if err != nil {
			return errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "todo: %v, err: %+v", req.Id, err)
		}
		old.Content = req.Content
		old.DueDate = time.Unix(req.DueDate, 0)
		old.Recurrence = req.Recurrence
		err = l.svcCtx.TodoListModel.Update(l.ctx, db, old)
		if err != nil {
			return errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "todo: %v, err: %+v", req.Id, err)
		}
		//修改定时任务
		user, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
		if err != nil {
			return errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "todo: %v, err: %+v", id, err)
		}
		l.svcCtx.Schedule.RemoveJob(old)
		err = l.svcCtx.Schedule.AddJob(old, user.Email)
		if err != nil {
			return errors.Wrapf(errorz.NewErrCode(errorz.OPERATE_CRON_ERROR), "todo: %v, err: %+v", id, err)
		}
		return nil
	})

	return err
}
