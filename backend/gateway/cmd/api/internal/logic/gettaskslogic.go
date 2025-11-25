// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"jike_todo/gateway/cmd/api/internal/svc"
	"jike_todo/gateway/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTasksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get task list
func NewGetTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTasksLogic {
	return &GetTasksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTasksLogic) GetTasks(req *types.GetTasksReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
