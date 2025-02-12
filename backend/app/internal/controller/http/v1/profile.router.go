package v1

import (
	"github.com/gin-gonic/gin"

	"backend.app/internal/controller/http/v1/handlers"
)

func ProfileRoutes(server *gin.RouterGroup, handler handlers.Operations) {
	profile := server.Group("/profile", handler.AuthenticatedUserMiddleware())
	{
		profile.PUT("", handler.UpdateUserProfile)
		profile.PUT("/add-profile-image", handler.UploadUserProfilePicture)
	}
}
