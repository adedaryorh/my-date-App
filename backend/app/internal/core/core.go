package core

import (
	"context"
	"errors"
	"fmt"

	"backend.app/configs"
	"backend.app/database"
	"backend.app/database/postgres"
	"backend.app/integrations/analytics/mix-panel"
	"backend.app/integrations/email"
	"backend.app/integrations/sms"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/internal/repo"
	"backend.app/internal/services/aiclient"
	"backend.app/internal/services/redisservice"
	"backend.app/internal/services/tokenservice"
	"backend.app/internal/services/upload"
	"backend.app/pkg/logger"
	"backend.app/pkg/middleware"
	"backend.app/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// Add AI service client to the Core struct
type Core struct {
	middleware    *middleware.Middleware
	redisService  redisservice.Redis
	TokenService  tokenservice.TokenService
	uploadService upload.Uploader
	aiClient      *aiclient.AIServiceClient
	config        *configs.Config
	log           *logger.Logger
	repo          repo.Operations
	phoneService  map[string]sms.PhoneService
	emailService  map[string]email.EmailService
	mixPanel      *mixpanel.MixPanel
	postgres      *postgres.Postgres
}

// Update the Operations interface to include AI methods
type Operations interface {
	Middleware() *middleware.Middleware

	// auth
	SignUpUser(ctx context.Context, data *dtos.UserSignUp) *dtos.ResponseObject
	SignUpBusiness(ctx context.Context, data *dtos.BusinessSignUp) *dtos.ResponseObject
	ConfirmPhone(ctx context.Context, data *dtos.ConfirmPhoneNumber) *dtos.ResponseObject
	ConfirmEmail(ctx context.Context, data *dtos.ConfirmEmail) *dtos.ResponseObject
	Login(ctx context.Context, data models.SignInDto) *dtos.ResponseObject
	SendResetPasswordToken(ctx context.Context, email string) *dtos.ResponseObject
	ResetPassword(ctx context.Context, data *dtos.ResetPassword) *dtos.ResponseObject
	AuthenticateUser(ctx context.Context, user *models.User) (*models.AuthenticatedUser, error)

	// blocked
	BlockUser(ctx context.Context, userId uuid.UUID, user *models.User) *dtos.ResponseObject
	GetBlockedUser(ctx context.Context, userId uuid.UUID, user *models.User) *dtos.ResponseObject
	GetAllBlockedUsers(ctx context.Context, user *models.User, query *dtos.APIPagingDto) *dtos.ResponseObject
	UnBlockUser(ctx context.Context, userId uuid.UUID, user *models.User) *dtos.ResponseObject

	// celebration
	CreateCelebrationDto(ctx context.Context, user *models.User, data *dtos.CreateCelebrationDto) *dtos.ResponseObject
	GetAllCelebrations(ctx context.Context, user *models.User, query *dtos.APIPagingDto) *dtos.ResponseObject
	GetSingleCelebration(ctx context.Context, celebrationId uuid.UUID) *dtos.ResponseObject
	DeclineCelebration(ctx context.Context, celebrationId uuid.UUID, user *models.User) *dtos.ResponseObject
	AcceptCelebration(ctx context.Context, celebrationId uuid.UUID, user *models.User) *dtos.ResponseObject

	// follower
	FollowUser(ctx context.Context, userId uuid.UUID, user *models.User) *dtos.ResponseObject
	UnFollowUser(ctx context.Context, userId uuid.UUID, user *models.User) *dtos.ResponseObject
	GetAllFollowers(ctx context.Context, user *models.User, query *dtos.APIPagingDto) *dtos.ResponseObject
	GetSingleFollower(ctx context.Context, userId uuid.UUID, user *models.User) *dtos.ResponseObject
	GetFriends(ctx context.Context, user *models.User, query *dtos.APIPagingDto) *dtos.ResponseObject

	// AI
	GetUserRecommendations(ctx context.Context, userID string, limit int) *dtos.ResponseObject
	ModerateCelebration(ctx context.Context, text string) *dtos.ResponseObject
	IndexCelebration(ctx context.Context, celebrationID string, text string) *dtos.ResponseObject
	SearchCelebrations(ctx context.Context, query string, limit int) *dtos.ResponseObject
	LogInteraction(ctx context.Context, userID string, targetID string, action string, metadata map[string]interface{}) *dtos.ResponseObject
	HealthCheck(ctx context.Context) *dtos.ResponseObject

	// notification
	GetNotificationById(ctx context.Context, notificationId uuid.UUID) *dtos.ResponseObject
	GetAllNotifications(ctx context.Context, query *dtos.APIPagingDto) *dtos.ResponseObject
	MarkNotificationAsRead(ctx context.Context, notificationId uuid.UUID) *dtos.ResponseObject
	DeleteNotification(ctx context.Context, notificationId uuid.UUID) *dtos.ResponseObject

	// settings
	ChangePassword(ctx context.Context, data *dtos.ChangePassword, user *models.User) *dtos.ResponseObject
	AddAreaOfInterests(ctx context.Context, data *dtos.AreaOfInterest, user *models.User) *dtos.ResponseObject
	AddPreferredLanguage(ctx context.Context, data *dtos.Language, user *models.User) *dtos.ResponseObject
	AddNotificationPreference(ctx context.Context, data *dtos.NotificationPreference, user *models.User) *dtos.ResponseObject
	TogglePushNotification(ctx context.Context, user *models.User) *dtos.ResponseObject
	ToggleLikesNotification(ctx context.Context, user *models.User) *dtos.ResponseObject
	ToggleCommentsNotification(ctx context.Context, user *models.User) *dtos.ResponseObject
	ToggleTagsAndMentionsNotification(ctx context.Context, user *models.User) *dtos.ResponseObject
	ToggleRepostNotification(ctx context.Context, user *models.User) *dtos.ResponseObject
	ToggleDirectMessageNotification(ctx context.Context, user *models.User) *dtos.ResponseObject
	ToggleLiveNotification(ctx context.Context, user *models.User) *dtos.ResponseObject
	ToggleNewFollowerNotification(ctx context.Context, user *models.User) *dtos.ResponseObject
	SetContentSettings(ctx context.Context, user *models.User, data *dtos.ContentSettingsDto) *dtos.ResponseObject
	SetBannedWords(ctx context.Context, data *dtos.SetBannedWordsDto, user *models.User) *dtos.ResponseObject

	// profile
	UploadUserProfileImage(ctx context.Context, user *models.User, data *dtos.UploadImage) *dtos.ResponseObject
	UpdateUserProfile(ctx context.Context, user *models.User, data *dtos.UpdateUserProfile) *dtos.ResponseObject

	// wallet
	GetUserWallet(ctx context.Context, user *models.User) *dtos.ResponseObject

	// OAuth
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	CreateUserFromGoogle(ctx context.Context, email string, firstName string, lastName string, picture string, googleID string) (*models.User, error)
	UpdateUser(ctx context.Context, userID string, fields map[string]interface{}) *dtos.ResponseObject

	// admin
	GetAllUsers(ctx context.Context, user *models.User, query *dtos.APIPagingDto) *dtos.ResponseObject
	UpdateUserRole(ctx context.Context, userID string, role string) *dtos.ResponseObject
	DeleteUser(ctx context.Context, userID string) *dtos.ResponseObject
	GetNearbyCelebrations(ctx context.Context, latitude, longitude, radiusKM float64, limit int) *dtos.ResponseObject
	GetModerationQueue(ctx context.Context, user *models.User, status string, limit int) *dtos.ResponseObject
	ResolveModeration(ctx context.Context, user *models.User, queueID uuid.UUID, decision string) *dtos.ResponseObject

	// websocket
	HandleWebsocketConnection(ctx *gin.Context, conn *websocket.Conn)
}

// Update the NewCore function to initialize the AI client
func NewCore(config *configs.Config, log *logger.Logger, db *database.DB, middleware *middleware.Middleware) Operations {
	redis := redisservice.Redis{Client: db.Redis.Client}
	repo := repo.NewRepo(db, log)

	// Initialize AI service client
	aiClient := aiclient.NewAIServiceClient(config, log)

	c := Core{
		config:        config,
		log:           log,
		redisService:  redis,
		TokenService:  tokenservice.NewTokenService(&redis, config, repo),
		uploadService: *upload.NewUpload(config, log),
		aiClient:      aiClient,
		middleware:    middleware,
		repo:          repo,
		phoneService: map[string]sms.PhoneService{
			"twilio": sms.NewTwilioService(config),
		},
		emailService: map[string]email.EmailService{"smtp": email.NewSMTP(config)},
		mixPanel:     mixpanel.NewMixPanel(config, log),
		postgres:     db.Postgres,
	}
	op := Operations(&c)

	return op
}

func (c *Core) Middleware() *middleware.Middleware {
	return c.middleware
}

func (c *Core) UploadFileToAwsS3(file upload.FileInput, maxFileSize int64, allowedFileTypes []upload.AttachmentKind) (*models.AWSObjectUrl, error) {
	if err := c.uploadService.ValidateFile(file, maxFileSize, allowedFileTypes); err != nil {
		return nil, err
	}
	savedFile, err := c.uploadService.UploadFile(file)
	if err != nil {
		return nil, errors.New("file upload failed")
	}
	return savedFile, nil
}

// Implement the AI methods in the Core struct

// GetUserRecommendations gets recommended users to follow for a given user
func (c *Core) GetUserRecommendations(ctx context.Context, userID string, limit int) *dtos.ResponseObject {
	c.log.Info("Getting user recommendations via AI service",
		zap.String("user_id", userID),
		zap.Int("limit", limit))

	// Call AI service for recommendations
	result, err := c.aiClient.RecommendUsers(ctx, &aiclient.RecommendUsersRequest{
		UserID:          userID,
		Limit:           limit,
		ExcludeFollowed: true,
		ExcludeBlocked:  true,
	})

	if err == nil {
		return response.SuccessResponse("User recommendations retrieved successfully", result)
	}

	// Handle error case
	return response.ServerErrorResponse(fmt.Errorf("failed to get user recommendations"), "Failed to get user recommendations")
}

// ModerateCelebration checks if celebration content is appropriate
func (c *Core) ModerateCelebration(ctx context.Context, text string) *dtos.ResponseObject {
	c.log.Info("Checking celebration content for toxicity via AI service",
		zap.String("text_preview", truncateString(text, 50)))

	// Call AI service for content moderation
	result, err := c.aiClient.ModerateCelebration(ctx, &aiclient.ModerateCelebrationRequest{
		Text: text,
	})

	if err == nil {
		return response.SuccessResponse("Content moderation completed", result)
	}

	// Handle error case
	return response.ServerErrorResponse(fmt.Errorf("failed to moderate celebration"), "Failed to moderate celebration")
}

// IndexCelebration indexes celebration content for semantic search
func (c *Core) IndexCelebration(ctx context.Context, celebrationID string, text string) *dtos.ResponseObject {
	c.log.Info("Indexing celebration for search via AI service",
		zap.String("celebration_id", celebrationID),
		zap.String("text_preview", truncateString(text, 50)))

	// Call AI service to index celebration
	result, err := c.aiClient.IndexCelebration(ctx, &aiclient.IndexCelebrationRequest{
		CelebrationID: celebrationID,
		Text:          text,
	})

	if err == nil {
		return response.SuccessResponse("Celebrity indexed successfully", result)
	}

	// Handle error case
	return response.ServerErrorResponse(fmt.Errorf("failed to index celebration"), "Failed to index celebration")
}

// SearchCelebrations searches for celebrations using semantic similarity
func (c *Core) SearchCelebrations(ctx context.Context, query string, limit int) *dtos.ResponseObject {
	c.log.Info("Searching celebrations by text via AI service",
		zap.String("query", query),
		zap.Int("limit", limit))

	// Call AI service to search celebrations
	result, err := c.aiClient.SearchCelebrations(ctx, &aiclient.SearchCelebrationsRequest{
		Query: query,
		Limit: limit,
	})

	if err == nil {
		return response.SuccessResponse("Celebrity search completed", result)
	}

	// Handle error case
	return response.ServerErrorResponse(fmt.Errorf("failed to search celebrations"), "Failed to search celebrations")
}

// LogInteraction logs a user interaction for the feedback loop
func (c *Core) LogInteraction(ctx context.Context, userID string, targetID string, action string, metadata map[string]interface{}) *dtos.ResponseObject {
	c.log.Info("Logging user interaction via AI service",
		zap.String("user_id", userID),
		zap.String("target_id", targetID),
		zap.String("action", action))

	// Call AI service to log interaction
	result, err := c.aiClient.LogInteraction(ctx, &aiclient.LogInteractionRequest{
		UserID:   userID,
		TargetID: targetID,
		Action:   action,
		Metadata: metadata,
	})

	if err == nil {
		return response.SuccessResponse("Interaction logged successfully", result)
	}

	// Handle error case
	return response.ServerErrorResponse(fmt.Errorf("failed to log interaction"), "Failed to log interaction")
}

// HealthCheck checks if the AI service is healthy
func (c *Core) HealthCheck(ctx context.Context) *dtos.ResponseObject {
	c.log.Info("Checking AI service health")
	result, err := c.aiClient.HealthCheck(ctx)
	if err != nil {
		return response.ServerErrorResponse(fmt.Errorf("AI service health check failed: %w", err), "AI service is unavailable")
	}
	return response.SuccessResponse("AI service is healthy", result)
}

// Helper function to truncate string for logging
func truncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length] + "..."
}

