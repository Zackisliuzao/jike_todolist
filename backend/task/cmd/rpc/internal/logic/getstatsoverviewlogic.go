package logic

import (
	"context"

	"jike_todo/task/cmd/rpc/internal/svc"
	"jike_todo/task/cmd/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatsOverviewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStatsOverviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatsOverviewLogic {
	return &GetStatsOverviewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Statistics operations
func (l *GetStatsOverviewLogic) GetStatsOverview(in *task.GetStatsRequest) (*task.GetStatsOverviewResponse, error) {
	// todo: add your logic here and delete this line

	return &task.GetStatsOverviewResponse{}, nil
}
