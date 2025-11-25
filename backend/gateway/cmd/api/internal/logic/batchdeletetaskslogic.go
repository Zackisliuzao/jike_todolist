// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"jike_todo/gateway/cmd/api/internal/svc"
	"jike_todo/gateway/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteTasksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Batch delete tasks
func NewBatchDeleteTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteTasksLogic {
	return &BatchDeleteTasksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchDeleteTasksLogic) BatchDeleteTasks(req *types.BatchDeleteReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
