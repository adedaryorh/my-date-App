package otp_generator

import "time"

type Generator interface {
	GenerateOTP(time time.Time) string
	VerifyOTP(code string, time time.Time) error
}
