package v1

import (
	"time"

	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// FollowerRoutes stores all follower routes
func FollowerRoutes(server *gin.RouterGroup, handler handlers.Operations, mw *middleware.Middleware) {
	follower := server.Group("/followers", handler.AuthenticatedUserMiddleware())
	{
		// Rate limiting for follower endpoints - moderate limits as these are frequently accessed
		followerLimiter := mw.rateLimiter.RateLimit(middleware.RateLimiterConfig{
			Requests: 50,    // 50 requests
			Window:   time.Minute, // per minute
			KeyFunc:  middleware.KeyFuncs.UserEndpoint,
		})

		follower.GET("", followerLimiter, handler.GetAllFollowers)
		follower.GET("/:id", followerLimiter, handler.GetSingleFollower)
		follower.GET("/friends", followerLimiter, handler.GetFriends)

	}
}
