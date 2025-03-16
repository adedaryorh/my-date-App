package dtos

import (
	"mime/multipart"
	"time"

	"backend.app/common/constants"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Celebration the celebration object
type Celebration struct {
	ID              uuid.UUID       `json:"id"`
	CelebrationType string          `json:"celebration_type,omitempty"`
	Audience        string          `json:"audience,omitempty"`
	Owner           string          `json:"owner,omitempty"`
	OwnerIdentifier string          `json:"owner_identifier,omitempty"`
	OwnerId         string          `json:"owner_id,omitempty"`
	Status          string          `json:"status,omitempty"`
	Frequency       string          `json:"frequency,omitempty"`
	CreatedBy       uuid.UUID       `json:"created_by,omitempty"`
	ExpiresAt       time.Time       `json:"expires_at,omitempty"`
	Long            float64         `json:"long,omitempty"`
	Lat             float64         `json:"lat,omitempty"`
	Notes           string          `json:"notes,omitempty"`
	SelectedFriends *pq.StringArray `json:"selected_friends,omitempty"`

	Caption         string     `json:"caption,omitempty"`
	CelebrationDate time.Time  `json:"celebration_date"`
	CelebrationKind string     `json:"celebration_kind,omitempty"`
	CreatedAt       time.Time  `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`

	Media []*Media `json:"media,omitempty"`
}

// Media the media object
type Media struct {
	ID        uuid.UUID  `json:"id,omitempty"`
	ObjectRef string     `json:"object_ref,omitempty"`
	ObjectId  uuid.UUID  `json:"object_id,omitempty"`
	MediaType string     `json:"media_type,omitempty"`
	MediaUrl  string     `json:"media_url,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// CelebrationsResponse the celebrations response object with pagination
type CelebrationsResponse struct {
	Celebrations []*Celebration `json:"celebrations"`
	PagingInfo   PagingInfo     `json:"paging_info"`
}

// CreateCelebrationDto the data object used to create a celebration
type CreateCelebrationDto struct {
	Audience        constants.CelebrationAudience  `form:"audience" validate:"required,is_enum"`
	CelebrationKind constants.CelebrationKind      `form:"celebration_kind" validate:"required,is_enum"`
	Owner           constants.CelebrationOwner     `form:"owner" validate:"required,is_enum"`
	OwnerId         *string                        `form:"owner_id" validate:"required_if=Owner friend,omitempty"`
	Frequency       constants.CelebrationFrequency `form:"frequency" validate:"required,is_enum"`
	CelebrationDate string                         `form:"celebration_date" validate:"omitempty,is_date"`
	ExpiresIn       *int                           `form:"expires_in" validate:"omitempty,numeric"`
	Longitude       float64                        `form:"longitude" validate:"required,numeric"`
	Latitude        float64                        `form:"latitude" validate:"required,numeric"`
	Notes           string                         `form:"notes" validate:"required,min=3,max=200"`
	Caption         string                         `form:"caption,omitempty" validate:"required,min=2,max=200"`
	SelectedFriends []string                       `form:"selected_friends" validate:"required_if=Audience selected-friends,omitempty,dive,is_uuid"`
	Media1          *multipart.FileHeader          `form:"media_1" validate:"omitempty,is_file"`
	Media2          *multipart.FileHeader          `form:"media_2" validate:"omitempty,is_file"`
	Media3          *multipart.FileHeader          `form:"media_3" validate:"omitempty,is_file"`
	Media4          *multipart.FileHeader          `form:"media_4" validate:"omitempty,is_file"`
	Media5          *multipart.FileHeader          `form:"media_5" validate:"omitempty,is_file"`
	Media6          *multipart.FileHeader          `form:"media_6" validate:"omitempty,is_file"`
}
