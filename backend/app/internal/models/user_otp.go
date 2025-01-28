package models

import "time"

type UserOTP struct {
	ID        int
	UserID    int
	OTP       string
	Counter   int64
	Mode      string
	Used      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
