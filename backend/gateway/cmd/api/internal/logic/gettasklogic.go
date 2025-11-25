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

type GetTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get task detail
func NewGetTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskLogic {
	return &GetTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskLogic) GetTask(taskIdStr string) (resp *types.BaseResponse, err error) {
	// 从JWT上下文中获取用户ID
	userId := ctxdata.GetUidFromCtx(l.ctx)

	// 解析任务ID
	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	if err != nil {
		l.Errorf("无效的任务ID: %s", taskIdStr)
		resp = &types.BaseResponse{
			Code:    400,
			Message: "无效的任务ID",
		}
		return
	}

	// 调用Task RPC获取任务详情
	rpcResp, err := l.svcCtx.TaskRpc.GetTask(l.ctx, &task.GetTaskRequest{
		TaskId: taskId,
		UserId: userId,
	})
	if err != nil {
		l.Errorf("调用Task RPC获取任务详情失败: %v", err)
		resp = &types.BaseResponse{
			Code:    500,
			Message: "获取任务详情失败",
		}
		return
	}

	// 构建任务信息响应
	taskInfo := &types.TaskInfo{
		ID:          rpcResp.Id,
		UserID:      rpcResp.UserId,
		Title:       rpcResp.Title,
		Description: rpcResp.Description,
		Category:    rpcResp.Category,
		Priority:    rpcResp.Priority,
		Status:      rpcResp.Status,
		DueDate:     rpcResp.DueDate,
		CreatedAt:   rpcResp.CreatedAt,
		UpdatedAt:   rpcResp.UpdatedAt,
		CompletedAt: rpcResp.CompletedAt,
	}

	resp = &types.BaseResponse{
		Code:    200,
		Message: "获取任务详情成功",
		Data:    taskInfo,
	}
	return
}
