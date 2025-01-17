package otp_generator

import (
	"fmt"
	"time"
)

import "github.com/uaraven/gotp"

type OTPGeneratorService struct {
	totp *gotp.TOTP
}

func NewOTPGeneratorService(secret string) *OTPGeneratorService {
	return &OTPGeneratorService{
		totp: gotp.NewDefaultTOTP([]byte(secret)),
	}
}

func (os *OTPGeneratorService) GenerateOTP(time time.Time) string {
	return os.totp.At(time)
}

func (os *OTPGeneratorService) VerifyOTP(code string, time time.Time) error {
	if !os.totp.VerifyAt(code, time) {
		return fmt.Errorf("invalid OTP code")
	}

	return nil
}