// GetAllUsers returns all users with pagination
func (c *Core) GetAllUsers(ctx context.Context, user *models.User, query *dtos.APIPagingDto) *dtos.ResponseObject {
	c.log.Info("Retrieving users list",
		zap.String("requester_id", user.ID.String()),
		zap.String("requester_email", user.Email),
		zap.Int("limit", query.Limit),
		zap.Int("page", query.Page),
		zap.String("action", "list_users"))

	// Check if user has admin permissions
	if user.Role != "admin" {
		c.log.Warn("Unauthorized attempt to list users",
			zap.String("requester_id", user.ID.String()),
			zap.String("requester_role", user.Role))
		return response.ForbiddenResponse(fmt.Errorf("insufficient permissions"), "Access denied")
	}

	result, err := c.repo.GetAllUsers(ctx, user, query)
	if err == nil {
		c.log.Info("Successfully retrieved users list",
			zap.String("requester_id", user.ID.String()),
			zap.Int("count", len(result.Users)),
			zap.String("action", "list_users"))
		return response.SuccessResponse("Users retrieved successfully", result)
	}

	// Handle error case
	c.log.Error("Failed to retrieve users list",
		zap.String("requester_id", user.ID.String()),
		zap.String("action", "list_users"))
	return response.ServerErrorResponse(fmt.Errorf("failed to get users"), "Failed to get users")
}

