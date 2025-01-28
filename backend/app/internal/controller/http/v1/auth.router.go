package v1

import (
	"backend.app/internal/controller/http/v1/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(server *gin.RouterGroup, handler handlers.Operations) {
	auth := server.Group("/auth")
	{
		auth.POST("/sign-up")
	}
}
