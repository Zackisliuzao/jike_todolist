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

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Change password
func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) (resp *types.BaseResponse, err error) {
	// 从JWT上下文中获取用户ID
	userId := ctxdata.GetUidFromCtx(l.ctx)

	// 调用RPC服务修改密码
	rpcResp, err := l.svcCtx.UserRpc.ChangePassword(l.ctx, &user.ChangePasswordRequest{
		UserId:      userId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		l.Errorf("调用RPC服务失败: %v", err)
		resp = &types.BaseResponse{
			Code:    500,
			Message: "修改密码失败: " + err.Error(),
			Data:    nil,
		}
		return
	}

	// 构建用户信息响应
	userInfo := &types.UserInfo{
		ID:          rpcResp.Id,
		Username:    rpcResp.Username,
		Email:       rpcResp.Email,
		Nickname:    rpcResp.Nickname,
		Avatar:      rpcResp.Avatar,
		Status:      rpcResp.Status,
		CreatedAt:   rpcResp.CreatedAt,
	}

	resp = &types.BaseResponse{
		Code:    200,
		Message: "密码修改成功",
		Data:    userInfo,
	}
	return
}
