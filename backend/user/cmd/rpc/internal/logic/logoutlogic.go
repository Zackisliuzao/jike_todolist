package logic

import (
	"context"
	"errors"

	"jike_todo/user/cmd/rpc/internal/svc"
	"jike_todo/user/cmd/rpc/user"
	"jike_todo/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *user.LogoutRequest) (*user.LogoutResponse, error) {
	// 验证用户是否存在
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Infof("登出用户不存在: %d", in.UserId)
		} else {
			l.Errorf("查询用户失败: %v", err)
		}
		// 即使查询失败也返回成功，因为登出操作不需要严格验证
	}

	// 在当前实现中，登出主要是客户端清除token
	// 服务端可以在这里记录日志、清理缓存等操作
	l.Infof("用户 %d 登出", in.UserId)

	return &user.LogoutResponse{
		Success: true,
	}, nil
}
