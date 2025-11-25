// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"jike_todo/gateway/cmd/api/internal/svc"
	"jike_todo/gateway/cmd/api/internal/types"
	"jike_todo/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// User login
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.AuthResp, err error) {
	l.Infof("网关收到登录请求，用户名: %s", req.Username)

	// 调用用户RPC服务
	rpcResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		l.Errorf("登录失败: %v", err)
		return nil, err
	}

	// 手动复制响应数据
	resp = &types.AuthResp{}
	resp.User = types.UserInfo{
		ID:        rpcResp.User.Id,
		Username:  rpcResp.User.Username,
		Email:     rpcResp.User.Email,
		Nickname:  rpcResp.User.Nickname,
		Avatar:    rpcResp.User.Avatar,
		Status:    rpcResp.User.Status,
		CreatedAt: rpcResp.User.CreatedAt,
	}
	resp.AccessToken = rpcResp.AccessToken
	resp.RefreshToken = rpcResp.RefreshToken
	resp.ExpiresAt = rpcResp.ExpiresAt

	return resp, nil
}
