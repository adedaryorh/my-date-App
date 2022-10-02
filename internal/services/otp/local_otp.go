package otp

import (
	"fmt"
	"time"
)

type LocalOTPService struct {
}

func NewLocalOTPService() *LocalOTPService {
	return &LocalOTPService{}
}

func (os *LocalOTPService) GenerateOTP(_ time.Time) string {
	return "123456"
}

func (os *LocalOTPService) VerifyOTP(code string, _ time.Time) error {
	if code != "123456" {
		return fmt.Errorf("invalid OTP code")
	}

	return nil
}
