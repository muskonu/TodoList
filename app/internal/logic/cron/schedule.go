package cron

import (
	"TodoList/model"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

type TodoSchedule struct {
	emailCancel sync.Map
	daoCancel   sync.Map
	model       model.TodoListModel
	cron        *cron.Cron
}

func NewTodoSchedule(model model.TodoListModel) *TodoSchedule {
	var emailCancelMap sync.Map
	var daoCancelMap sync.Map
	t := &TodoSchedule{daoCancel: daoCancelMap, emailCancel: emailCancelMap, cron: cron.New(cron.WithSeconds()), model: model}
	t.cron.Start()
	return t
}

func (s *TodoSchedule) AddJob(todo *model.TodoList, email string) error {
	j := NewTodoJob(todo, email, s.model)
	c := make(chan struct{}, 1)
	s.emailCancel.Store(j.list.Id, c)
	switch j.list.Recurrence {
	case 0:
		if !todo.DueDate.Add(-15 * time.Minute).After(time.Now()) {
			return nil
		}
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
		//事务添加email job和dao job
		err := func() (err error) {
			rollback := func(r func(cron.EntryID), id cron.EntryID) {
				if err != nil {
					r(id)
				}
			}
			emailCmd, err := s.cron.AddJob(j.emailJob.ParseToCrontab(), j.emailJob)
			if err != nil {
				return err
			}
			defer rollback(s.rollback, emailCmd)
			daoCmd, err := s.cron.AddJob(j.ParseToCrontab(), j)
			if err != nil {
				return err
			}
			s.emailCancel.Store(j.list.Id, emailCmd)
			s.daoCancel.Store(j.list.Id, daoCmd)
			return nil
		}()
		return err
	}
	return nil
}

func (s *TodoSchedule) RemoveJob(list *model.TodoList) {
	switch list.Recurrence {
	case 0:
		channel, ok := s.emailCancel.Load(list.Id)
		if !ok {
			return
		}
		channel.(chan struct{}) <- struct{}{}
		s.emailCancel.Delete(list.Id)
	default:
		emailCmd, ok := s.emailCancel.Load(list.Id)
		if !ok {
			return
		}
		daoCmd, _ := s.emailCancel.Load(list.Id)
		s.cron.Remove(emailCmd.(cron.EntryID))
		s.cron.Remove(daoCmd.(cron.EntryID))
		s.emailCancel.Delete(list.Id)
		s.daoCancel.Delete(list.Id)
	}
}

func (s *TodoSchedule) rollback(id cron.EntryID) {
	s.cron.Remove(id)
}
