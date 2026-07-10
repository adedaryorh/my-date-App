package v1

import (
	"backend.app/internal/controller/http/v1/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(server *gin.RouterGroup, handler handlers.Operations) {
	auth := server.Group("/auth")
	{
		auth.POST("/sign-up/user", handler.SignUpUser)
		auth.POST("/sign-up/business", handler.SignUpBusiness)
		auth.POST("/login", handler.Login)
		auth.PATCH("/confirm-phone", handler.ConfirmPhone)
		auth.POST("/reset-password", handler.SendResetPasswordToken)
		auth.PATCH("/reset-password", handler.ResetPassword)
		auth.GET("/self", handler.AuthenticatedUserMiddleware(), handler.Me)
	}
}
