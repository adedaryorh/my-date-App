package dtos

import (
	"github.com/google/uuid"
	"time"
)

type NearbyCelebration struct {
	ID              uuid.UUID `json:"id"`
	Caption         string    `json:"caption"`
	Notes           string    `json:"notes"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	DistanceKM      float64   `json:"distance_km"`
	CelebrationDate time.Time `json:"celebration_date"`
	CreatedAt       time.Time `json:"created_at"`
}

type ModerationQueueItem struct {
	ID             uuid.UUID  `json:"id"`
	CelebrationID  uuid.UUID  `json:"celebration_id"`
	UserID         uuid.UUID  `json:"user_id"`
	ToxicityScore  float64    `json:"toxicity_score"`
	FlaggedContent string     `json:"flagged_content"`
	Status         string     `json:"status"`
	ModeratorID    *uuid.UUID `json:"moderated_by,omitempty"`
	ModeratedAt    *time.Time `json:"moderated_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
}

type ResolveModerationRequest struct {
	Decision string `json:"decision" validate:"required,oneof=approved rejected"`
}
