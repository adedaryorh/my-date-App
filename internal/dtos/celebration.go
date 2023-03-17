package dtos

import (
	"celebut-api/internal/models"
	"time"
)

type Celebration struct {
	PostID         string               `json:"post_id"`
	Message        string               `json:"message"`
	Author         UserInfo             `json:"author"`
	Media          []PostMedia          `json:"media"`
	Comments       []CelebrationComment `json:"comments"`
	UserReaction   []UserReaction       `json:"reactions"`
	CommentsCount  int                  `json:"comments_count"`
	ReactionsCount int                  `json:"reactions_count"`
	CreatedAt      time.Time            `json:"created_at"`
	UpdatedAt      time.Time            `json:"updated_at"`
}

type CelebrationComment struct {
}

type NewCelebration struct {
	Message *string
	Author  models.User
	Media   []string
}

type PagedCelebrations Paged[Celebration]
