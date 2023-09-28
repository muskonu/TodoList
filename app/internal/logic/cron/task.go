package cron

import (
	"TodoList/app/internal/logic/email"
	"TodoList/model"
	"fmt"
	"time"
)

type TodoJob struct {
	list  *model.TodoList
	email string
}

func (todo *TodoJob) Run() {
	_ = email.SendWarning(todo.email, todo.list)
}

func (todo *TodoJob) ParseToCrontab() string {
	due := todo.list.DueDate.Add(-15 * time.Minute)
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

func NewTodoJob(todo *model.TodoList, email string) *TodoJob {
	return &TodoJob{email: email, list: todo}
}
