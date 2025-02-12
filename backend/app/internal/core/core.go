package core

import (
	"context"
	"errors"

	"backend.app/configs"
	"backend.app/database"
	"backend.app/integrations/sms"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/internal/repo"
	"backend.app/internal/services/redisservice"
	"backend.app/internal/services/tokenservice"
	"backend.app/internal/services/upload"
	"backend.app/pkg/logger"
	"backend.app/pkg/middleware"
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

	// settings
	ChangePassword(ctx context.Context, data *dtos.ChangePassword, user *models.User) *dtos.ResponseObject

	// profile
	UploadUserProfileImage(ctx context.Context, user *models.User, data *dtos.UploadImage) *dtos.ResponseObject
	UpdateUserProfile(ctx context.Context, user *models.User, data *dtos.UpdateUserProfile) *dtos.ResponseObject

	// wallet
	GetUserWallet(ctx context.Context, user *models.User) *dtos.ResponseObject
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
