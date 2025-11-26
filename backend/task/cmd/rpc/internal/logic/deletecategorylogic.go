package logic

import (
	"context"
	"errors"

	"jike_todo/task/cmd/rpc/internal/svc"
	"jike_todo/task/cmd/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(in *task.DeleteCategoryRequest) (*task.DeleteCategoryResponse, error) {
	// 先检查分类是否存在且属于该用户
	category, err := l.svcCtx.CateModel.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Errorf("查询分类失败: %v", err)
		return nil, err
	}

	// 验证分类是否属于该用户
	if category.UserId != in.UserId {
		l.Error("用户无权限删除此分类")
		return nil, errors.New("无权限删除此分类")
	}

	// 删除分类
	err = l.svcCtx.CateModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Errorf("删除分类失败: %v", err)
		return nil, err
	}

	l.Infof("用户 %d 成功删除分类 %d", in.UserId, in.Id)

	return &task.DeleteCategoryResponse{
		Success: true,
	}, nil
}
