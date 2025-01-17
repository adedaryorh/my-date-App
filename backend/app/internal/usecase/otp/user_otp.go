package otp

import (
	"context"
	"fmt"

	"backend.app/internal/services/mailer"
	"backend.app/internal/services/user_otp"
)

// UserOTPUseCase -.
type UserOTPUseCase struct {
	otpService user_otp.UserOTP
	mailer     mailer.Mailer
}

// NewUserOTPUseCase -.
func NewUserOTPUseCase(o user_otp.UserOTP, m mailer.Mailer) *UserOTPUseCase {
	return &UserOTPUseCase{
		otpService: o,
		mailer:     m,
	}
}

// Create - Create a new User OTP
func (uc *UserOTPUseCase) Create(ctx context.Context, userID int, mode string, recipient Recipient) error {
	otp, err := uc.otpService.Create(ctx, userID, mode)

	if err != nil {
		return fmt.Errorf("UserOTPUseCase - UserOTPs - uc.repo.Create: %w", err)
	}

	if recipient.Type == "email" {
		err = uc.mailer.SendOTP(recipient.Info, otp)
	}

	return nil
}
