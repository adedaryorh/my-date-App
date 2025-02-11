package v1

import (
	"github.com/gin-gonic/gin"

	"backend.app/internal/controller/http/v1/handlers"
)

func UserRoutes(server *gin.RouterGroup, handler handlers.Operations) {
	user := server.Group("/users", handler.AuthenticatedUserMiddleware())
	{
		user.PUT("/update-profile", handler.UpdateUserProfile)
		user.PUT("/add-profile-image", handler.UploadUserProfilePicture)
	}
}
