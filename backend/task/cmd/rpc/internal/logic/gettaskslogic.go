package logic

import (
	"context"

	"jike_todo/task/cmd/rpc/internal/svc"
	"jike_todo/task/cmd/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTasksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTasksLogic {
	return &GetTasksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTasksLogic) GetTasks(in *task.GetTasksRequest) (*task.GetTasksResponse, error) {
	// 由于go-zero生成的模型限制，这里实现一个简化的分页查询
	// 实际项目中可以：
	// 1. 扩展模型添加自定义查询方法
	// 2. 使用原生SQL查询
	// 3. 创建自定义数据访问层

	var tasks []*task.Task
	var total int32

	// 这里先通过遍历可能的任务ID范围来查询用户任务
	// 简化实现：假设用户任务ID范围，实际项目中需要优化
	maxTaskId := int64(1000) // 假设的最大任务ID

	for i := int64(1); i <= maxTaskId; i++ {
		taskModel, err := l.svcCtx.TaskModel.FindOne(l.ctx, i)
		if err != nil {
			// 如果任务不存在，继续下一个
			continue
			}

		// 检查是否属于当前用户
		if taskModel.UserId == in.UserId {
			// 应用筛选条件
			if in.Category != "" && taskModel.Category != in.Category {
				continue
			}
			if in.Status != 0 && int32(taskModel.Status) != in.Status {
				continue
			}

			taskPB := &task.Task{
				Id:          taskModel.Id,
				UserId:      taskModel.UserId,
				Title:       taskModel.Title,
				Description: taskModel.Description.String,
				Category:    taskModel.Category,
				Priority:    int32(taskModel.Priority),
				Status:      int32(taskModel.Status),
				DueDate:     "", // 简化处理，实际项目中应该格式化
				CreatedAt:   taskModel.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
				UpdatedAt:   taskModel.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
				CompletedAt: "", // 简化处理
			}
			tasks = append(tasks, taskPB)
		}
	}

	// 应用分页
	total = int32(len(tasks))
	start := (in.Page - 1) * in.Size
	end := start + in.Size
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	if start < end {
		tasks = tasks[start:end]
	} else {
		tasks = []*task.Task{}
	}

	return &task.GetTasksResponse{
		Tasks: tasks,
		Total: total,
	}, nil
}
