package middleware

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"backend.app/pkg/logger"
)

// RateLimiterConfig holds configuration for rate limiter
type RateLimiterConfig struct {
	Requests int           // max requests allowed
	Window   time.Duration // time window for rate limit
	KeyFunc  func(*gin.Context) string
}

// RateLimiter implements rate limiting using token bucket algorithm with Redis backend
type RateLimiter struct {
	redisClient RedisClient
	logger      *logger.Logger
}

// RedisClient defines the Redis operations needed for rate limiting
type RedisClient interface {
	Incr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, expiration time.Duration) bool
	Get(ctx context.Context, key string) string
}

// NewRateLimiter creates a new rate limiter middleware
func NewRateLimiter(redisClient RedisClient, logger *logger.Logger) *RateLimiter {
	return &RateLimiter{
		redisClient: redisClient,
		logger:      logger,
	}
}

// RateLimit returns a gin.HandlerFunc that implements rate limiting
func (r *RateLimiter) RateLimit(config RateLimiterConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Requests <= 0 {
			c.Next()
			return
		}

		var key string
		if config.KeyFunc != nil {
			key = config.KeyFunc(c)
		} else {
			// Default key: IP address + endpoint
			key = c.ClientIP() + ":" + c.FullPath()
		}

		// Add rate limit prefix to key
		key = "rate_limit:" + key

		ctx := c.Request.Context()

		// Increment request counter
		count, err := r.redisClient.Incr(ctx, key)
		if err != nil {
			r.logger.Error("rate limiter increment failed: %v", err)
			// Fail open - allow request through on Redis error
			c.Next()
			return
		}

		// Set expiration on first request
		if count == 1 {
			expired := r.redisClient.Expire(ctx, key, config.Window)
			if !expired {
				r.logger.Error("failed to set expiration for rate limit key: %s", key)
			}
		}

		// Check if limit exceeded
		if count > int64(config.Requests) {
			// Calculate retry after seconds
			ttl := ttlFromRedis(c, r.redisClient, key, config.Window)
			retryAfter := int64(0)
			if ttl > 0 {
				retryAfter = int64(ttl.Seconds())
			}

			c.Header("X-RateLimit-Limit", strconv.Itoa(config.Requests))
			c.Header("X-RateLimit-Remaining", "0")
			c.Header("X-RateLimit-Reset", strconv.FormatInt(retryAfter, 10))
			c.Header("Retry-After", strconv.FormatInt(retryAfter, 10))

			c.AbortWithStatusJSON(429, gin.H{
				"error":   "rate_limit_exceeded",
				"message": "Too many requests, please try again later.",
			})
			return
		}

		// Set headers for remaining requests
		remaining := config.Requests - int(count)
		c.Header("X-RateLimit-Limit", strconv.Itoa(config.Requests))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(remaining))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(int64(config.Window.Seconds()), 10))

		c.Next()
	}
}

// Helper to get TTL from Redis (simplified - in real implementation would use TTL command)
func ttlFromRedis(c *gin.Context, rc RedisClient, key string, d time.Duration) time.Duration {
	_ = rc.Get(c.Request.Context(), key)
	// This is a simplification - real implementation would use Redis TTL command
	// For now, we'll return the full window as approximation
	return d
}

// KeyFuncs provides common key functions for rate limiting
var KeyFuncs = struct {
	IP           func(*gin.Context) string
	UserID       func(*gin.Context) string
	Endpoint     func(*gin.Context) string
	IPEndpoint   func(*gin.Context) string
	UserEndpoint func(*gin.Context) string
}{
	IP: func(c *gin.Context) string {
		return c.ClientIP()
	},
	UserID: func(c *gin.Context) string {
		if user, exists := c.Get("authUser"); exists {
			if u, ok := user.(interface{ ID() string }); ok {
				return u.ID()
			}
		}
		return "anonymous"
	},
	Endpoint: func(c *gin.Context) string {
		return c.FullPath()
	},
	IPEndpoint: func(c *gin.Context) string {
		return c.ClientIP() + ":" + c.FullPath()
	},
	UserEndpoint: func(c *gin.Context) string {
		userID := "anonymous"
		if user, exists := c.Get("authUser"); exists {
			if u, ok := user.(interface{ ID() string }); ok {
				userID = u.ID()
			}
		}
		return userID + ":" + c.FullPath()
	},
}
