package dtos

import (
	"time"

	"backend.app/internal/models"
)

type Post struct {
	PostID         string         `json:"post_id"`
	Message        string         `json:"message"`
	Author         UserInfo       `json:"author"`
	Media          []PostMedia    `json:"media"`
	Comments       []Post         `json:"comments"`
	UserReaction   []UserReaction `json:"reactions"`
	CommentsCount  int            `json:"comments_count"`
	ReactionsCount int            `json:"reactions_count"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

type PagedPosts struct {
	Page  *int   `json:"page"`
	Limit *int   `json:"limit"`
	Posts []Post `json:"posts"`
}

type PostMedia struct {
	URL           string `json:"url"`
	AlternateText string `json:"alternate_text"`
}

type NewPost struct {
	Message *string
	Author  models.User
	Media   []string
}
