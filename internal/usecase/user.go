package usecase

import (
	"celebut-api/internal/models"
	"celebut-api/internal/repo"
	"celebut-api/internal/services/otp_generator"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
	"time"
)

// User -.
type User interface {
	CompleteRegistration(context.Context, *models.User) error
	Login(context.Context, string, string, int) (*models.User, error)
	ValidateExistenceByField(context.Context, string, string) error
	UserIsEnabled(ctx context.Context, field string, value string) error
	UserByField(ctx context.Context, field string, value string) (*models.User, error)
	UserByEmail(ctx context.Context, email string) (*models.User, error)
	RegisterUserWithOTP(context.Context, string, string, int) (*models.UserOTP, error)
	ValidateOTP(ctx context.Context, userID int, otp string, mode string) error
	UsersByField(ctx context.Context, field string, values []string) ([]models.User, error)
}

var (
	registerOTPMode = "register"

	statusPending = "pending"
	statusEnabled = "enabled"
)

// UserUseCase -.
type UserUseCase struct {
	userRepo   repo.User
	otpRepo    repo.UserOTP
	otpService otp_generator.Generator
}

// NewUserUseCase -.
func NewUserUseCase(r repo.User, otp repo.UserOTP, otps otp_generator.Generator) *UserUseCase {
	return &UserUseCase{
		userRepo:   r,
		otpRepo:    otp,
		otpService: otps,
	}
}

// CompleteRegistration - completes registration of a new user.
func (uc *UserUseCase) CompleteRegistration(ctx context.Context, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("unable to generate password hash: %w", err)
	}

	user.Password = string(hashedPassword)

	user.Status = statusEnabled

	err = uc.userRepo.UpdateUser(ctx, user, true)
	if err != nil {
		return fmt.Errorf("unable to create user: %w", err)
	}

	return nil
}

// Login -.
func (uc *UserUseCase) Login(ctx context.Context, email string, password string, accountTypeId int) (*models.User, error) {
	user, err := uc.userRepo.GetUserByField(ctx, "email", email)
	if err != nil {
		return nil, fmt.Errorf("user - login - s.userRepo.GetUserByEmail: %w", err)
	}

	if user == nil {
		return nil, errors.New("user - login - s.userRepo.GetUserByEmail: user is nil")
	}

	if user.AccountType.ID != accountTypeId {
		return nil, errors.New("invalid credentials: user authentication failed")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials: user authentication failed")
	}

	return user, nil
}

// ValidateExistenceByField -.
func (uc *UserUseCase) ValidateExistenceByField(ctx context.Context, field string, value string) error {
	user, err := uc.userRepo.GetUserByField(ctx, field, value)

	if err != nil {
		return err
	}

	if user != nil {
		return fmt.Errorf("%s already exists", field)
	}

	return nil
}

// UserIsEnabled -.
func (uc *UserUseCase) UserIsEnabled(ctx context.Context, field string, value string) error {
	user, err := uc.userRepo.GetUserByField(ctx, field, value)
	if err != nil {
		return err
	}

	if user == nil {
		return nil
	}

	if user.Status == "enabled" {
		return fmt.Errorf("%s already exists", field)
	}

	return nil
}

func (uc *UserUseCase) RegisterUserWithOTP(ctx context.Context, accountID string, accountInfoType string, accountTypeID int) (*models.UserOTP, error) {
	user, err := uc.userRepo.GetUserByField(ctx, accountInfoType, accountID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		// initialise and create a new user
		user = uc.initialiseNewUser(accountInfoType, accountID, accountTypeID)
		err = uc.createUser(ctx, user)

		if err != nil {
			return nil, fmt.Errorf("unable to create user: %w", err)
		}
	}

	model, err := uc.otpRepo.GetOTPByUserIDAndMode(ctx, user.ID, registerOTPMode)
	if err != nil {
		return nil, fmt.Errorf("unable to get OTP by user ID and mode: %w", err)
	}

	otpTimeStamp := time.Now()

	model = &models.UserOTP{
		UserID:    user.ID,
		OTP:       uc.otpService.GenerateOTP(otpTimeStamp),
		Mode:      registerOTPMode,
		CreatedAt: otpTimeStamp,
	}

	// create OTP in database
	err = uc.otpRepo.CreateOTP(ctx, model)
	if err != nil {
		return nil, fmt.Errorf("unable to create OTP: %w", err)
	}

	return model, nil
}

func (uc *UserUseCase) initialiseNewUser(accountInfoType string, accountID string, accountType int) *models.User {
	user := &models.User{
		AccountType: models.AccountType{
			ID: accountType,
		},
		Status: statusPending,
	}

	if accountInfoType == "email" {
		user.Email = accountID
	} else {
		user.PhoneNumber = &accountID
	}

	return user
}

func (uc *UserUseCase) createUser(ctx context.Context, user *models.User) error {
	userId, err := uc.generateUserID(10)

	if err != nil {
		return fmt.Errorf("unable to generate user ID: %w", err)
	}

	user.UserID = userId

	err = uc.userRepo.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("unable to create user: %w", err)
	}

	return nil
}

func (uc *UserUseCase) UserByField(ctx context.Context, field string, value string) (*models.User, error) {
	user, err := uc.userRepo.GetUserByField(ctx, field, value)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) UserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := uc.userRepo.GetUserByField(ctx, "email", email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) ValidateOTP(ctx context.Context, userID int, otp string, mode string) error {
	userOTP, err := uc.otpRepo.GetOTPByUserIDAndMode(ctx, userID, mode)

	//return error if an error occurred or OTP does not exist
	if err != nil || userOTP == nil {
		return fmt.Errorf("unable to get OTP by user ID and mode: %w", err)
	}

	if userOTP.OTP != otp {
		return fmt.Errorf("incorrect user OTP: %w", err)
	}

	if err = uc.otpService.VerifyOTP(userOTP.OTP, userOTP.CreatedAt); err != nil {
		return fmt.Errorf("invalid user OTP: %w", err)
	}

	if err = uc.otpRepo.UseOTP(ctx, userOTP); err != nil {
		return fmt.Errorf("an error occurred while using OTP: %w", err)
	}

	return nil
}

func (uc *UserUseCase) UsersByField(ctx context.Context, field string, values []string) ([]models.User, error) {
	users, err := uc.userRepo.UsersByField(ctx, field, values)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *UserUseCase) generateUserID(size int) (string, error) {
	// First create a slice of bytes
	b := make([]byte, size)
	// Read size number of bytes into b
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// Encode our bytes as a base64 encoded string using URLEncoding
	encoded := base64.URLEncoding.EncodeToString(b)

	user, _ := uc.userRepo.GetUserByField(context.Background(), "user_id", encoded)

	if user != nil {
		return uc.generateUserID(size)
	}

	return encoded, nil
}
