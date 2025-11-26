// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"strconv"

	"jike_todo/common/ctxdata"
	"jike_todo/gateway/cmd/api/internal/svc"
	"jike_todo/gateway/cmd/api/internal/types"
	"jike_todo/task/cmd/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update category
func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(categoryIdStr string, req *types.UpdateCategoryReq) (resp *types.BaseResponse, err error) {
	// 解析分类ID
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		l.Error("无效的分类ID格式")
		resp = &types.BaseResponse{
			Code:    400,
			Message: "无效的分类ID",
		}
		return
	}

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

	// 调用Task RPC更新分类
	rpcResp, err := l.svcCtx.TaskRpc.UpdateCategory(l.ctx, &task.UpdateCategoryRequest{
		Id:     categoryId,
		UserId: userId,
		Name:   req.Name,
		Color:  req.Color,
	})
	if err != nil {
		l.Errorf("调用Task RPC更新分类失败: %v", err)
		resp = &types.BaseResponse{
			Code:    500,
			Message: "更新分类失败",
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
		Message: "更新分类成功",
		Data:    categoryInfo,
	}
	return
}
