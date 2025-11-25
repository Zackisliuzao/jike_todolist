package logic

import (
	"context"
	"fmt"
	"time"

	"jike_todo/common/tool"
	"jike_todo/task/cmd/rpc/internal/svc"
	"jike_todo/task/cmd/rpc/task"
	"jike_todo/task/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskLogic {
	return &GetTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskLogic) GetTask(in *task.GetTaskRequest) (*task.Task, error) {
	// 查询任务
	taskRecord, err := l.svcCtx.TaskModel.FindOne(l.ctx, in.TaskId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, fmt.Errorf("任务不存在")
		}
		l.Errorf("查询任务失败: %v", err)
		return nil, fmt.Errorf("获取任务失败")
	}

	// 验证任务所有权
	if taskRecord.UserId != in.UserId {
		return nil, fmt.Errorf("无权限访问此任务")
	}

	// 构建响应
	return &task.Task{
		Id:          taskRecord.Id,
		UserId:      taskRecord.UserId,
		Title:       taskRecord.Title,
		Description: taskRecord.Description.String,
		Category:    taskRecord.Category,
		Priority:    int32(taskRecord.Priority),
		Status:      int32(taskRecord.Status),
		DueDate:     tool.FormatNullableDate(taskRecord.DueDate),
		CreatedAt:   taskRecord.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   taskRecord.UpdatedAt.Format(time.RFC3339),
		CompletedAt: tool.FormatNullableTime(taskRecord.CompletedAt),
	}, nil
}
