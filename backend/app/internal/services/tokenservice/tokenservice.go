package tokenservice

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"backend.app/common/helpers"
	"backend.app/common/messages"
	"backend.app/internal/models"
	"backend.app/internal/repo"
	"backend.app/internal/services/redisservice"
	"github.com/google/uuid"
)

var defaultTtl = "1h"

type TokenService struct {
	redis *redisservice.Redis
	repo  repo.Operations
}

func NewTokenService(redis *redisservice.Redis, repo repo.Operations) TokenService {
	tokenService := TokenService{
		redis: redis,
		repo:  repo,
	}
	return tokenService
}

func (t *TokenService) SetToken(ctx context.Context, key string, ttl *time.Duration) string {
	if ttl == nil {
		d := helpers.GetDurationFromTimeString(defaultTtl)
		ttl = &d
	}

	code := helpers.GenerateRandomNumber(6)
	// hash code
	codeHash := helpers.HashString(code)

	// send to redis
	err := t.redis.Set(ctx, key, codeHash, *ttl)
	fmt.Println("TOKEN ERR", err)
	return code
}

func (t *TokenService) ValidateToken(ctx context.Context, key, token string) bool {
	return t.redis.GetValue(ctx, key) == helpers.HashString(token)
}

func (t *TokenService) SetStateToken(ctx context.Context, userID string, ttl *time.Duration, isMobileFriendly bool) string {
	if ttl == nil {
		d := helpers.GetDurationFromTimeString(defaultTtl)
		ttl = &d
	}
	code := helpers.GenerateRandomNumber(6)
	if !isMobileFriendly {
		code = helpers.GenerateRandomByte(16)
	}
	// hash code
	codeHash := helpers.HashString(code)

	// send to redis
	t.redis.Set(ctx, fmt.Sprintf("%s:%s:%s", models.RedisKeys.DataAuthStateTokens, userID, codeHash), "", *ttl)

	strToReturn := fmt.Sprintf("%s:%s", userID, code)
	return base64.RawURLEncoding.EncodeToString([]byte(strToReturn))
}

func (t *TokenService) DecodeStateToken(stateToken string) models.DecodedStateToken {
	var decoded models.DecodedStateToken
	byteStr, _ := base64.RawURLEncoding.DecodeString(stateToken)
	splitted := strings.Split(string(byteStr), ":")
	if len(splitted) > 1 {
		decoded.UserUUID = splitted[0]
		decoded.Code = splitted[1]
	}
	return decoded
}

func (t *TokenService) ValidateStateToken(ctx context.Context, decoded models.DecodedStateToken) (*models.User, error) {
	value := t.redis.KeyExists(ctx, fmt.Sprintf("%s:%s:%s", models.RedisKeys.DataAuthStateTokens, decoded.UserUUID, helpers.HashString(decoded.Code)))
	if value < 1 {
		return nil, messages.ErrInvalidToken
	}
	id, err := uuid.Parse(decoded.UserUUID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	user, err := t.repo.GetUserByField(ctx, helpers.Map{"user_id": id})
	if err != nil {
		return nil, err
	}
	return user, nil
}
