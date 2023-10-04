package cron

import (
	"TodoList/app/internal/logic/email"
	"TodoList/common/globalkey"
	"TodoList/model"
	"context"
	"fmt"
	"time"
)

type TodoJob struct {
	*emailJob
	todoListModel model.TodoListModel
}

type emailJob struct {
	list  *model.TodoList
	email string
}

func (e *emailJob) Run() {
	_ = email.SendWarning(e.email, e.list)
}

func (t *TodoJob) Run() {
	t.list.DueDate = t.list.DueDate.Add(globalkey.Period(t.list.Recurrence))
	t.list.IsCompleted = false
	_ = t.todoListModel.Update(context.Background(), nil, t.list)
}

func (email *emailJob) ParseToCrontab() string {
	due := email.list.DueDate.Add(-15 * time.Minute)
	switch email.list.Recurrence {
	case 1:
		return fmt.Sprintf("%d %d %d * * *", due.Second(), due.Minute(), due.Hour())
	case 2:
		return fmt.Sprintf("%d %d %d * * %d", due.Second(), due.Minute(), due.Hour(), due.Weekday())
	case 3:
		return fmt.Sprintf("%d %d %d %d * *", due.Second(), due.Minute(), due.Hour(), due.Day())
	case 4:
		return fmt.Sprintf("%d %d %d %d %d *", due.Second(), due.Minute(), due.Hour(), due.Day(), due.Month())
	default:
		return ""
	}
}

func (todo *TodoJob) ParseToCrontab() string {
	due := todo.list.DueDate
	switch todo.list.Recurrence {
	case 1:
		return fmt.Sprintf("%d %d %d * * *", due.Second(), due.Minute(), due.Hour())
	case 2:
		return fmt.Sprintf("%d %d %d * * %d", due.Second(), due.Minute(), due.Hour(), due.Weekday())
	case 3:
		return fmt.Sprintf("%d %d %d %d * *", due.Second(), due.Minute(), due.Hour(), due.Day())
	case 4:
		return fmt.Sprintf("%d %d %d %d %d *", due.Second(), due.Minute(), due.Hour(), due.Day(), due.Month())
	default:
		return ""
	}
}

func NewTodoJob(todo *model.TodoList, email string, model model.TodoListModel) *TodoJob {
	return &TodoJob{&emailJob{email: email, list: todo}, model}
}