// UpdateUserRole updates a user's role
func (c *Core) UpdateUserRole(ctx context.Context, userID string, role string) *dtos.ResponseObject {
	c.log.Info("Attempting to update user role",
		zap.String("target_user_id", userID),
		zap.String("role", role),
		zap.String("action", "update_user_role"))

	// Validate role
	validRoles := map[string]bool{
		"user":      true,
		"admin":     true,
		"moderator": true,
	}
	if !validRoles[role] {
		c.log.Warn("Invalid role specified for user update",
			zap.String("target_user_id", userID),
			zap.String("role", role),
			zap.String("action", "update_user_role"))
		return response.BadRequestResponse(fmt.Errorf("invalid role"), "Invalid role specified")
	}

	err := c.repo.UpdateUserRole(ctx, userID, role)
	if err != nil {
		c.log.Error("Failed to update user role",
			zap.String("target_user_id", userID),
			zap.String("role", role),
			zap.Error(err),
			zap.String("action", "update_user_role"))
		return response.ServerErrorResponse(err, "Failed to update user role")
	}

	c.log.Info("Successfully updated user role",
		zap.String("target_user_id", userID),
		zap.String("role", role),
		zap.String("action", "update_user_role"))
	return response.SuccessResponse("User role updated successfully", map[string]string{
		"user_id": userID,
		"role":    role,
	})
}

