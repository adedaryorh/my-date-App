package v1

import (
	"time"

	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// NotificationRoutes stores all notification routes
func NotificationRoutes(server *gin.RouterGroup, handler handlers.Operations, mw *middleware.Middleware) {
	notification := server.Group("/notifications", handler.AuthenticatedUserMiddleware())
	{
		// Rate limiting for notification endpoints - moderate limits
		notificationLimiter := mw.RateLimiter.RateLimit(middleware.RateLimiterConfig{
			Requests: 30,          // 30 requests
			Window:   time.Minute, // per minute
			KeyFunc:  middleware.KeyFuncs.UserEndpoint,
		})

		// Get all notifications with pagination
		notification.GET("", notificationLimiter, handler.GetAllNotifications)
		// Get single notification by ID
		notification.GET("/:id", notificationLimiter, handler.GetNotificationById)
		// Mark notification as read
		notification.PATCH("/:id/read", notificationLimiter, handler.MarkNotificationAsRead)
		// Delete notification
		notification.DELETE("/:id", notificationLimiter, handler.DeleteNotification)
	}
}
