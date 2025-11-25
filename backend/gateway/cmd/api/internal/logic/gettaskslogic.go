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

type GetTasksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get task list
func NewGetTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTasksLogic {
	return &GetTasksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTasksLogic) GetTasks(req *types.GetTasksReq) (resp *types.BaseResponse, err error) {
	// 从JWT上下文中获取用户ID
	userId := ctxdata.GetUidFromCtx(l.ctx)

	// 调用Task RPC获取任务列表
	rpcResp, err := l.svcCtx.TaskRpc.GetTasks(l.ctx, &task.GetTasksRequest{
		UserId:  userId,
		Page:    req.Page,
		Size:    req.Size,
		Category: req.Category,
		Status:  req.Status,
		SortBy:  req.SortBy,
	})
	if err != nil {
		l.Errorf("调用Task RPC获取任务列表失败: %v", err)
		resp = &types.BaseResponse{
			Code:    500,
			Message: "获取任务列表失败",
		}
		return
	}

	// 转换任务列表
	var taskList []types.TaskInfo
	for _, t := range rpcResp.Tasks {
		taskInfo := types.TaskInfo{
			ID:          t.Id,
			UserID:      t.UserId,
			Title:       t.Title,
			Description: t.Description,
			Category:    t.Category,
			Priority:    t.Priority,
			Status:      t.Status,
			DueDate:     t.DueDate,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
			CompletedAt: t.CompletedAt,
		}
		taskList = append(taskList, taskInfo)
	}

	// 构建响应
	getTasksResp := types.GetTasksResp{
		Tasks: taskList,
		Total: rpcResp.Total,
	}

	resp = &types.BaseResponse{
		Code:    200,
		Message: "获取任务列表成功",
		Data:    getTasksResp,
	}
	return
}
