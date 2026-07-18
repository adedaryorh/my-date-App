package v1

import (
	"time"

	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// BlockRoutes stores all block routes
func BlockRoutes(server *gin.RouterGroup, handler handlers.Operations, mw *middleware.Middleware) {
	block := server.Group("/block", handler.AuthenticatedUserMiddleware())
	{
		// Rate limiting for block endpoints - moderate limits as these modify user relationships
		blockLimiter := mw.RateLimiter.RateLimit(middleware.RateLimiterConfig{
			Requests: 30,          // 30 requests
			Window:   time.Minute, // per minute
			KeyFunc:  middleware.KeyFuncs.UserEndpoint,
		})

		block.GET("", blockLimiter, handler.GetAllBlockedUsers)
		block.GET("/:id", blockLimiter, handler.GetBlockedUser)
	}
}
