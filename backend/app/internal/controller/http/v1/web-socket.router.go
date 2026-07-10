package v1

import (
	"backend.app/internal/controller/http/v1/handlers"
	"github.com/gin-gonic/gin"
)

func WebSocketRoutes(server *gin.RouterGroup, handler handlers.Operations) {
	socket := server.Group("/ws")
	{
		socket.GET("", handler.WebSocketHandler)
	}
}
