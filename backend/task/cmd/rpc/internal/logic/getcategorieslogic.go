package logic

import (
	"context"
	"time"

	"jike_todo/task/cmd/rpc/internal/svc"
	"jike_todo/task/cmd/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoriesLogic {
	return &GetCategoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoriesLogic) GetCategories(in *task.GetCategoriesRequest) (*task.GetCategoriesResponse, error) {
	// 简化实现：返回默认分类和示例数据
	// 实际项目中需要从数据库查询用户的分类

	categories := []*task.Category{
		{
			Id:        1,
			UserId:    in.UserId,
			Name:      "工作",
			Color:     "#1890ff",
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
		{
			Id:        2,
			UserId:    in.UserId,
			Name:      "学习",
			Color:     "#52c41a",
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
		{
			Id:        3,
			UserId:    in.UserId,
			Name:      "生活",
			Color:     "#faad14",
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
	}

	return &task.GetCategoriesResponse{
		Categories: categories,
	}, nil
}
