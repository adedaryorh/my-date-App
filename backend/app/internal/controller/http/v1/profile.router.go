package v1

import (
	"time"

	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// ProfileRoutes stores all profile routes
func ProfileRoutes(server *gin.RouterGroup, handler handlers.Operations, mw *middleware.Middleware) {
	profile := server.Group("/profile", handler.AuthenticatedUserMiddleware())
	{
		// Rate limiting for profile endpoints - moderate limits as these modify user data
		profileLimiter := mw.RateLimiter.RateLimit(middleware.RateLimiterConfig{
			Requests: 30,          // 30 requests
			Window:   time.Minute, // per minute
			KeyFunc:  middleware.KeyFuncs.UserEndpoint,
		})

		profile.PATCH("", profileLimiter, handler.UpdateUserProfile)
		profile.PATCH("/add-profile-image", profileLimiter, handler.UploadUserProfilePicture)
		profile.PATCH("/:id/block", profileLimiter, handler.BlockUser)
		profile.PATCH("/:id/unblock", profileLimiter, handler.UnBlockUser)
		profile.PATCH("/:id/follow", profileLimiter, handler.FollowUser)
		profile.PATCH("/:id/unfollow", profileLimiter, handler.UnFollowUser)
	}
}
