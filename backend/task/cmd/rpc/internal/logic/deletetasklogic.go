package logic

import (
	"context"
	"fmt"

	"jike_todo/task/cmd/rpc/internal/svc"
	"jike_todo/task/cmd/rpc/task"
	"jike_todo/task/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTaskLogic {
	return &DeleteTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTaskLogic) DeleteTask(in *task.DeleteTaskRequest) (*task.DeleteTaskResponse, error) {
	// 查询任务
	taskRecord, err := l.svcCtx.TaskModel.FindOne(l.ctx, in.TaskId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, fmt.Errorf("任务不存在")
		}
		l.Errorf("查询任务失败: %v", err)
		return nil, fmt.Errorf("删除任务失败")
	}

	// 验证任务所有权
	if taskRecord.UserId != in.UserId {
		return nil, fmt.Errorf("无权限删除此任务")
	}

	// 删除任务
	err = l.svcCtx.TaskModel.Delete(l.ctx, in.TaskId)
	if err != nil {
		l.Errorf("删除任务失败: %v", err)
		return nil, fmt.Errorf("删除任务失败")
	}

	return &task.DeleteTaskResponse{
		Success: true,
	}, nil
}
