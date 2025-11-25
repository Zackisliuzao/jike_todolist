package svc

import (
	"jike_todo/task/cmd/rpc/internal/config"
	"jike_todo/task/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	TaskModel model.TasksModel
	CateModel model.CategoriesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		TaskModel: model.NewTasksModel(sqlx.NewMysql(c.DataSource), c.Cache),
		CateModel: model.NewCategoriesModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
