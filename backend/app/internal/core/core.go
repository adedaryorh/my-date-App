package core

import (
	"context"
	"errors"

	"backend.app/configs"
	"backend.app/database"
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
}

type Operations interface {
	Middleware() *middleware.Middleware

	SignUpUser(ctx context.Context)
}

func NewCore(config *configs.Config, log *logger.Logger, db *database.DB, middleware *middleware.Middleware) Operations {
	redis := redisservice.Redis{Client: db.Redis.Client}
	repo := repo.NewRepo(db)

	c := Core{
		config:        config,
		log:           log,
		redisService:  redis,
		TokenService:  tokenservice.NewTokenService(&redis, repo),
		uploadService: *upload.NewUpload(config, log),
		middleware:    middleware,
		repo:          repo,
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
