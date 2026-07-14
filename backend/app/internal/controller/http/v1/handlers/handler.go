package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/configs"
	"backend.app/database"
	"backend.app/internal/core"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/pkg/logger"
	"backend.app/pkg/middleware"
	"backend.app/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

type Handler struct {
	core        core.Operations
	log         *logger.Logger
	config      *configs.Config
	oauthConfig *oauth2.Config
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

	// Notifications
	GetNotificationById(c *gin.Context)
	GetAllNotifications(c *gin.Context)
	MarkNotificationAsRead(c *gin.Context)
	DeleteNotification(c *gin.Context)

	// RBAC
	GetAllUsers(c *gin.Context)
	UpdateUserRole(c *gin.Context)
	DeleteUser(c *gin.Context)

	// OAuth
	GoogleLogin(c *gin.Context)
	GoogleCallback(c *gin.Context)
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
		oauthConfig: &oauth2.Config{
			ClientID:     config.GoogleClientID,
			ClientSecret: config.GoogleClientSecret,
			RedirectURL:  config.AppHost + ":" + config.Port + "/v1/auth/google/callback",
			Scopes:       []string{"openid", "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		},
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

const oauthStateCookieName = "oauth_state"

func generateStateOauthCookie() string {
	var b []byte
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	state := base64.URLEncoding.EncodeToString(b)
	return state
}

func setStateOauthCookie(c *gin.Context, state string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     oauthStateCookieName,
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		// Secure:   true, // Uncomment when serving via HTTPS
		MaxAge: 3600,
	})
}

func getStateOauthCookie(c *gin.Context) (string, error) {
	cookie, err := c.Cookie(oauthStateCookieName)
	if err != nil {
		return "", err
	}
	return cookie, nil
}

func decodeGoogleIDToken(idToken string) (string, error) {
	parts := strings.Split(idToken, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid ID token")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("failed to decode payload: %w", err)
	}
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return "", fmt.Errorf("failed to unmarshal claims: %w", err)
	}
	if sub, ok := claims["sub"].(string); ok {
		return sub, nil
	}
	return "", fmt.Errorf("sub claim not found")
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
	user := c.MustGet("authUser").(models.User)
	input.UserID = user.ID.String()

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

func (h *Handler) GetAllUsers(c *gin.Context) {
	query := getPagingInfo(c)
	user := c.MustGet("authUser").(models.User)

	result := h.core.GetAllUsers(c.Request.Context(), &user, &query)
	if result != nil {
		c.JSON(result.Code, result)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to get users"), "Failed to get users")
	c.JSON(result.Code, result)
}

func (h *Handler) UpdateUserRole(c *gin.Context) {
	var input struct {
		UserID string `json:"user_id" binding:"required"`
		Role   string `json:"role" binding:"required,oneof=user admin moderator"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	user := c.MustGet("authUser").(models.User)
	if user.Role != "admin" {
		result := response.ForbiddenResponse(fmt.Errorf("insufficient permissions"), "Access denied")
		c.JSON(result.Code, result)
		return
	}

	result := h.core.UpdateUserRole(c.Request.Context(), input.UserID, input.Role)
	if result != nil {
		c.JSON(result.Code, result)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to update user role"), "Failed to update user role")
	c.JSON(result.Code, result)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		result := response.BadRequestResponse(fmt.Errorf("user_id is required"), constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	user := c.MustGet("authUser").(models.User)
	if user.Role != "admin" {
		result := response.ForbiddenResponse(fmt.Errorf("insufficient permissions"), "Access denied")
		c.JSON(result.Code, result)
		return
	}

	// Optional: prevent self-delete
	if user.ID.String() == userID {
		result := response.BadRequestResponse(fmt.Errorf("cannot delete yourself"), constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	result := h.core.DeleteUser(c.Request.Context(), userID)
	if result != nil {
		c.JSON(result.Code, result)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to delete user"), "Failed to delete user")
	c.JSON(result.Code, result)
}

// GoogleLogin handles the Google login request
func (h *Handler) GoogleLogin(c *gin.Context) {
	state := generateStateOauthCookie()
	setStateOauthCookie(c, state)
	url := h.oauthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback handles Google callback
func (h *Handler) GoogleCallback(c *gin.Context) {
	state, err := getStateOauthCookie(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if c.Query("state") != state {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	code := c.Query("code")
	if code == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.oauthConfig.Exchange(c.Request.Context(), code)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Extract Google ID from ID token
	var googleID string
	if token.Extra("id_token") != nil {
		idToken := token.Extra("id_token").(string)
		googleID, err = decodeGoogleIDToken(idToken)
		if err != nil {
			h.log.Error("Failed to decode Google ID token", zap.Error(err))
			// Continue without Google ID - not critical for basic functionality
		}
	}

	// Get user info from Google
	client := &http.Client{}
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email     string `json:"email"`
		FirstName string `json:"given_name"`
		LastName  string `json:"family_name"`
		Picture   string `json:"picture"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Check if user already exists by email
	existingUser, err := h.core.GetUserByEmail(c.Request.Context(), userInfo.Email)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var user *models.User
	if existingUser != nil {
		// User exists, update GoogleID if not set
		if existingUser.GoogleID == "" && googleID != "" {
			// Update the user's GoogleID
			updateErr := h.core.UpdateUser(c.Request.Context(), existingUser.ID.String(), map[string]interface{}{"google_id": googleID})
			if updateErr != nil {
				h.log.Warn("Failed to update GoogleID for existing user", zap.Error(updateErr), zap.String("user_id", existingUser.ID.String()))
			}
			// Refresh the user object with the updated GoogleID
			existingUser.GoogleID = googleID
		}
		user = existingUser
	} else {
		// Create new user
		user, err = h.core.CreateUserFromGoogle(c.Request.Context(), userInfo.Email, userInfo.FirstName, userInfo.LastName, userInfo.Picture, googleID)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	// Generate JWT token for the user
	tokenString, err := h.core.TokenService.GenerateToken(user.ID.String())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Return the token to the client
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user":  user,
	})
}

// GetNotificationById returns a notification by its ID
func (h *Handler) GetNotificationById(c *gin.Context) {
	h.core.GetNotificationById(c)
}

// GetAllNotifications returns all notifications with pagination
func (h *Handler) GetAllNotifications(c *gin.Context) {
	h.core.GetAllNotifications(c)
}

// DeleteNotification deletes a notification by its ID
func (h *Handler) DeleteNotification(c *gin.Context) {
	h.core.DeleteNotification(c)
}

// MarkNotificationAsRead marks a notification as read
func (h *Handler) MarkNotificationAsRead(c *gin.Context) {
	h.core.MarkNotificationAsRead(c)
}
