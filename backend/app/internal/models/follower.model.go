package models

import (
	"time"

	"github.com/google/uuid"
)

type Follower struct {
	ID         uuid.UUID  `json:"id" gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	UserId     uuid.UUID  `json:"user_id,omitempty"`
	FollowerId uuid.UUID  `json:"follower_id,omitempty"`
	CreatedAt  time.Time  `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

type Blocked struct {
	ID            uuid.UUID  `json:"id" gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	UserId        uuid.UUID  `json:"user_id,omitempty"`
	BlockedUserId uuid.UUID  `json:"blocked_user_id,omitempty"`
	CreatedAt     time.Time  `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}
