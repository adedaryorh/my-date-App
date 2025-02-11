package v1

import (
	"github.com/gin-gonic/gin"

	"backend.app/internal/controller/http/v1/handlers"
)

func WalletRoutes(server *gin.RouterGroup, handler handlers.Operations) {
	wallet := server.Group("/wallets", handler.AuthenticatedUserMiddleware())
	{
		wallet.GET("", handler.GetUserWallet)

	}
}
