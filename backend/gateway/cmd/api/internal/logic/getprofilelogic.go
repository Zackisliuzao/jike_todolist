// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"jike_todo/common/ctxdata"
	"jike_todo/gateway/cmd/api/internal/svc"
	"jike_todo/gateway/cmd/api/internal/types"
	"jike_todo/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get user profile
func NewGetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProfileLogic {
	return &GetProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProfileLogic) GetProfile() (resp *types.BaseResponse, err error) {
	// 从 JWT 上下文获取用户ID
	userId := ctxdata.GetUidFromCtx(l.ctx)
	if userId == 0 {
		l.Error("无法获取用户ID")
		return nil, errors.New("用户未认证")
	}

	// 调用用户RPC服务获取用户信息
	rpcResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserRequest{
		UserId: userId,
	})
	if err != nil {
		l.Errorf("获取用户信息失败: %v", err)
		return nil, err
	}

	// 构建用户信息响应
	userInfo := types.UserInfo{
		ID:        rpcResp.Id,
		Username:  rpcResp.Username,
		Email:     rpcResp.Email,
		Nickname:  rpcResp.Nickname,
		Avatar:    rpcResp.Avatar,
		Status:    rpcResp.Status,
		CreatedAt: rpcResp.CreatedAt,
	}

	return &types.BaseResponse{
		Code:    0,
		Message: "获取用户信息成功",
		Data:    userInfo,
	}, nil
}
