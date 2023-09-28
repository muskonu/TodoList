package todo

import (
	"TodoList/common/errorz"
	"context"
	"github.com/pkg/errors"

	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTodoLogic {
	return &DelTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelTodoLogic) DelTodo(req *types.DelTodoRequest) error {
	todo, err := l.svcCtx.TodoListModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "todo: %v, err: %+v", req.Id, err)
	}

	err = l.svcCtx.TodoListModel.Delete(l.ctx, nil, req.Id)
	if err != nil {
		return errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "todo: %v, err: %+v", req.Id, err)
	}

	l.svcCtx.Schedule.RemoveJob(todo)
	return nil
}
