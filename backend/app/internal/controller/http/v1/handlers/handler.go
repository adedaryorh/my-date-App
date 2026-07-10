package handlers

import (
	"fmt"
	"strconv"

	"backend.app/configs"
	"backend.app/database"
	"backend.app/internal/core"
	"backend.app/internal/dtos"
	"backend.app/pkg/logger"
	"backend.app/pkg/middleware"
	"backend.app/pkg/response"
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

	// AI
	GetUserRecommendations(c *gin.Context)
	ModerateCelebration(c *gin.Context)
	IndexCelebration(c *gin.Context)
	SearchCelebrations(c *gin.Context)
	LogInteraction(c *gin.Context)
	HealthCheck(c *gin.Context)

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

// Implement the Operations interface by delegating to the core methods and handling responses
func (h *Handler) GetUserRecommendations(c *gin.Context) {
	var input *dtos.RecommendUsersRequest
	if err := c.Bind(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	user := c.MustGet("authUser").(models.User) // auth user

	// Override user_id from token for security
	input.UserID = user.ID.String()

	result := h.core.GetUserRecommendations(c.Request.Context(), input.UserID, input.Limit)
	if result != nil {
		c.JSON(result.Code, result)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to get user recommendations"), "Failed to get user recommendations")
	c.JSON(result.Code, result)
}

func (h *Handler) ModerateCelebration(c *gin.Context) {
	var input *dtos.ModerateCelebrationRequest
	if err := c.Bind(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	result := h.core.ModerateCelebration(c.Request.Context(), input.Text)
	if result != nil {
		c.JSON(result.Code, result)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to moderate celebration"), "Failed to moderate celebration")
	c.JSON(result.Code, result)
}

func (h *Handler) IndexCelebration(c *gin.Context) {
	var input *dtos.IndexCelebrationRequest
	if err := c.Bind(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	result := h.core.IndexCelebration(c.Request.Context(), input.CelebrationID, input.Text)
	if result != nil {
		c.JSON(result.Code, result)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to index celebration"), "Failed to index celebration")
	c.JSON(result.Code, result)
}

func (h *Handler) SearchCelebrations(c *gin.Context) {
	var input *dtos.SearchCelebrationsRequest
	if err := c.Bind(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	result := h.core.SearchCelebrations(c.Request.Context(), input.Query, input.Limit)
	if result != nil {
		c.JSON(result.Code, result)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to search celebrations"), "Failed to search celebrations")
	c.JSON(result.Code, result)
}

func (h *Handler) LogInteraction(c *gin.Context) {
	var input *dtos.LogInteractionRequest
	if err := c.Bind(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	result := h.core.LogInteraction(c.Request.Context(), input.UserID, input.TargetID, input.Action, input.Metadata)
	if result != nil {
		c.JSON(result.Code, result)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to log interaction"), "Failed to log interaction")
	c.JSON(result.Code, result)
}

func (h *Handler) HealthCheck(c *gin.Context) {
	result := h.core.HealthCheck(c.Request.Context())
	if result != nil {
		c.JSON(result.Code, result)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to check health"), "Failed to check health")
	c.JSON(result.Code, result)
}

// Delegate other methods to core handler (keeping existing functionality)
func (h *Handler) SignUpUser(c *gin.Context) {
	h.core.SignUpUser(c)
}

func (h *Handler) SignUpBusiness(c *gin.Context) {
	h.core.SignUpBusiness(c)
}

func (h *Handler) Login(c *gin.Context) {
	h.core.Login(c)
}

func (h *Handler) ConfirmPhone(c *gin.Context) {
	h.core.ConfirmPhone(c)
}

func (h *Handler) SendResetPasswordToken(c *gin.Context) {
	h.core.SendResetPasswordToken(c)
}

func (h *Handler) ResetPassword(c *gin.Context) {
	h.core.ResetPassword(c)
}

func (h *Handler) Me(c *gin.Context) {
	h.core.Me(c)
}

func (h *Handler) GetAllBlockedUsers(c *gin.Context) {
	h.core.GetAllBlockedUsers(c)
}

func (h *Handler) GetBlockedUser(c *gin.Context) {
	h.core.GetBlockedUser(c)
}

func (h *Handler) CreateCelebration(c *gin.Context) {
	h.core.CreateCelebration(c)
}

func (h *Handler) GetCelebrations(c *gin.Context) {
	h.core.GetCelebrations(c)
}

func (h *Handler) GetAllFollowers(c *gin.Context) {
	h.core.GetAllFollowers(c)
}

func (h *Handler) GetSingleFollower(c *gin.Context) {
	h.core.GetSingleFollower(c)
}

func (h *Handler) GetFriends(c *gin.Context) {
	h.core.GetFriends(c)
}

func (h *Handler) AuthenticatedUserMiddleware() gin.HandlerFunc {
	return h.core.AuthenticatedUserMiddleware()
}

func (h *Handler) UpdateUserProfile(c *gin.Context) {
	h.core.UpdateUserProfile(c)
}

func (h *Handler) UploadUserProfilePicture(c *gin.Context) {
	h.core.UploadUserProfilePicture(c)
}

func (h *Handler) BlockUser(c *gin.Context) {
	h.core.BlockUser(c)
}

func (h *Handler) UnBlockUser(c *gin.Context) {
	h.core.UnBlockUser(c)
}

func (h *Handler) FollowUser(c *gin.Context) {
	h.core.FollowUser(c)
}

func (h *Handler) UnFollowUser(c *gin.Context) {
	h.core.UnFollowUser(c)
}

func (h *Handler) LogoutMiddleware() gin.HandlerFunc {
	return h.core.LogoutMiddleware()
}

func (h *Handler) ChangePrice(c *gin.Context) {
	h.core.ChangePrice(c)
}

func (h *Handler) AddAreaOfInterests(c *gin.Context) {
	h.core.AddAreaOfInterests(c)
}

func (h *Handler) AddPreferredLanguage(c *gin.Context) {
	h.core.AddPreferredLanguage(c)
}

func (h *Handler) AddNotificationPreference(c *gin.Context) {
	h.core.AddNotificationPreference(c)
}

func (h *Handler) TogglePushNotification(c *gin.Context) {
	h.core.TogglePushNotification(c)
}

func (h *Handler) ToggleLikesNotification(c *gin.Context) {
	h.core.ToggleLikesNotification(c)
}

func (h *Handler) ToggleCommentsNotification(c *gin.Context) {
	h.core.ToggleCommentsNotification(c)
}

func (h *Handler) ToggleTagsAndMentionNotification(c *gin.Context) {
	h.core.ToggleTagsAndMentionNotification(c)
}

func (h *Handler) ToggleRepostNotification(c *gin.Context) {
	h.core.ToggleRepostNotification(c)
}

func (h *Handler) ToggleDirectMessageNotification(c *gin.Context) {
	h.core.ToggleDirectMessageNotification(c)
}

func (h *Handler) ToggleLiveNotification(c *gin.Context) {
	h.core.ToggleLiveNotification(c)
}

func (h *Handler) ToggleNewFollower(c *gin.Context) {
	h.core.ToggleNewFollower(c)
}

func (h *Handler) SetContentSettings(c *gin.Context) {
	h.core.SetContentSettings(c)
}

func (h *Handler) SetBannedWords(c *gin.Context) {
	h.core.SetBannedWords(c)
}

func (h *Handler) WebSocketHandler(c *gin.Context) {
	h.core.WebSocket(c)
}

func (h *Handler) GetUserWallet(c *gin.Context) {
	h.core.GetUserWallet(c)
}
