package logic

import (
	"context"

	"jike_todo/task/cmd/rpc/internal/svc"
	"jike_todo/task/cmd/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchUpdateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchUpdateStatusLogic {
	return &BatchUpdateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Batch operations
func (l *BatchUpdateStatusLogic) BatchUpdateStatus(in *task.BatchUpdateStatusRequest) (*task.BatchUpdateStatusResponse, error) {
	// todo: add your logic here and delete this line

	return &task.BatchUpdateStatusResponse{}, nil
}
