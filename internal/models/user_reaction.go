package models

import "time"

type UserReaction struct {
	ID        int
	UserID    int
	PostID    string
	Reaction  string
	CreatedAt time.Time
	UpdatedAt time.Time

	User User
	Post Post
}
