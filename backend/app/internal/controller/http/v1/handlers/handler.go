package handlers

import (
	"backend.app/configs"
	"backend.app/database"
	"backend.app/internal/core"
	"backend.app/pkg/logger"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	core   core.Operations
	log    *logger.Logger
	config *configs.Config
}

type Operations interface {
	// Auth
	SignUpUser(c *gin.Context)
	SignUpBusiness(c *gin.Context)
	Login(c *gin.Context)
	ConfirmPhone(c *gin.Context)
	SendResetPasswordToken(c *gin.Context)
	ResetPassword(c *gin.Context)
	Me(c *gin.Context)

	// Middleware
	AuthenticatedUserMiddleware() gin.HandlerFunc

	// Profile
	UpdateUserProfile(c *gin.Context)
	UploadUserProfilePicture(c *gin.Context)

	// Settings
	LogoutMiddleware() gin.HandlerFunc

	// Wallet
	GetUserWallet(c *gin.Context)
}

func NewHandler(log *logger.Logger, config *configs.Config, db *database.DB) Operations {
	newMiddleware, err := middleware.NewMiddleware(db, config, log)
	if err != nil {
		log.Fatal("middleware error: %v", err)
	}

	h := Handler{
		//controller: controllers.New(db, config, newMiddleware, log),
		config: config,
		log:    log,
		core:   core.NewCore(config, log, db, newMiddleware),
	}
	op := Operations(&h)
	return op
}
