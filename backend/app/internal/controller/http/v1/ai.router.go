package v1

import (
	"backend.app/internal/controller/http/v1/handlers"
	"github.com/gin-gonic/gin"
)

// AIRoutes stores all AI routes
func AIRoutes(server *gin.RouterGroup, handler handlers.Operations) {
	ai := server.Group("/ai", handler.AuthenticatedUserMiddleware())
	{
		ai.GET("/health", handler.HealthCheck)

		// User recommendations
		ai.POST("/users/recommendations", handler.GetUserRecommendations)

		// Content moderation
		ai.POST("/celebrations/moderate", handler.ModerateCelebration)
		ai.POST("/celebrations/index", handler.IndexCelebration)
		ai.POST("/celebrations/search", handler.SearchCelebrations)

		// Interaction logging
		ai.POST("/interactions", handler.LogInteraction)
	}
}