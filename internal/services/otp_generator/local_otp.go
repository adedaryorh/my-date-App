package otp_generator

import (
	"fmt"
	"time"
)

type LocalOTPGeneratorService struct {
}

func NewLocalOTPGeneratorService() *LocalOTPGeneratorService {
	return &LocalOTPGeneratorService{}
}

func (os *LocalOTPGeneratorService) GenerateOTP(_ time.Time) string {
	return "123456"
}

func (os *LocalOTPGeneratorService) VerifyOTP(code string, _ time.Time) error {
	if code != "123456" {
		return fmt.Errorf("invalid OTP code")
	}

	return nil
}
