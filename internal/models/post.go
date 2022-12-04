package models

import "time"

type Post struct {
	ID        int
	ParentID  *int
	PostID    string
	User      User
	Message   *string
	Media     []PostMedia
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostMedia struct {
	ID        int
	Post      Post
	Source    string
	Extension string
	Type      string
	CreatedAt time.Time
}