// DeleteUser deletes a user by ID
func (c *Core) DeleteUser(ctx context.Context, userID string) *dtos.ResponseObject {
	c.log.Info("Attempting to delete user",
		zap.String("target_user_id", userID),
		zap.String("action", "delete_user"))

	// Prevent self-deletion warning (in a real app, we'd compare with current user ID from context)
	// Since we don't have current user in this method signature, we log a warning to highlight
	// that self-deletion prevention should happen at the API layer
	c.log.Warn("Delete user requested - self-deletion prevention should be verified at API layer",
		zap.String("target_user_id", userID),
		zap.String("action", "delete_user"))

	err := c.repo.DeleteUser(ctx, userID)
	if err != nil {
		c.log.Error("Failed to delete user",
			zap.String("target_user_id", userID),
			zap.Error(err),
			zap.String("action", "delete_user"))
		return response.ServerErrorResponse(err, "Failed to delete user")
	}

	c.log.Info("Successfully deleted user",
		zap.String("target_user_id", userID),
		zap.String("action", "delete_user"))
	return response.SuccessResponse("User deleted successfully", map[string]string{
		"user_id": userID,
	})
}

// GetUserByEmail returns a user by email
func (c *Core) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	c.log.Info("Getting user by email", zap.String("email", email))
	return c.repo.GetUserByEmail(ctx, email)
}

// CreateUserFromGoogle creates a new user from Google OAuth data
func (c *Core) CreateUserFromGoogle(ctx context.Context, email string, firstName string, lastName string, picture string, googleID string) (*models.User, error) {
	c.log.Info("Creating user from Google", zap.String("email", email))
	return c.repo.CreateUserFromGoogle(ctx, email, firstName, lastName, picture, googleID)
}

// UpdateUser updates a user's fields
func (c *Core) UpdateUser(ctx context.Context, userID string, fields map[string]interface{}) *dtos.ResponseObject {
	c.log.Info("Updating user",
		zap.String("user_id", userID))

	// In a real implementation, we would get the current user from context and check permissions
	// For now, we'll assume the caller has validated permissions

	id, parseErr := uuid.Parse(userID)
	if parseErr != nil {
		return response.BadRequestResponse(fmt.Errorf("invalid user ID"))
	}
	err := c.repo.UpdateUser(ctx, id, fields)
	if err != nil {
		return response.ServerErrorResponse(err, "Failed to update user")
	}

	return response.SuccessResponse("User updated successfully", map[string]string{
		"user_id": userID,
	})
}

// The rest of the file would contain the existing implementations for other features...
// For brevity, I'm not including them here as they should already exist in the file
