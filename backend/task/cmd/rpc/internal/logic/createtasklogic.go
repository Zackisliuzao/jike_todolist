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

type CreateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Task CRUD operations
func (l *CreateTaskLogic) CreateTask(in *task.CreateTaskRequest) (*task.Task, error) {
	// 设置默认值
	if in.Category == "" {
		in.Category = "default"
	}
	if in.Priority == 0 {
		in.Priority = 1 // 默认低优先级
	}

	// 处理截止日期
	dueDate, err := tool.ParseNullableDate(in.DueDate)
	if err != nil {
		l.Errorf("无效的截止日期格式: %s, error: %v", in.DueDate, err)
		return nil, fmt.Errorf("截止日期格式无效，请使用 YYYY-MM-DD 格式")
	}

	// 创建任务记录
	taskRecord := &model.Tasks{
		UserId:      in.UserId,
		Title:       in.Title,
		Description: sql.NullString{String: in.Description, Valid: in.Description != ""},
		Category:    in.Category,
		Priority:    int64(in.Priority),
		Status:      0, // 默认未完成
		DueDate:     dueDate,
	}

	result, err := l.svcCtx.TaskModel.Insert(l.ctx, taskRecord)
	if err != nil {
		l.Errorf("插入任务失败: %v", err)
		return nil, fmt.Errorf("创建任务失败")
	}

	taskId, err := result.LastInsertId()
	if err != nil {
		l.Errorf("获取任务ID失败: %v", err)
		return nil, fmt.Errorf("创建任务失败")
	}

	// 查询刚创建的任务
	createdTask, err := l.svcCtx.TaskModel.FindOne(l.ctx, taskId)
	if err != nil {
		l.Errorf("查询创建的任务失败: %v", err)
		return nil, fmt.Errorf("创建任务失败")
	}

	// 构建响应
	return &task.Task{
		Id:          createdTask.Id,
		UserId:      createdTask.UserId,
		Title:       createdTask.Title,
		Description: createdTask.Description.String,
		Category:    createdTask.Category,
		Priority:    int32(createdTask.Priority),
		Status:      int32(createdTask.Status),
		DueDate:     tool.FormatNullableDate(createdTask.DueDate),
		CreatedAt:   createdTask.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   createdTask.UpdatedAt.Format(time.RFC3339),
		CompletedAt: tool.FormatNullableTime(createdTask.CompletedAt),
	}, nil
}
