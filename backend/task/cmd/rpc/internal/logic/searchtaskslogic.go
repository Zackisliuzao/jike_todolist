package logic

import (
	"context"

	"jike_todo/task/cmd/rpc/internal/svc"
	"jike_todo/task/cmd/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTasksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTasksLogic {
	return &SearchTasksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Search functionality
func (l *SearchTasksLogic) SearchTasks(in *task.SearchTasksRequest) (*task.SearchTasksResponse, error) {
	// todo: add your logic here and delete this line

	return &task.SearchTasksResponse{}, nil
}
