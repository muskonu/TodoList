package cron

import (
	"TodoList/model"
	"github.com/robfig/cron/v3"
	"time"
)

type TodoSchedule struct {
	onceCancel   map[int64]chan struct{}
	repeatCancel map[int64]cron.EntryID
	cron         *cron.Cron
}

func NewTodoSchedule() *TodoSchedule {
	t := &TodoSchedule{onceCancel: make(map[int64]chan struct{}), repeatCancel: make(map[int64]cron.EntryID), cron: cron.New(cron.WithSeconds())}
	t.cron.Start()
	return t
}

func (s *TodoSchedule) AddJob(todo *model.TodoList, email string) error {
	if !todo.DueDate.Add(-15 * time.Minute).After(time.Now()) {
		return nil
	}
	j := NewTodoJob(todo, email)
	c := make(chan struct{}, 1)
	s.onceCancel[j.list.Id] = c
	switch j.list.Recurrence {
	case 0:
		go func() {
			timer := time.After(j.list.DueDate.Add(-15 * time.Minute).Sub(time.Now()))
			select {
			case <-timer:
				j.Run()
			case <-s.onceCancel[j.list.Id]:
				return
			}
		}()
	default:
		cmd, err := s.cron.AddJob(j.ParseToCrontab(), j)
		if err != nil {
			return err
		}
		s.repeatCancel[j.list.Id] = cmd
	}
	return nil
}

func (s *TodoSchedule) RemoveJob(list *model.TodoList) {
	switch list.Recurrence {
	case 0:
		select {
		case s.onceCancel[list.Id] <- struct{}{}:
			delete(s.onceCancel, list.Id)
			return
		default:
			return
		}
	default:
		s.cron.Remove(s.repeatCancel[list.Id])
		delete(s.repeatCancel, list.Id)
	}

}
