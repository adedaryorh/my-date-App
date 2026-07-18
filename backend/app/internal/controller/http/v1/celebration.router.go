package v1

import (
	"time"

	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// CelebrationRoutes stores all celebration routes
func CelebrationRoutes(server *gin.RouterGroup, handler handlers.Operations, mw *middleware.Middleware) {
	celebration := server.Group("/celebrations", handler.AuthenticatedUserMiddleware())
	{
		// Rate limiting for celebration endpoints - higher limits as these are content creation/celebration events
		celebrationLimiter := mw.RateLimiter.RateLimit(middleware.RateLimiterConfig{
			Requests: 50,          // 50 requests
			Window:   time.Minute, // per minute
			KeyFunc:  middleware.KeyFuncs.UserEndpoint,
		})

		celebration.POST("", celebrationLimiter, handler.CreateCelebration)
		celebration.GET("", celebrationLimiter, handler.GetCelebrations)
		celebration.GET("/nearby", celebrationLimiter, handler.GetNearbyCelebrations)
	}
}
