package otp

import (
	"context"
)

type UserOTP interface {
	Create(ctx context.Context, userID int, mode string, recipient Recipient) error
}

type Recipient struct {
	Info string
	Type string
}
