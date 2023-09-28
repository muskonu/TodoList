// Code generated by goctl. DO NOT EDIT!

package model

import (
	"TodoList/common/globalkey"
	"context"
	"fmt"
	"time"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheTodoListTodoListIdPrefix = "cache:todoList:todoList:id:"
)

type (
	todoListModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *TodoList) error

		FindOne(ctx context.Context, id int64) (*TodoList, error)
		Update(ctx context.Context, tx *gorm.DB, data *TodoList) error
		ChangeComplete(ctx context.Context, tx *gorm.DB, data *TodoList) error
		FindId(ctx context.Context, userId int64, isCompleted bool, cursor int64, pageSize int64) ([]*TodoList, error)

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultTodoListModel struct {
		gormc.CachedConn
		table string
	}

	TodoList struct {
		Id          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
		Content     string    `gorm:"column:content;not null" json:"content"`
		DueDate     time.Time `gorm:"column:due_date;not null" json:"due_date"`
		Recurrence  int64     `gorm:"column:recurrence;not null" json:"recurrence"`
		IsCompleted bool      `gorm:"column:is_completed;not null" json:"is_completed"`
		CreateAt    time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
		UpdateAt    time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP" json:"update_at"`
		UserId      int64     `gorm:"column:user_id;not null" json:"user_id"`
	}
)

func (TodoList) TableName() string {
	return "`todo_list`"
}

func newTodoListModel(conn *gorm.DB, c cache.CacheConf) *defaultTodoListModel {
	return &defaultTodoListModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`todo_list`",
	}
}

func (m *defaultTodoListModel) Insert(ctx context.Context, tx *gorm.DB, data *TodoList) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultTodoListModel) FindOne(ctx context.Context, id int64) (*TodoList, error) {
	todoListTodoListIdKey := fmt.Sprintf("%s%v", cacheTodoListTodoListIdPrefix, id)
	var resp TodoList
	err := m.QueryCtx(ctx, &resp, todoListTodoListIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&TodoList{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTodoListModel) FindId(ctx context.Context, userId int64, isCompleted bool, cursor int64, pageSize int64) ([]*TodoList, error) {
	var resp []*TodoList
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&TodoList{}).Select("id").Where("`user_id` = ? AND `is_completed` = ?", userId, isCompleted).
			Order("due_date desc").Limit(int(pageSize)).Offset(int(cursor)).Find(&resp).Error
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultTodoListModel) Update(ctx context.Context, tx *gorm.DB, data *TodoList) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultTodoListModel) ChangeComplete(ctx context.Context, tx *gorm.DB, data *TodoList) error {

	data.IsCompleted = !data.IsCompleted
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultTodoListModel) getCacheKeys(data *TodoList) []string {
	if data == nil {
		return []string{}
	}
	todoListTodoListIdKey := fmt.Sprintf("%s%v", cacheTodoListTodoListIdPrefix, data.Id)
	cacheKeys := []string{
		todoListTodoListIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultTodoListModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		if err == ErrNotFound {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&TodoList{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultTodoListModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultTodoListModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTodoListTodoListIdPrefix, primary)
}

func (m *defaultTodoListModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&TodoList{}).Where("`id` = ?", primary).Take(v).Error
}

func (t TodoList) Repeat() string {
	return globalkey.Recurrence(t.Recurrence)
}

func (t TodoList) Deadline() int64 {
	return t.DueDate.Unix()
}
