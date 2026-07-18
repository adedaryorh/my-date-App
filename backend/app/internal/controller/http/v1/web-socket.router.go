package v1

import (
	"time"

	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func WebSocketRoutes(server *gin.RouterGroup, handler handlers.Operations, mw *middleware.Middleware) {
	socket := server.Group("/ws")
	{
		// Rate limiting for websocket connection attempts - prevent abuse of connection attempts
		wsLimiter := mw.RateLimiter.RateLimit(middleware.RateLimiterConfig{
			Requests: 10,                     // 10 connection attempts
			Window:   time.Minute,            // per minute
			KeyFunc:  middleware.KeyFuncs.IP, // Limit by IP for connection attempts
		})

		socket.GET("", wsLimiter, handler.WebSocketHandler)
	}
}
