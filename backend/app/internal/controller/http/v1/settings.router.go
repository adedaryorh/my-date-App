package v1

import (
	"backend.app/internal/controller/http/v1/handlers"
	"github.com/gin-gonic/gin"
)

func SettingsRoutes(server *gin.RouterGroup, handler handlers.Operations) {
	settings := server.Group("/settings", handler.AuthenticatedUserMiddleware())
	{
		settings.PUT("/logout", handler.LogoutMiddleware())
	}
}
