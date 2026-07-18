package v1

import (
	"time"

	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// AIRoutes stores all AI routes
func AIRoutes(server *gin.RouterGroup, handler handlers.Operations, mw *middleware.Middleware) {
	ai := server.Group("/ai", handler.AuthenticatedUserMiddleware())
	{
		// Rate limiting for AI endpoints - moderate limits as these may be used frequently
		aiLimiter := mw.RateLimiter.RateLimit(middleware.RateLimiterConfig{
			Requests: 50,          // 50 requests
			Window:   time.Minute, // per minute
			KeyFunc:  middleware.KeyFuncs.UserEndpoint,
		})

		// Health check
		ai.GET("/health", aiLimiter, handler.HealthCheck)

		// User recommendations
		ai.POST("/users/recommendations", aiLimiter, handler.GetUserRecommendations)

		// Content moderation
		ai.POST("/celebrations/moderate", aiLimiter, handler.ModerateCelebration)
		ai.POST("/celebrations/index", aiLimiter, handler.IndexCelebration)
		ai.POST("/celebrations/search", aiLimiter, handler.SearchCelebrations)

		// Interaction logging
		ai.POST("/interactions", aiLimiter, handler.LogInteraction)
	}
}
