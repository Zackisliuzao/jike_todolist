package logic

import (
	"context"
	"fmt"
	"time"

	"jike_todo/task/cmd/rpc/internal/svc"
	"jike_todo/task/cmd/rpc/task"
	"jike_todo/task/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Category operations
func (l *CreateCategoryLogic) CreateCategory(in *task.CreateCategoryRequest) (*task.Category, error) {
	// 设置默认颜色
	if in.Color == "" {
		in.Color = "#1890ff"
	}

	// 创建分类记录
	categoryRecord := &model.Categories{
		UserId:    in.UserId,
		Name:      in.Name,
		Color:     in.Color,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	l.Infof("准备插入分类记录: %+v", categoryRecord)
	result, err := l.svcCtx.CateModel.Insert(l.ctx, categoryRecord)
	if err != nil {
		l.Errorf("插入分类失败: %v", err)
		return nil, fmt.Errorf("创建分类失败: %v", err)
	}
	l.Infof("插入分类成功")

	categoryId, err := result.LastInsertId()
	if err != nil {
		l.Errorf("获取分类ID失败: %v", err)
		return nil, fmt.Errorf("创建分类失败")
	}

	// 查询刚创建的分类
	createdCategory, err := l.svcCtx.CateModel.FindOne(l.ctx, categoryId)
	if err != nil {
		l.Errorf("查询创建的分类失败: %v", err)
		return nil, fmt.Errorf("创建分类失败")
	}

	// 构建响应
	return &task.Category{
		Id:        createdCategory.Id,
		UserId:    createdCategory.UserId,
		Name:      createdCategory.Name,
		Color:     createdCategory.Color,
		CreatedAt: createdCategory.CreatedAt.Format(time.RFC3339),
		UpdatedAt: createdCategory.UpdatedAt.Format(time.RFC3339),
	}, nil
}
