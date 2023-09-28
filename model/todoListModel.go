package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ TodoListModel = (*customTodoListModel)(nil)

type (
	// TodoListModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTodoListModel.
	TodoListModel interface {
		todoListModel
		customTodoListLogicModel
	}

	customTodoListModel struct {
		*defaultTodoListModel
	}

	customTodoListLogicModel interface {
	}
)

// NewTodoListModel returns a model for the database table.
func NewTodoListModel(conn *gorm.DB, c cache.CacheConf) TodoListModel {
	return &customTodoListModel{
		defaultTodoListModel: newTodoListModel(conn, c),
	}
}

func (m *defaultTodoListModel) customCacheKeys(data *TodoList) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
