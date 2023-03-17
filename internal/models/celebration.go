package models

import "time"

type Celebration struct {
	ID             int
	FlaggedCounter int
	UserID         int
	CelebrationID  string
	User           User
	Message        *string
	Comments       []CelebrationComment
	Media          []CelebrationMedia
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CelebrationComment struct {
	ID             int
	CommentID      string
	FlaggedCounter int
	UserID         int
	CelebrationID  int
	Celebration    Celebration
	User           User
	Text           *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CelebrationMedia struct {
	ID          int
	Celebration Celebration
	Source      string
	Extension   string
	Type        string
	CreatedAt   time.Time
}
