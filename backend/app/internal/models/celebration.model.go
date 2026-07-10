package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Celebration the celebration model object
type Celebration struct {
	ID              uuid.UUID       `json:"id" gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	Audience        string          `json:"audience,omitempty"`
	CelebrationType string          `json:"celebration_type,omitempty"`
	Owner           string          `json:"owner,omitempty"`
	OwnerIdentifier string          `json:"owner_identifier,omitempty"`
	OwnerId         string          `json:"owner_id,omitempty"`
	Status          string          `json:"status,omitempty"`
	Frequency       string          `json:"frequency,omitempty"`
	CreatedBy       uuid.UUID       `json:"created_by,omitempty"`
	ExpiresAt       time.Time       `json:"expires_at,omitempty"`
	Long            float64         `json:"longitude,omitempty"`
	Lat             float64         `json:"latitude,omitempty"`
	Notes           string          `json:"notes,omitempty"`
	SelectedFriends *pq.StringArray `json:"selected_friends,omitempty"`

	Caption         string `json:"caption,omitempty"`
	CelebrationDate time.Time
	CelebrationKind string

	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Media struct {
	ID        uuid.UUID  `json:"id" gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	ObjectRef string     `json:"object_ref,omitempty"`
	ObjectId  uuid.UUID  `json:"object_id,omitempty"`
	OwnerId   string     `json:"owner_id,omitempty"`
	MediaType string     `json:"media_type,omitempty"`
	MediaUrl  string     `json:"media_url,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
