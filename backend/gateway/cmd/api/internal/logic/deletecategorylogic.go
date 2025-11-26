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

type DeleteCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete category
func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(categoryIdStr string) (resp *types.BaseResponse, err error) {
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

	// 调用Task RPC删除分类
	rpcResp, err := l.svcCtx.TaskRpc.DeleteCategory(l.ctx, &task.DeleteCategoryRequest{
		Id:     categoryId,
		UserId: userId,
	})
	if err != nil {
		l.Errorf("调用Task RPC删除分类失败: %v", err)
		resp = &types.BaseResponse{
			Code:    500,
			Message: "删除分类失败",
		}
		return
	}

	if rpcResp.Success {
		resp = &types.BaseResponse{
			Code:    200,
			Message: "删除分类成功",
		}
	} else {
		resp = &types.BaseResponse{
			Code:    500,
			Message: "删除分类失败",
		}
	}
	return
}
