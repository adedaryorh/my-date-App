package v1

import (
	"backend.app/internal/controller/http/v1/handlers"
	"github.com/gin-gonic/gin"
)

// CelebrationRoutes stores all celebration routes
func CelebrationRoutes(server *gin.RouterGroup, handler handlers.Operations) {
	celebration := server.Group("/celebrations", handler.AuthenticatedUserMiddleware())
	{
		celebration.POST("", handler.CreateCelebration)
		celebration.GET("", handler.GetCelebrations)
	}
}
