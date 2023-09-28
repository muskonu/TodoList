package todo

import (
	"TodoList/common/ctxdata"
	"TodoList/common/errorz"
	"TodoList/model"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"TodoList/app/internal/svc"
	"TodoList/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTodoLogic {
	return &AddTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTodoLogic) AddTodo(req *types.AddTodoRequest) error {
	id := ctxdata.GetUidFromCtx(l.ctx)
	todo := &model.TodoList{UserId: id, Content: req.Content,
		Recurrence: req.Recurrence, DueDate: time.Unix(req.DueDate, 0), IsCompleted: false}

	//添加定时任务
	err := l.svcCtx.UserModel.Transaction(l.ctx, func(db *gorm.DB) error {
		user, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
		if err != nil {
			return errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "todo: %v, err: %+v", id, err)
		}
		//插入数据库
		err = l.svcCtx.TodoListModel.Insert(l.ctx, db, todo)
		if err != nil {
			return errors.Wrapf(errorz.NewErrCode(errorz.DB_ERROR), "user: %v,addtodo err: %+v", id, err)
		}
		err = l.svcCtx.Schedule.AddJob(todo, user.Email)
		if err != nil {
			return errors.Wrapf(errorz.NewErrCode(errorz.OPERATE_CRON_ERROR), "todo: %v, err: %+v", id, err)
		}
		return nil
	})
	return err
}
