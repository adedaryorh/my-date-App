package v1

import (
	"time"

	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(server *gin.RouterGroup, handler handlers.Operations, mw *middleware.Middleware) {
	auth := server.Group("/auth")
	{
		// Rate limiting for auth endpoints - stricter limits
		authLimiter := mw.RateLimiter.RateLimit(middleware.RateLimiterConfig{
			Requests: 5,           // 5 requests
			Window:   time.Minute, // per minute
			KeyFunc:  middleware.KeyFuncs.UserEndpoint,
		})

		auth.POST("/sign-up/user", authLimiter, handler.SignUpUser)
		auth.POST("/sign-up/business", authLimiter, handler.SignUpBusiness)
		auth.POST("/login", authLimiter, handler.Login)
		auth.PATCH("/confirm-phone", authLimiter, handler.ConfirmPhone)
		auth.PATCH("/confirm-email", authLimiter, handler.ConfirmEmail)
		auth.POST("/reset-password", authLimiter, handler.SendResetPasswordToken)
		auth.PATCH("/reset-password", authLimiter, handler.ResetPassword)
		auth.GET("/self", handler.AuthenticatedUserMiddleware(), handler.Me)
		// Google OAuth - slightly higher limit as it's less frequent but still needs protection
		authGETLimiter := mw.RateLimiter.RateLimit(middleware.RateLimiterConfig{
			Requests: 10,          // 10 requests
			Window:   time.Minute, // per minute
			KeyFunc:  middleware.KeyFuncs.UserEndpoint,
		})
		auth.GET("/google/login", authGETLimiter, handler.GoogleLogin)
		auth.GET("/google/callback", authGETLimiter, handler.GoogleCallback)
	}
}
