// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"jike_todo/gateway/cmd/api/internal/svc"
	"jike_todo/gateway/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTasksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Search tasks
func NewSearchTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTasksLogic {
	return &SearchTasksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchTasksLogic) SearchTasks(req *types.SearchTasksReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
