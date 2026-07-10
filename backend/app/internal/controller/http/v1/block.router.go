package v1

import (
	"github.com/gin-gonic/gin"

	"backend.app/internal/controller/http/v1/handlers"
)

// BlockRoutes stores all block routes
func BlockRoutes(server *gin.RouterGroup, handler handlers.Operations) {
	block := server.Group("/block", handler.AuthenticatedUserMiddleware())
	{
		block.GET("", handler.GetAllBlockedUsers)
		block.GET("/:id", handler.GetBlockedUser)
	}
}
