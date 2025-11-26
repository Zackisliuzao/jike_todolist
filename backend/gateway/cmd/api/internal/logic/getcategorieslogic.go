// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"jike_todo/common/ctxdata"
	"jike_todo/gateway/cmd/api/internal/svc"
	"jike_todo/gateway/cmd/api/internal/types"
	"jike_todo/task/cmd/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get category list
func NewGetCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoriesLogic {
	return &GetCategoriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoriesLogic) GetCategories() (resp *types.BaseResponse, err error) {
	// 从JWT上下文中获取用户ID
	userId := ctxdata.GetUidFromCtx(l.ctx)
	if userId == 0 {
		l.Error("无法获取用户ID")
		resp = &types.BaseResponse{
			Code:    401,
			Message: "用户未认证",
		}
		return
	}

	// 调用Task RPC获取分类列表
	rpcResp, err := l.svcCtx.TaskRpc.GetCategories(l.ctx, &task.GetCategoriesRequest{
		UserId: userId,
	})
	if err != nil {
		l.Errorf("调用Task RPC获取分类列表失败: %v", err)
		resp = &types.BaseResponse{
			Code:    500,
			Message: "获取分类列表失败",
		}
		return
	}

	// 转换RPC响应为API响应
	categories := make([]*types.CategoryInfo, 0, len(rpcResp.Categories))
	for _, category := range rpcResp.Categories {
		categories = append(categories, &types.CategoryInfo{
			ID:        category.Id,
			UserID:    category.UserId,
			Name:      category.Name,
			Color:     category.Color,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		})
	}

	resp = &types.BaseResponse{
		Code:    200,
		Message: "获取分类列表成功",
		Data:    categories,
	}
	return
}
