package models

import "time"

type Post struct {
	ID        int
	PostID    string
	User      User
	Message   *string
	Media     []PostMedia
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostMedia struct {
	ID        string
	Post      Post
	Source    string
	Type      string
	CreatedAt time.Time
}
