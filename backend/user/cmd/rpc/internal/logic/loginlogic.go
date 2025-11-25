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

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.AuthResponse, error) {
	fmt.Printf("=== RPC LOGIN START ===\n")
	fmt.Printf("用户名: %s\n", in.Username)
	fmt.Printf("密码: %s (长度: %d)\n", in.Password, len(in.Password))

	// 通过用户名查找用户
	userRecord, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		fmt.Printf("查询用户失败: %v\n", err)
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, fmt.Errorf("登录失败: %v", err)
	}

	fmt.Printf("找到用户: ID=%d, 状态=%d\n", userRecord.Id, userRecord.Status)

	// 检查用户状态
	if userRecord.Status != 1 {
		fmt.Printf("用户状态被禁用: %d\n", userRecord.Status)
		return nil, errors.New("用户账号已被禁用")
	}

	fmt.Printf("开始密码验证，存储哈希长度: %d\n", len(userRecord.PasswordHash))

	// 验证密码
	passwordValid := tool.CheckPasswordHash(in.Password, userRecord.PasswordHash)
	fmt.Printf("密码验证结果: %t\n", passwordValid)

	if !passwordValid {
		return nil, errors.New("用户名或密码错误")
	}

	fmt.Printf("密码验证成功\n")

	// 生成JWT令牌
	accessToken, refreshToken, expiresAt, err := l.generateTokens(userRecord.Id)
	if err != nil {
		l.Errorf("生成令牌失败: %v", err)
		return nil, fmt.Errorf("登录失败")
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

	return &user.AuthResponse{
		User:         userInfo,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}

func (l *LoginLogic) generateTokens(userId int64) (string, string, int64, error) {
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
