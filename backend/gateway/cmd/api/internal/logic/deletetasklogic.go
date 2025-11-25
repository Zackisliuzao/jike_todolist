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

type DeleteTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete task
func NewDeleteTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTaskLogic {
	return &DeleteTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTaskLogic) DeleteTask(taskIdStr string) (resp *types.BaseResponse, err error) {
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

	// 调用Task RPC删除任务
	rpcResp, err := l.svcCtx.TaskRpc.DeleteTask(l.ctx, &task.DeleteTaskRequest{
		TaskId: taskId,
		UserId: userId,
	})
	if err != nil {
		l.Errorf("调用Task RPC删除任务失败: %v", err)
		resp = &types.BaseResponse{
			Code:    500,
			Message: "删除任务失败",
		}
		return
	}

	resp = &types.BaseResponse{
		Code:    200,
		Message: "删除任务成功",
		Data: map[string]interface{}{
			"success": rpcResp.Success,
		},
	}
	return
}
