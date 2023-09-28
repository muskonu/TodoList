package svc

import (
	"TodoList/app/internal/config"
	"TodoList/app/internal/logic/cron"
	"TodoList/model"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

type ServiceContext struct {
	Config        config.Config
	Validate      *validator.Validate
	Redis         *redis.Client
	UserModel     model.UserModel
	TodoListModel model.TodoListModel
	Schedule      *cron.TodoSchedule
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := gormc.ConnectMysql(c.Mysql)
	return &ServiceContext{
		Config:        c,
		Validate:      validator.New(),
		Redis:         redis.NewClient(&redis.Options{Addr: c.Redis.Host, Password: c.Redis.Password}),
		UserModel:     model.NewUserModel(db, c.CacheRedis),
		TodoListModel: model.NewTodoListModel(db, c.CacheRedis),
		Schedule:      cron.NewTodoSchedule(),
	}
}
