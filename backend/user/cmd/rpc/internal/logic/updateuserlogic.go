package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"jike_todo/user/cmd/rpc/internal/svc"
	"jike_todo/user/cmd/rpc/user"
	"jike_todo/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserRequest) (*user.User, error) {
	// 查询现有用户信息
	userRecord, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.New("用户不存在")
		}
		l.Errorf("查询用户失败: %v", err)
		return nil, fmt.Errorf("更新用户信息失败")
	}

	// 检查用户状态
	if userRecord.Status != 1 {
		return nil, errors.New("用户账号已被禁用")
	}

	// 更新用户信息（只更新提供的字段）
	if in.Nickname != "" {
		userRecord.Nickname = in.Nickname
	}
	if in.Avatar != "" {
		userRecord.Avatar = in.Avatar
	}

	userRecord.UpdatedAt = time.Now()

	// 保存更新
	err = l.svcCtx.UserModel.Update(l.ctx, userRecord)
	if err != nil {
		l.Errorf("更新用户信息失败: %v", err)
		return nil, fmt.Errorf("更新用户信息失败")
	}

	// 构建响应
	userInfo := &user.User{
		Id:        userRecord.Id,
		Username:  userRecord.Username,
		Email:     userRecord.Email,
		Nickname:  userRecord.Nickname,
		Avatar:    userRecord.Avatar,
		Status:    int32(userRecord.Status),
		CreatedAt: userRecord.CreatedAt.Format(time.RFC3339),
	}

	return userInfo, nil
}
