package v1

import (
	"backend.app/internal/controller/http/v1/handlers"
	"github.com/gin-gonic/gin"
)

func SettingsRoutes(server *gin.RouterGroup, handler handlers.Operations) {
	settings := server.Group("/settings", handler.AuthenticatedUserMiddleware())
	{
		settings.PUT("/logout", handler.LogoutMiddleware())
		settings.PATCH("/change-password", handler.ChangePassword)
		settings.PATCH("/area-of-interests", handler.AddAreaOfInterests)
		settings.PATCH("/add-language", handler.AddPreferredLanguage)
		settings.PATCH("/add-notification-preferences", handler.AddNotificationPreference)
		settings.PATCH("/toggle-push-notification", handler.TogglePushNotification)
		settings.PATCH("/toggle-likes-notification", handler.ToggleLikesNotification)
		settings.PATCH("/toggle-live-notification", handler.ToggleLiveNotification)
		settings.PATCH("/toggle-comment-notification", handler.ToggleCommentsNotification)
		settings.PATCH("/toggle-tags-and-mention-notification", handler.ToggleTagsAndMentionNotification)
		settings.PATCH("/toggle-repost-notification", handler.ToggleRepostNotification)
		settings.PATCH("/toggle-direct-message-notification", handler.ToggleDirectMessageNotification)
		settings.PATCH("/toggle-new-follower-notification", handler.ToggleNewFollower)
		settings.PATCH("/content-settings", handler.SetContentSettings)
		settings.PATCH("/banned-words", handler.SetBannedWords)
	}

}
