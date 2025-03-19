package core

import (
	"backend.app/database/postgres"
	"context"
	"errors"

	"backend.app/configs"
	"backend.app/database"
	mixpanel "backend.app/integrations/analytics/mix-panel"
	"backend.app/integrations/sms"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/internal/repo"
	"backend.app/internal/services/redisservice"
	"backend.app/internal/services/tokenservice"
	"backend.app/internal/services/upload"
	"backend.app/pkg/logger"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Core struct {
	middleware    *middleware.Middleware
	redisService  redisservice.Redis
	TokenService  tokenservice.TokenService
	uploadService upload.Uploader
	config        *configs.Config
	log           *logger.Logger
	repo          repo.Operations
	phoneService  map[string]sms.PhoneService
	mixPanel      *mixpanel.MixPanel
	postgres      *postgres.Postgres
}

type Operations interface {
	Middleware() *middleware.Middleware

	// auth
	SignUpUser(ctx context.Context, data *dtos.UserSignUp) *dtos.ResponseObject
	SignUpBusiness(ctx context.Context, data *dtos.BusinessSignUp) *dtos.ResponseObject
	ConfirmPhone(ctx context.Context, data *dtos.ConfirmPhoneNumber) *dtos.ResponseObject
	Login(ctx context.Context, data models.SignInDto) *dtos.ResponseObject
	SendResetPasswordToken(ctx context.Context, email string) *dtos.ResponseObject
	ResetPassword(ctx context.Context, data *dtos.ResetPassword) *dtos.ResponseObject

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

	// notification

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

	// websocket
	HandleWebsocketConnection(ctx *gin.Context, conn *websocket.Conn)
}

func NewCore(config *configs.Config, log *logger.Logger, db *database.DB, middleware *middleware.Middleware) Operations {
	redis := redisservice.Redis{Client: db.Redis.Client}
	repo := repo.NewRepo(db, log)

	c := Core{
		config:        config,
		log:           log,
		redisService:  redis,
		TokenService:  tokenservice.NewTokenService(&redis, config, repo),
		uploadService: *upload.NewUpload(config, log),
		middleware:    middleware,
		repo:          repo,
		phoneService: map[string]sms.PhoneService{
			"twilio": sms.NewTwilioService(config),
		},
		mixPanel: mixpanel.NewMixPanel(config, log),
	}
	op := Operations(&c)

	return op
}
func (c *Core) Middleware() *middleware.Middleware {
	return c.middleware
}

func (c *Core) UploadFileToAwsS3(file upload.FileInput, maxFileSize int64, allowedFileTypes []upload.AttachmentKind) (*models.AWSObjectUrl, error) {
	//validate image
	err := c.uploadService.ValidateFile(file, maxFileSize, allowedFileTypes) // validating the file(s)
	if err != nil {
		c.log.Debug("UploadFileToAwsS3:validateFile [%v] : (%v)", file.Name, err)
		return nil, err
	}

	//upload file into destination  folder
	savedFile, err := c.uploadService.UploadFile(file)
	if err != nil {
		c.log.Debug("UploadFileToAwsS3:UploadFile file error: (%v)", err)
		return nil, errors.New("file upload failed")
	}

	return savedFile, nil
}
