package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	// Balance this is the balance's model object
	Balance struct {
		ID               uuid.UUID  `json:"id"`
		AvailableBalance *int64     `json:"available_balance"`
		Currency         string     `json:"currency"`
		Status           string     `json:"status"`
		ChangeAmount     int64      `json:"change_amount"`
		LockedAmount     *int64     `json:"locked_amount"`
		Mode             string     `json:"mode"`
		TransactionId    *uuid.UUID `json:"transaction_id"`
		Hash             string     `json:"-"`
		Sequence         int64      `json:"-"`
		PreviousHash     string     `json:"-"`
		HashKey          string     `json:"hash_key"`
		TxnTime          time.Time  `json:"txn_time"`
		CreatedAt        time.Time  `json:"created_at"`
		UpdatedAt        time.Time  `json:"updated_at"`
	}

	// Lien this is the lien's model object
	Lien struct {
		Id            uuid.UUID `json:"id"`
		BalanceId     uuid.UUID
		TransactionId uuid.UUID
		LienAmount    int64
		Currency      string
		Status        string
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}

	// BalanceHistory this is the balances' history model object
	BalanceHistory struct {
		Balance
		BalanceId uuid.UUID
		Operation string
	}

	BalanceTxnResponse struct {
		DebitTransactionId  uuid.UUID
		CreditTransactionId uuid.UUID
	}
)
