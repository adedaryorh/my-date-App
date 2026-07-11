package v1

import (
	"github.com/gin-gonic/gin"
	"backend.app/internal/controller/http/v1/handlers"
)

// AdminRoutes registers admin-related routes
func AdminRoutes(v *gin.RouterGroup, handler handlers.Operations) {
	admin := v.Group("/admin")
	// Optional: add middleware for admin auth, e.g., adminAuthMiddleware()
	admin.GET("/users", handler.GetAllUsers)
	admin.PUT("/users/:user_id/role", handler.UpdateUserRole)
	admin.DELETE("/users/:user_id", handler.DeleteUser)
}