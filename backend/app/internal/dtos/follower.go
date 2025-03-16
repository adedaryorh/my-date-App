package dtos

import (
	"time"

	"github.com/google/uuid"
)

type Follower struct {
	ID           uuid.UUID `json:"id" gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	UserId       uuid.UUID `json:"user_id,omitempty"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Username     string    `json:"username,omitempty"`
	ProfileImage string    `json:"profile_image,omitempty"`
	Following    bool      `json:"following,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

type Blocked struct {
	ID            uuid.UUID `json:"id" gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	BlockedUserId uuid.UUID `json:"blocked_user_id,omitempty"`
	FirstName     string    `json:"first_name,omitempty"`
	LastName      string    `json:"last_name,omitempty"`
	Username      string    `json:"username,omitempty"`
	ProfileImage  string    `json:"profile_image,omitempty"`
	BlockedAt     time.Time `json:"blocked_at,omitempty"`
}

// FollowersResponse the followers response object with pagination
type FollowersResponse struct {
	Followers  []*Follower `json:"followers"`
	PagingInfo PagingInfo  `json:"paging_info"`
}

// BlockedResponse the blocked response object with pagination
type BlockedResponse struct {
	Blocked    []*Blocked `json:"blocked"`
	PagingInfo PagingInfo `json:"paging_info"`
}
