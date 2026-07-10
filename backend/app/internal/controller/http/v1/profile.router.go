package v1

import (
	"github.com/gin-gonic/gin"

	"backend.app/internal/controller/http/v1/handlers"
)

// ProfileRoutes stores all profile routes
func ProfileRoutes(server *gin.RouterGroup, handler handlers.Operations) {
	profile := server.Group("/profile", handler.AuthenticatedUserMiddleware())
	{
		profile.PATCH("", handler.UpdateUserProfile)
		profile.PATCH("/add-profile-image", handler.UploadUserProfilePicture)
		profile.PATCH("/:id/block", handler.BlockUser)
		profile.PATCH("/:id/unblock", handler.UnBlockUser)
		profile.PATCH("/:id/follow", handler.FollowUser)
		profile.PATCH("/:id/unfollow", handler.UnFollowUser)
	}
}
