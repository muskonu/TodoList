package todo

import (
	"TodoList/common/ctxdata"
	"TodoList/common/errorz"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTodoLogic {
	return &QueryTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryTodoLogic) QueryTodo(req *types.QueryTodoRequest) (resp *types.QueryTodoResponse, err error) {
	//参数矫正
	if req.PageSize < 1 || req.PageSize > 20 {
		req.PageSize = 10
	}
	if req.Cursor < 0 {
		req.Cursor = 0
	}
	//进入数据库查询
	userId := ctxdata.GetUidFromCtx(l.ctx)
	todos, err := l.svcCtx.TodoListModel.FindId(l.ctx, userId, req.IsCompleted, req.Cursor, req.PageSize)
	if err != nil {
		return nil, errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "user: %v,addtodo err: %+v", userId, err)
	}
	for index, todo := range todos {
		todos[index], _ = l.svcCtx.TodoListModel.FindOne(l.ctx, todo.Id)
	}
	var res []types.Todo
	err = copier.Copy(&res, todos)
	if err != nil {
		return nil, errorz.NewErrMsg("copy发生错误")
	}
	return &types.QueryTodoResponse{Todos: res}, nil
}
