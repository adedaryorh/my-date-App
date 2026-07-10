package models

import (
	"time"

	"github.com/google/uuid"
)

// Wallet The wallet Object model
type Wallet struct {
	ID        uuid.UUID `json:"id"`
	BalanceID uuid.UUID `json:"balance_id"`
	Currency  string    `json:"currency"`
	Status    string    `json:"status"`
	OwnerType string    `json:"owner_type"`
	OwnerId   uuid.UUID `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
