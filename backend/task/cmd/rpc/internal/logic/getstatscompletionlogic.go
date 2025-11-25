package logic

import (
	"context"

	"jike_todo/task/cmd/rpc/internal/svc"
	"jike_todo/task/cmd/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatsCompletionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStatsCompletionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatsCompletionLogic {
	return &GetStatsCompletionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStatsCompletionLogic) GetStatsCompletion(in *task.GetStatsRequest) (*task.GetStatsCompletionResponse, error) {
	// todo: add your logic here and delete this line

	return &task.GetStatsCompletionResponse{}, nil
}
