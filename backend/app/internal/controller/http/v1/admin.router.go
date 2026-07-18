package v1

import (
	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// AdminRoutes registers admin-related routes
func AdminRoutes(v *gin.RouterGroup, handler handlers.Operations, _ *middleware.Middleware) {
	admin := v.Group("/admin", handler.AuthenticatedUserMiddleware())
	admin.GET("/users", handler.GetAllUsers)
	admin.PUT("/users/:user_id/role", handler.UpdateUserRole)
	admin.DELETE("/users/:user_id", handler.DeleteUser)
	admin.GET("/moderation-queue", handler.GetModerationQueue)
	admin.POST("/moderation-queue/:id/resolve", handler.ResolveModeration)
}
