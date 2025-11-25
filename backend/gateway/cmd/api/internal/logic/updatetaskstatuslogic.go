// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"jike_todo/gateway/cmd/api/internal/svc"
	"jike_todo/gateway/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update task status
func NewUpdateTaskStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskStatusLogic {
	return &UpdateTaskStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTaskStatusLogic) UpdateTaskStatus(req *types.UpdateTaskStatusReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
