package middleware

import (
	"errors"

	"context"
	"fmt"
	"strings"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/common/messages"
	"backend.app/configs"
	"backend.app/database"
	"backend.app/internal/models"
	"backend.app/internal/repo"
	"backend.app/internal/services/redisservice"
	"backend.app/internal/services/tokenservice"
	"backend.app/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	authorizationHeader = "Authorization"
	authorizationBearer = "Bearer"
)

type TokenMaker interface {
	CreateAuthRefreshTokens(ctx context.Context, user models.User) (*models.AuthTokens, error)
	VerifyToken(token string) (*Payload, error)
}

type Middleware struct {
	Jwt          TokenMaker
	logger       *logger.Logger
	db           *database.DB
	repo         repo.Operations
	config       *configs.Config
	tokenService tokenservice.TokenService
	redis        redisservice.Redis
}

func NewMiddleware(db *database.DB, config *configs.Config, log *logger.Logger) (*Middleware, error) {

	redis := redisservice.Redis{Client: db.Redis.Client}
	jwt, err := NewJwtMaker(config, &redis)
	if err != nil {
		return nil, err
	}
	repo := repo.NewRepo(db)
	tokenService := tokenservice.NewTokenService(&redis, repo)
	m := &Middleware{
		Jwt:          jwt,
		logger:       log,
		config:       config,
		db:           db,
		repo:         repo,
		tokenService: tokenService,
		redis:        redis,
	}

	return m, nil
}

// JwtUserAuth hybrid middleware returns an authorized user
func (m *Middleware) JwtUserAuth(c *gin.Context) (*models.User, error) {
	verified, err := m.fetchPayloadFromContext(c)
	if err != nil {
		return nil, err
	}
	return m.getUserFromToken(c, verified, models.RedisKeys.AccessToken)
}

func (m *Middleware) JwtLogUserOut(c *gin.Context) error {
	verified, err := m.fetchPayloadFromContext(c)
	if err != nil {
		return err
	}

	// remove auth from redis
	m.redis.Delete(c, fmt.Sprintf("%s:%s:%s", models.RedisKeys.AccessToken, verified.UserID, verified.ID))
	m.redis.Delete(c, fmt.Sprintf("%s:%s:%s", models.RedisKeys.RefreshToken, verified.UserID, verified.ID))

	return nil
}

func (m *Middleware) fetchPayloadFromContext(c *gin.Context) (*Payload, error) {
	authorization := c.GetHeader(authorizationHeader)
	if len(authorization) < 1 {
		return nil, messages.ErrInvalidToken
	}

	fields := strings.Fields(authorization)
	if len(fields) != 2 {
		return nil, messages.ErrInvalidToken
	}

	verified, err := m.Jwt.VerifyToken(fields[1])
	if err != nil {
		return nil, err
	}

	return verified, nil
}

func (m *Middleware) JwtRefreshTokenAuth(c *gin.Context, token string, redisKey string) (*models.User, error) {
	verified, err := m.Jwt.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	return m.getUserFromToken(c, verified, redisKey)
}

func (m *Middleware) getUserFromToken(ctx context.Context, verified *Payload, redisKey string) (*models.User, error) {
	id, err := uuid.Parse(verified.UserID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	user, err := m.repo.GetUserByField(ctx, helpers.Map{"user_id": id})
	if err != nil {
		return nil, err
	}

	if user.Status != string(constants.UserStatusActive) {
		return nil, messages.ErrInactiveUser
	}
	// ensure token is valid on redis too
	value := m.redis.KeyExists(ctx, fmt.Sprintf("%s:%s:%s", redisKey, verified.UserID, verified.ID))
	if value < 1 {
		return nil, messages.ErrNoActiveSession
	}
	return user, nil
}

func (m *Middleware) StateTokenAuth(c *gin.Context) (*models.User, error) {
	stateToken := c.Query("stateToken")
	if len(stateToken) < 1 {
		return nil, messages.ErrInvalidToken
	}
	decoded := m.tokenService.DecodeStateToken(stateToken)
	user, err := m.tokenService.ValidateStateToken(c, decoded)
	if err != nil {
		return nil, err
	}
	return user, nil
}
