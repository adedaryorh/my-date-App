package handlers

import (
	"strconv"

	"backend.app/configs"
	"backend.app/database"
	"backend.app/internal/core"
	"backend.app/internal/dtos"
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

	// block
	GetAllBlockedUsers(c *gin.Context)
	GetBlockedUser(c *gin.Context)

	// celebration
	CreateCelebration(c *gin.Context)
	GetCelebrations(c *gin.Context)

	// follower
	GetAllFollowers(c *gin.Context)
	GetSingleFollower(c *gin.Context)
	GetFriends(c *gin.Context)

	// Middleware
	AuthenticatedUserMiddleware() gin.HandlerFunc

	// Profile
	UpdateUserProfile(c *gin.Context)
	UploadUserProfilePicture(c *gin.Context)
	BlockUser(c *gin.Context)
	UnBlockUser(c *gin.Context)
	FollowUser(c *gin.Context)
	UnFollowUser(c *gin.Context)

	// Settings
	LogoutMiddleware() gin.HandlerFunc
	ChangePassword(c *gin.Context)
	AddAreaOfInterests(c *gin.Context)
	AddPreferredLanguage(c *gin.Context)
	AddNotificationPreference(c *gin.Context)
	TogglePushNotification(c *gin.Context)
	ToggleLikesNotification(c *gin.Context)
	ToggleCommentsNotification(c *gin.Context)
	ToggleTagsAndMentionNotification(c *gin.Context)
	ToggleRepostNotification(c *gin.Context)
	ToggleDirectMessageNotification(c *gin.Context)
	ToggleLiveNotification(c *gin.Context)
	ToggleNewFollower(c *gin.Context)
	SetContentSettings(c *gin.Context)
	SetBannedWords(c *gin.Context)

	// Web socket
	WebSocketHandler(c *gin.Context)

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

func getPagingInfo(c *gin.Context) dtos.APIPagingDto {
	var paging dtos.APIPagingDto

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	paging.Filter = c.Query("filter")
	paging.Cursor = c.Query("cursor")

	paging.Limit = limit
	paging.Page = page
	return paging
}
