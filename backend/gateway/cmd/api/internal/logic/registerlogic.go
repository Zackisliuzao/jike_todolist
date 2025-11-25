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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// User registration
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.AuthResp, err error) {
	// 调用用户RPC服务
	rpcResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	if err != nil {
		l.Errorf("注册失败: %v", err)
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
