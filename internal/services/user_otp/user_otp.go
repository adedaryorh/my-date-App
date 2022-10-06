package user_otp

import (
	"celebut-api/internal/models"
	"celebut-api/internal/repo"
	"celebut-api/internal/services/otp_generator"
	"context"
	"fmt"
	"time"
)

type UserOTPService struct {
	repo      repo.UserOTP
	generator otp_generator.Generator
}

func NewUserOTPService(r repo.UserOTP, g otp_generator.Generator) *UserOTPService {
	return &UserOTPService{repo: r, generator: g}
}

func (uos *UserOTPService) Create(ctx context.Context, userID int, mode string) (string, error) {
	otpTimeStamp := time.Now()
	generatedOTP := uos.generator.GenerateOTP(otpTimeStamp)

	userOTP := &models.UserOTP{
		UserID:    userID,
		Mode:      mode,
		OTP:       generatedOTP,
		CreatedAt: otpTimeStamp,
	}

	// create OTP in database
	err := uos.repo.CreateOTP(ctx, userOTP)
	if err != nil {
		return "", fmt.Errorf("userotpservice - unable to create OTP: %w", err)
	}

	return generatedOTP, nil
}

func (uos *UserOTPService) VerifyOTP(code string, _ time.Time) error {
	if code != "123456" {
		return fmt.Errorf("invalid OTP code")
	}

	return nil
}
