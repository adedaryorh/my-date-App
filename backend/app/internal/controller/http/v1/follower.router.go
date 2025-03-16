package v1

import (
	"github.com/gin-gonic/gin"

	"backend.app/internal/controller/http/v1/handlers"
)

// FollowerRoutes stores all follower routes
func FollowerRoutes(server *gin.RouterGroup, handler handlers.Operations) {
	follower := server.Group("/followers", handler.AuthenticatedUserMiddleware())
	{
		follower.GET("", handler.GetAllFollowers)
		follower.GET("/:id", handler.GetSingleFollower)
		follower.GET("/friends", handler.GetFriends)

	}
}
