// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"jike_todo/gateway/cmd/api/internal/config"
	"jike_todo/task/cmd/rpc/taskservice"
	"jike_todo/user/cmd/rpc/authservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc authservice.AuthService
	TaskRpc taskservice.TaskService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: authservice.NewAuthService(zrpc.MustNewClient(c.UserRpcClient)),
		TaskRpc: taskservice.NewTaskService(zrpc.MustNewClient(c.TaskRpcClient)),
	}
}
