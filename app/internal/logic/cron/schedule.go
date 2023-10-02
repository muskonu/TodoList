package cron

import (
	"TodoList/model"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

type TodoSchedule struct {
	cancel *sync.Map
	cron   *cron.Cron
}

func NewTodoSchedule() *TodoSchedule {
	var cancelMap *sync.Map
	t := &TodoSchedule{cancel: cancelMap, cron: cron.New(cron.WithSeconds())}
	t.cron.Start()
	return t
}

func (s *TodoSchedule) AddJob(todo *model.TodoList, email string) error {
	if !todo.DueDate.Add(-15 * time.Minute).After(time.Now()) {
		return nil
	}
	j := NewTodoJob(todo, email)
	c := make(chan struct{}, 1)
	s.cancel.Store(j.list.Id, c)
	switch j.list.Recurrence {
	case 0:
		go func() {
			timer := time.After(j.list.DueDate.Add(-15 * time.Minute).Sub(time.Now()))
			select {
			case <-timer:
				j.Run()
			case <-c:
				return
			}
		}()
	default:
		cmd, err := s.cron.AddJob(j.ParseToCrontab(), j)
		if err != nil {
			return err
		}
		s.cancel.Store(j.list.Id, cmd)
	}
	return nil
}

func (s *TodoSchedule) RemoveJob(list *model.TodoList) {
	switch list.Recurrence {
	case 0:
		channel, ok := s.cancel.Load(list.Id)
		if !ok {
			return
		}
		channel.(chan struct{}) <- struct{}{}
		s.cancel.Delete(list.Id)
	default:
		cmd, ok := s.cancel.Load(list.Id)
		if !ok {
			return
		}
		s.cron.Remove(cmd.(cron.EntryID))
		s.cancel.Delete(list.Id)
	}

}
