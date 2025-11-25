// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"jike_todo/common/ctxdata"
	"jike_todo/gateway/cmd/api/internal/svc"
	"jike_todo/gateway/cmd/api/internal/types"
	"jike_todo/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Logout
func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() (resp *types.BaseResponse, err error) {
	// 从 JWT 上下文获取用户ID
	userId := ctxdata.GetUidFromCtx(l.ctx)
	if userId == 0 {
		l.Error("无法获取用户ID")
		return &types.BaseResponse{
			Code:    401,
			Message: "用户未认证",
		}, nil
	}

	// 调用用户RPC服务处理登出
	_, err = l.svcCtx.UserRpc.Logout(l.ctx, &user.LogoutRequest{
		UserId: userId,
	})
	if err != nil {
		l.Errorf("登出失败: %v", err)
		// 登出失败不影响客户端清除token
	}

	return &types.BaseResponse{
		Code:    0,
		Message: "登出成功",
		Data:    nil,
	}, nil
}
