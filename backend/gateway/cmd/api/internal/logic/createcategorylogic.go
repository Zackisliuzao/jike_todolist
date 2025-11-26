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

type CreateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Create category
func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCategoryLogic) CreateCategory(req *types.CreateCategoryReq) (resp *types.BaseResponse, err error) {
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

	// 调用Task RPC创建分类
	rpcResp, err := l.svcCtx.TaskRpc.CreateCategory(l.ctx, &task.CreateCategoryRequest{
		UserId: userId,
		Name:   req.Name,
		Color:  req.Color,
	})
	if err != nil {
		l.Errorf("调用Task RPC创建分类失败: %v", err)
		resp = &types.BaseResponse{
			Code:    500,
			Message: "创建分类失败",
		}
		return
	}

	// 构建分类信息响应
	categoryInfo := &types.CategoryInfo{
		ID:        rpcResp.Id,
		UserID:    rpcResp.UserId,
		Name:      rpcResp.Name,
		Color:     rpcResp.Color,
		CreatedAt: rpcResp.CreatedAt,
		UpdatedAt: rpcResp.UpdatedAt,
	}

	resp = &types.BaseResponse{
		Code:    200,
		Message: "创建分类成功",
		Data:    categoryInfo,
	}
	return
}
