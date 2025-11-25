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

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.User, error) {
	// 查询用户信息
	userRecord, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.New("用户不存在")
		}
		l.Errorf("查询用户失败: %v", err)
		return nil, fmt.Errorf("获取用户信息失败")
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
