package otp

import (
	"fmt"
	"time"
)

import "github.com/uaraven/gotp"

type OTPService struct {
	totp *gotp.TOTP
}

func NewOTPService(secret string) *OTPService {
	return &OTPService{
		totp: gotp.NewDefaultTOTP([]byte(secret)),
	}
}

func (os *OTPService) GenerateOTP(time time.Time) string {
	return os.totp.At(time)
}

func (os *OTPService) VerifyOTP(code string, time time.Time) error {
	if !os.totp.VerifyAt(code, time) {
		return fmt.Errorf("invalid OTP code")
	}

	return nil
}
