package models

import "time"

type Celebration struct {
	ID            int
	CelebrationID string
	User          User
	Message       *string
	Media         []CelebrationMedia
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CelebrationMedia struct {
	ID          string
	Celebration Celebration
	Source      string
	Type        string
	CreatedAt   time.Time
}
