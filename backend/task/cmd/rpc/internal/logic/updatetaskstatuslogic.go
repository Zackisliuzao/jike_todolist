package logic

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"jike_todo/common/tool"
	"jike_todo/task/cmd/rpc/internal/svc"
	"jike_todo/task/cmd/rpc/task"
	"jike_todo/task/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTaskStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskStatusLogic {
	return &UpdateTaskStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTaskStatusLogic) UpdateTaskStatus(in *task.UpdateTaskStatusRequest) (*task.Task, error) {
	// 查询任务
	taskRecord, err := l.svcCtx.TaskModel.FindOne(l.ctx, in.TaskId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, fmt.Errorf("任务不存在")
		}
		l.Errorf("查询任务失败: %v", err)
		return nil, fmt.Errorf("更新任务状态失败")
	}

	// 验证任务所有权
	if taskRecord.UserId != in.UserId {
		return nil, fmt.Errorf("无权限更新此任务")
	}

	// 更新任务状态
	taskRecord.Status = int64(in.Status)

	// 如果任务完成，设置完成时间
	if in.Status == 1 { // 已完成
		taskRecord.CompletedAt = sql.NullTime{Time: time.Now(), Valid: true}
	} else {
		taskRecord.CompletedAt = sql.NullTime{Valid: false}
	}

	err = l.svcCtx.TaskModel.Update(l.ctx, taskRecord)
	if err != nil {
		l.Errorf("更新任务状态失败: %v", err)
		return nil, fmt.Errorf("更新任务状态失败")
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
