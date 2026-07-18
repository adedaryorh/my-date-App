package v1

import (
	"time"

	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func SettingsRoutes(server *gin.RouterGroup, handler handlers.Operations, mw *middleware.Middleware) {
	settings := server.Group("/settings", handler.AuthenticatedUserMiddleware())
	{
		// Rate limiting for settings endpoints - modify user data, so moderate limits
		settingsLimiter := mw.RateLimiter.RateLimit(middleware.RateLimiterConfig{
			Requests: 20,          // 20 requests
			Window:   time.Minute, // per minute
			KeyFunc:  middleware.KeyFuncs.UserEndpoint,
		})

		settings.PUT("/logout", settingsLimiter, handler.LogoutMiddleware())
		settings.PATCH("/change-password", settingsLimiter, handler.ChangePassword)
		settings.PATCH("/area-of-interests", settingsLimiter, handler.AddAreaOfInterests)
		settings.PATCH("/add-language", settingsLimiter, handler.AddPreferredLanguage)
		settings.PATCH("/add-notification-preferences", settingsLimiter, handler.AddNotificationPreference)
		settings.PATCH("/toggle-push-notification", settingsLimiter, handler.TogglePushNotification)
		settings.PATCH("/toggle-likes-notification", settingsLimiter, handler.ToggleLikesNotification)
		settings.PATCH("/toggle-live-notification", settingsLimiter, handler.ToggleLiveNotification)
		settings.PATCH("/toggle-comment-notification", settingsLimiter, handler.ToggleCommentsNotification)
		settings.PATCH("/toggle-tags-and-mention-notification", settingsLimiter, handler.ToggleTagsAndMentionNotification)
		settings.PATCH("/toggle-repost-notification", settingsLimiter, handler.ToggleRepostNotification)
		settings.PATCH("/toggle-direct-message-notification", settingsLimiter, handler.ToggleDirectMessageNotification)
		settings.PATCH("/toggle-new-follower-notification", settingsLimiter, handler.ToggleNewFollower)
		settings.PATCH("/content-settings", settingsLimiter, handler.SetContentSettings)
		settings.PATCH("/banned-words", settingsLimiter, handler.SetBannedWords)
	}

}
