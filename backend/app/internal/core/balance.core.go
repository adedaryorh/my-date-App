package core

import (
	"context"
	"errors"
	"fmt"
	"time"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/internal/models"
	"github.com/google/uuid"
)

func (c *Core) CreateBalance(ctx context.Context, currency constants.Currency) (*uuid.UUID, error) {

	balance := models.Balance{
		ID:               uuid.New(),
		AvailableBalance: helpers.Int64ToPointer(0),
		ChangeAmount:     0,
		LockedAmount:     helpers.Int64ToPointer(0),
		Currency:         string(currency),
		Mode:             string(constants.TransactionModeCredit),
		Status:           string(constants.StatusActive),
		HashKey:          helpers.GenerateRandomUppercase(32),
		TxnTime:          time.Now(),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	newBalance, err := c.CreateBalanceWithHistory(ctx, &balance)
	if err != nil {
		return nil, err
	}

	id := newBalance.ID
	return &id, nil
}

func (c *Core) CreateBalanceWithHistory(ctx context.Context, balance *models.Balance) (*models.Balance, error) {
	// hash wallet
	_, hashedWallet := helpers.GetEntityComputedHash(balance)

	balance.Hash = hashedWallet
	balance.Sequence = 1

	// save to db
	return c.repo.CreateBalanceWithHistory(ctx, *balance)

}

func (c *Core) FindAndVerifyBalance(ctx context.Context, balanceId uuid.UUID, limit int64) (*models.Balance, error) {
	balance, err := c.repo.GetBalanceByField(ctx, helpers.Map{"id": balanceId})
	if err != nil {
		return nil, err
	}

	verified, err := c.verifyWallet(ctx, balance)
	if err != nil {
		return nil, err
	}
	if !verified {
		return nil, errors.New("invalid wallet")
	}

	return balance, nil
}

func (c *Core) verifyWallet(ctx context.Context, balance *models.Balance) (bool, error) {
	var hasHistory bool
	_, hashedWallet := helpers.GetEntityComputedHash(balance)
	fmt.Println(balance.Hash, hashedWallet, "######### BALANCE HASH")
	if balance.Hash != hashedWallet {
		return false, errors.New("invalid wallet")
	}
	filter := map[string]interface{}{
		"balance_id": balance.ID,
	}
	sort := "sequence:desc"

	balanceHistory, err := c.repo.GetBalanceHistory(ctx, filter, sort, 10)
	var balanceHistoryFlip []*models.BalanceHistory
	for i := len(balanceHistory) - 1; i >= 0; i-- {
		balanceHistoryFlip = append(balanceHistoryFlip, balanceHistory[i])
	}
	if err != nil {
		return false, err
	}
	if len(balanceHistory) > 0 {
		hasHistory = true
	}
	var prevHash = ""
	for _, history := range balanceHistoryFlip {
		_, walletHistoryHash := helpers.GetEntityComputedHash(history)
		fmt.Println(history.Sequence)
		if prevHash == "" {
			prevHash = history.PreviousHash
		}
		fmt.Println(history.Hash, walletHistoryHash, history.Sequence, "######### HISTORY HASH")
		if history.Hash != walletHistoryHash {
			return false, nil
		}

		if prevHash != history.PreviousHash {
			return false, nil
		}
		prevHash = history.Hash
	}

	if hashedWallet != "" {
		return prevHash == hashedWallet, nil
	}

	return hasHistory, nil
}

func (c *Core) GetBalanceByID(ctx context.Context, balanceId uuid.UUID) (*models.Balance, error) {
	return c.repo.GetBalanceByField(ctx, helpers.Map{"id": balanceId})
}
