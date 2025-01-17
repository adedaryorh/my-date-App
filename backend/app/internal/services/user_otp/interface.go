package user_otp

import "context"

type UserOTP interface {
	Create(ctx context.Context, userID int, mode string) (string, error)
}
