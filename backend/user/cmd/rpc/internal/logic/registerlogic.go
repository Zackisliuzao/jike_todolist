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

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.AuthResponse, error) {
	// 检查用户名是否已存在
	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err == nil {
		return nil, errors.New("用户名已存在")
	}
	if !errors.Is(err, model.ErrNotFound) {
		l.Errorf("检查用户名失败: %v", err)
		return nil, fmt.Errorf("检查用户名失败")
	}

	// 检查邮箱是否已存在
	_, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err == nil {
		return nil, errors.New("邮箱已被注册")
	}
	if !errors.Is(err, model.ErrNotFound) {
		l.Errorf("检查邮箱失败: %v", err)
		return nil, fmt.Errorf("检查邮箱失败")
	}

	// 密码哈希
	passwordHash, err := tool.BcryptByString(in.Password)
	if err != nil {
		l.Errorf("密码哈希失败: %v", err)
		return nil, fmt.Errorf("密码处理失败")
	}

	// 创建用户记录
	userRecord := &model.Users{
		Username:     in.Username,
		Email:        in.Email,
		PasswordHash: passwordHash,
		Nickname:     in.Nickname,
		Avatar:       "",
		Status:       1, // 正常状态
	}

	result, err := l.svcCtx.UserModel.Insert(l.ctx, userRecord)
	if err != nil {
		l.Errorf("插入用户失败: %v", err)
		return nil, fmt.Errorf("注册失败")
	}

	userId, err := result.LastInsertId()
	if err != nil {
		l.Errorf("获取用户ID失败: %v", err)
		return nil, fmt.Errorf("注册失败")
	}

	userRecord.Id = userId

	// 生成JWT令牌
	accessToken, refreshToken, expiresAt, err := l.generateTokens(userId)
	if err != nil {
		l.Errorf("生成令牌失败: %v", err)
		return nil, fmt.Errorf("登录失败")
	}

	// 构建响应
	userInfo := &user.User{
		Id:        userId,
		Username:  userRecord.Username,
		Email:     userRecord.Email,
		Nickname:  userRecord.Nickname,
		Avatar:    userRecord.Avatar,
		Status:    int32(userRecord.Status),
		CreatedAt: userRecord.CreatedAt.Format(time.RFC3339),
	}

	return &user.AuthResponse{
		User:         userInfo,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}

func (l *RegisterLogic) generateTokens(userId int64) (string, string, int64, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire) * time.Second)

	// 生成访问令牌
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    expiresAt.Unix(),
		"iat":    now.Unix(),
		"iss":    l.svcCtx.Config.JwtAuth.Issuer,
	})

	accessTokenString, err := accessToken.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))
	if err != nil {
		return "", "", 0, err
	}

	// 生成刷新令牌 (有效期更长)
	refreshExpiresAt := now.Add(time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire) * time.Second * 7) // 7倍有效期
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    refreshExpiresAt.Unix(),
		"iat":    now.Unix(),
		"iss":    l.svcCtx.Config.JwtAuth.Issuer,
		"type":   "refresh",
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))
	if err != nil {
		return "", "", 0, err
	}

	return accessTokenString, refreshTokenString, expiresAt.Unix(), nil
}
