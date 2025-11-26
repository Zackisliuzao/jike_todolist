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
	// 从数据库查询用户的分类
	categoryRecords, err := l.svcCtx.CateModel.FindByUserId(l.ctx, in.UserId)
	if err != nil {
		l.Errorf("查询用户分类失败: %v", err)
		return nil, err
	}

	// 转换为RPC响应格式
	categories := make([]*task.Category, 0, len(categoryRecords))
	for _, record := range categoryRecords {
		categories = append(categories, &task.Category{
			Id:        record.Id,
			UserId:    record.UserId,
			Name:      record.Name,
			Color:     record.Color,
			CreatedAt: record.CreatedAt.Format(time.RFC3339),
			UpdatedAt: record.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &task.GetCategoriesResponse{
		Categories: categories,
	}, nil
}
