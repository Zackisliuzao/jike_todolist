package logic

import (
	"context"
	"errors"
	"time"

	"jike_todo/task/cmd/rpc/internal/svc"
	"jike_todo/task/cmd/rpc/task"
	"jike_todo/task/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *task.UpdateCategoryRequest) (*task.Category, error) {
	// 先检查分类是否存在且属于该用户
	category, err := l.svcCtx.CateModel.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Errorf("查询分类失败: %v", err)
		return nil, err
	}

	// 验证分类是否属于该用户
	if category.UserId != in.UserId {
		l.Error("用户无权限更新此分类")
		return nil, errors.New("无权限更新此分类")
	}

	// 更新分类信息
	updatedCategory := &model.Categories{
		Id:        in.Id,
		UserId:    in.UserId,
		Name:      in.Name,
		Color:     in.Color,
		CreatedAt: category.CreatedAt,
		UpdatedAt: time.Now(),
	}

	err = l.svcCtx.CateModel.Update(l.ctx, updatedCategory)
	if err != nil {
		l.Errorf("更新分类失败: %v", err)
		return nil, err
	}

	// 查询更新后的分类
	category, err = l.svcCtx.CateModel.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Errorf("查询更新后的分类失败: %v", err)
		return nil, err
	}

	l.Infof("用户 %d 成功更新分类 %d", in.UserId, in.Id)

	return &task.Category{
		Id:        category.Id,
		UserId:    category.UserId,
		Name:      category.Name,
		Color:     category.Color,
		CreatedAt: category.CreatedAt.Format(time.RFC3339),
		UpdatedAt: category.UpdatedAt.Format(time.RFC3339),
	}, nil
}
