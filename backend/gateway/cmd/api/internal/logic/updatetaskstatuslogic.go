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

func (l *UpdateTaskStatusLogic) UpdateTaskStatus(req *types.UpdateTaskStatusReq, taskIdStr string) (resp *types.BaseResponse, err error) {
	// 从JWT上下文中获取用户ID
	userId := ctxdata.GetUidFromCtx(l.ctx)

	// 解析任务ID
	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	if err != nil {
		l.Errorf("解析任务ID失败: %v", err)
		resp = &types.BaseResponse{
			Code:    400,
			Message: "无效的任务ID",
			Data:    nil,
		}
		return
	}

	// 调用RPC服务更新任务状态
	rpcResp, err := l.svcCtx.TaskRpc.UpdateTaskStatus(l.ctx, &task.UpdateTaskStatusRequest{
		TaskId: taskId,
		UserId: userId,
		Status: req.Status,
	})
	if err != nil {
		l.Errorf("调用RPC服务失败: %v", err)
		resp = &types.BaseResponse{
			Code:    500,
			Message: "更新任务状态失败: " + err.Error(),
			Data:    nil,
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
		Message: "更新任务状态成功",
		Data:    taskInfo,
	}
	return
}
