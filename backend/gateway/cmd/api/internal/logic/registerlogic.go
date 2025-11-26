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

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.BaseResponse, err error) {
	// 调用用户RPC服务
	rpcResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	if err != nil {
		l.Errorf("注册失败: %v", err)
		return &types.BaseResponse{
			Code:    500,
			Message: "注册失败: " + err.Error(),
		}, err
	}

	// 构建认证响应数据
	authData := &types.AuthResp{}
	authData.User = types.UserInfo{
		ID:        rpcResp.User.Id,
		Username:  rpcResp.User.Username,
		Email:     rpcResp.User.Email,
		Nickname:  rpcResp.User.Nickname,
		Avatar:    rpcResp.User.Avatar,
		Status:    rpcResp.User.Status,
		CreatedAt: rpcResp.User.CreatedAt,
	}
	authData.AccessToken = rpcResp.AccessToken
	authData.RefreshToken = rpcResp.RefreshToken
	authData.ExpiresAt = rpcResp.ExpiresAt

	// 返回统一格式的响应
	return &types.BaseResponse{
		Code:    200,
		Message: "注册成功",
		Data:    authData,
	}, nil
}
