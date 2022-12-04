package users

import (
	"celebut-api/internal/dtos"
	"celebut-api/internal/mappers"
	"celebut-api/internal/models"
	"celebut-api/internal/repo"
	"celebut-api/internal/services/file"
	"celebut-api/internal/usecase"
	"celebut-api/internal/validators"
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"mime/multipart"
	"path/filepath"
	"strings"
)

type User interface {
	ValidateUser(ctx context.Context, userID string) (*models.User, error)
	GetUser(ctx context.Context, userID string) (*dtos.UserProfile, error)
	EditUser(ctx context.Context, user *models.User) error
	UploadProfileImage(ctx context.Context, image multipart.FileHeader) (string, error)
}

type UserService struct {
	usecase      usecase.User
	repo         repo.User
	mapper       mappers.UserMapper
	uploadClient file.UploadFileClient
}

func NewUserService(r repo.User, uploadClient file.UploadFileClient, usecase usecase.User, mapper mappers.UserMapper) *UserService {
	return &UserService{repo: r, uploadClient: uploadClient, usecase: usecase, mapper: mapper}
}

func (us *UserService) EditUser(ctx context.Context, user *models.User) error {
	if user == nil || user.ID == 0 {
		return errors.New("user service - user/user ID not defined")
	}

	if user.ProfileImage != nil {
		url, err := us.UploadProfileImage(ctx, *user.ProfileImage)
		if err != nil {
			return fmt.Errorf("user service - error updating user: %w", err)
		}

		user.ProfileImageURL = &url
	}

	err := us.repo.UpdateUser(ctx, user, false)
	if err != nil {
		return fmt.Errorf("user service - error updating user: %w", err)
	}

	return nil
}

func (us *UserService) GetUser(ctx context.Context, userID string) (*dtos.UserProfile, error) {
	user, err := us.usecase.UserByField(ctx, "user_id", userID)
	if err != nil {
		return nil, fmt.Errorf("user service - error fetching user: %w", err)
	}

	dto := us.mapper.MapToUserProfileDto(*user)

	return &dto, nil
}

func (us *UserService) ValidateUser(ctx context.Context, userID string) (*models.User, error) {
	user, err := us.usecase.UserByField(ctx, "user_id", userID)
	if err != nil {
		return nil, fmt.Errorf("user service - error fetching user: %w", err)
	}

	return user, nil
}

func (us *UserService) UploadProfileImage(ctx context.Context, image multipart.FileHeader) (string, error) {
	imgFile, err := image.Open()
	defer imgFile.Close()

	if err != nil {
		return "", fmt.Errorf("unable to open file: %w", err)
	}

	// We only accept PNG, JPEG and JPG for profile images for now...
	_, err = validators.ValidateFileExtension(imgFile, []string{"image/png", "image/jpeg", "image/jpg"})
	if err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}

	fileName, err := generateRandomFileName()
	if err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}

	ext := filepath.Ext(image.Filename)
	fileName = fmt.Sprintf("profile-images/%s.%s", fileName, ext)
	profileImageURL, err := us.uploadClient.UploadToBucket(ctx, "celebut", fileName, imgFile)

	if err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}

	return *profileImageURL, nil
}

func generateRandomFileName() (string, error) {
	newUUID, err := uuid.DefaultGenerator.NewV4()

	if err != nil {
		return "", fmt.Errorf("unable to generate UUID: %s", err)
	}

	fileName := strings.ReplaceAll(newUUID.String(), "-", "")

	return fileName, nil
}
