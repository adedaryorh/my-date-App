package v1

import (
	"time"

	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func WalletRoutes(server *gin.RouterGroup, handler handlers.Operations, mw *middleware.Middleware) {
	wallet := server.Group("/wallets", handler.AuthenticatedUserMiddleware())
	{
		// Rate limiting for wallet endpoints - moderate limits as these are typically accessed infrequently
		walletLimiter := mw.RateLimiter.RateLimit(middleware.RateLimiterConfig{
			Requests: 30,          // 30 requests
			Window:   time.Minute, // per minute
			KeyFunc:  middleware.KeyFuncs.UserEndpoint,
		})

		wallet.GET("", walletLimiter, handler.GetUserWallet)

	}
}
