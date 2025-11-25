package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"jike_todo/common/tool"
	"jike_todo/user/cmd/rpc/internal/svc"
	"jike_todo/user/cmd/rpc/user"
	"jike_todo/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePasswordLogic) ChangePassword(in *user.ChangePasswordRequest) (*user.User, error) {
	// 查询用户信息
	userRecord, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.New("用户不存在")
		}
		l.Errorf("查询用户失败: %v", err)
		return nil, fmt.Errorf("修改密码失败")
	}

	// 检查用户状态
	if userRecord.Status != 1 {
		return nil, errors.New("用户账号已被禁用")
	}

	// 验证旧密码
	if !tool.CheckPasswordHash(in.OldPassword, userRecord.PasswordHash) {
		return nil, errors.New("原密码错误")
	}

	// 生成新密码哈希
	newPasswordHash, err := tool.BcryptByString(in.NewPassword)
	if err != nil {
		l.Errorf("新密码哈希失败: %v", err)
		return nil, fmt.Errorf("修改密码失败")
	}

	// 更新密码
	userRecord.PasswordHash = newPasswordHash
	userRecord.UpdatedAt = time.Now()

	err = l.svcCtx.UserModel.Update(l.ctx, userRecord)
	if err != nil {
		l.Errorf("更新密码失败: %v", err)
		return nil, fmt.Errorf("修改密码失败")
	}

	// 构建响应（不返回敏感信息）
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
