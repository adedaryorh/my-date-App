package core

import (
	"context"
	"errors"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/common/messages"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
)

// CreateWallet creates a new wallet
func (c *Core) CreateWallet(ctx context.Context, user *models.User, currency constants.Currency) error {
	// ensure user has no wallet of similar currency created earlier
	existingWallet, err := c.repo.GetWalletByField(ctx, helpers.Map{"owner_id": user.ID, "w.currency": currency})
	if err != nil && err != messages.ErrWalletNotFound {
		return err
	}

	if existingWallet != nil {
		return errors.New("user already has wallet for this currency")
	}

	// create balance
	balanceId, err := c.CreateBalance(ctx, currency)
	if err != nil {
		return err
	}
	// create wallet proper
	wallet := models.Wallet{
		BalanceID: *balanceId,
		Currency:  string(currency),
		Status:    string(constants.StatusInActive),
		OwnerType: user.AccountType,
		OwnerId:   user.ID,
	}
	return c.repo.CreateWallet(ctx, &wallet)
}

// GetUserWallet method to fetch user's wallet with balance
func (c *Core) GetUserWallet(ctx context.Context, user *models.User) *dtos.ResponseObject {
	result, err := c.repo.GetWalletByField(ctx, helpers.Map{"owner_id": user.ID})
	if err != nil {
		return ServerErrorResponse(err)
	}
	return SuccessResponse("user wallet successfully fetched", result)
}
