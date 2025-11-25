// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"jike_todo/gateway/cmd/api/internal/svc"
	"jike_todo/gateway/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatsOverviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get overview statistics
func NewGetStatsOverviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatsOverviewLogic {
	return &GetStatsOverviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStatsOverviewLogic) GetStatsOverview() (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
