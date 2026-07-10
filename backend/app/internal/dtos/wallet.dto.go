package dtos

import "github.com/google/uuid"

type Wallet struct {
	ID               uuid.UUID `json:"id"`
	BalanceID        uuid.UUID `json:"balance_id"`
	Currency         string    `json:"currency"`
	Status           string    `json:"status"`
	OwnerType        string    `json:"owner_type"`
	OwnerId          uuid.UUID `json:"owner_id"`
	AvailableBalance *int64    `json:"available_balance"`
	ChangeAmount     int64     `json:"change_amount"`
	LockedAmount     *int64    `json:"locked_amount"`
}
