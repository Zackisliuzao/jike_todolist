// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"jike_todo/gateway/cmd/api/internal/svc"
	"jike_todo/gateway/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatsCompletionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get completion statistics
func NewGetStatsCompletionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatsCompletionLogic {
	return &GetStatsCompletionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStatsCompletionLogic) GetStatsCompletion() (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
