package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"jike_todo/user/cmd/rpc/internal/svc"
	"jike_todo/user/cmd/rpc/user"
	"jike_todo/user/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokenLogic) RefreshToken(in *user.RefreshTokenRequest) (*user.AuthResponse, error) {
	// 验证刷新令牌
	token, err := jwt.Parse(in.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(l.svcCtx.Config.JwtAuth.AccessSecret), nil
	})

	if err != nil {
		l.Errorf("解析刷新令牌失败: %v", err)
		return nil, errors.New("无效的刷新令牌")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("无效的刷新令牌")
	}

	// 检查令牌类型
	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "refresh" {
		return nil, errors.New("无效的令牌类型")
	}

	// 获取用户ID
	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return nil, errors.New("无效的令牌")
	}
	userId := int64(userIdFloat)

	// 验证用户是否存在且状态正常
	userRecord, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.New("用户不存在")
		}
		l.Errorf("查询用户失败: %v", err)
		return nil, fmt.Errorf("刷新令牌失败")
	}

	if userRecord.Status != 1 {
		return nil, errors.New("用户账号已被禁用")
	}

	// 生成新的JWT令牌
	accessToken, refreshToken, expiresAt, err := l.generateTokens(userId)
	if err != nil {
		l.Errorf("生成新令牌失败: %v", err)
		return nil, fmt.Errorf("刷新令牌失败")
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

func (l *RefreshTokenLogic) generateTokens(userId int64) (string, string, int64, error) {
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
