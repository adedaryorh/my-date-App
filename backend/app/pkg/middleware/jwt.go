package middleware

import (
	"context"
	"fmt"
	"time"

	"backend.app/common/helpers"
	"backend.app/common/messages"
	"backend.app/configs"
	"backend.app/internal/models"
	"backend.app/internal/services/redisservice"

	"github.com/golang-jwt/jwt/v5"
)

type (
	Payload struct {
		UserID string `json:"sub"`
		Email  string `json:"email"`
		jwt.RegisteredClaims
	}

	JwtMaker struct {
		secretKey []byte
		config    *configs.Config
		redis     *redisservice.Redis
	}
)

func NewPayload(user models.User, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		user.ID.String(),
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        helpers.GenerateRandomUppercase(26),
		},
	}

	return payload, nil
}

func (p Payload) Valid() error {
	if time.Now().After(p.RegisteredClaims.ExpiresAt.Time) {
		return fmt.Errorf("token is expired")
	}
	return nil
}

func NewJwtMaker(config *configs.Config, redis *redisservice.Redis) (TokenMaker, error) {

	return &JwtMaker{
		secretKey: []byte(config.JwtSecret),
		config:    config,
		redis:     redis,
	}, nil
}

func (j JwtMaker) createToken(ctx context.Context, user models.User, duration time.Duration, redisKey string) (string, error) {
	payload, _ := NewPayload(user, duration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}

	j.redis.Set(ctx, fmt.Sprintf("%s:%s:%s", redisKey, payload.UserID, payload.RegisteredClaims.ID), "", duration)
	return tokenString, nil
}

func (j JwtMaker) createAccessToken(ctx context.Context, user models.User) (string, error) {
	duration, _ := time.ParseDuration(j.config.JwtSecretExpiry)
	return j.createToken(ctx, user, duration, models.RedisKeys.AccessToken)
}

func (j JwtMaker) createRefreshToken(ctx context.Context, user models.User) (string, error) {
	duration, _ := time.ParseDuration("720h") // 1 month
	return j.createToken(ctx, user, duration, models.RedisKeys.RefreshToken)
}

func (j JwtMaker) CreateAuthRefreshTokens(ctx context.Context, user models.User) (*models.AuthTokens, error) {
	accessToken, err := j.createAccessToken(ctx, user)
	if err != nil {
		return nil, err
	}
	refreshToken, err := j.createRefreshToken(ctx, user)
	if err != nil {
		return nil, err
	}

	tokens := models.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &tokens, nil
}

func (j JwtMaker) VerifyToken(tokenString string) (*Payload, error) {
	payload := Payload{}

	tkn, err := jwt.ParseWithClaims(tokenString, &payload, func(token *jwt.Token) (any, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, messages.ErrInvalidToken
	}
	return &payload, nil
}
