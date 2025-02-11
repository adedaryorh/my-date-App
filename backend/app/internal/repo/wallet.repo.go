package repo

import (
	"context"
	"errors"
	"fmt"

	"backend.app/common/messages"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

// CreateWallet  creates a new wallet record
func (r *Repo) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	sql, args, err := r.postgres.Builder.
		Insert(WalletsTable).
		Columns("balance_id, currency, status, owner_type, owner_id").
		Values(wallet.BalanceID, wallet.Currency, wallet.Status, wallet.OwnerType, wallet.OwnerId).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return fmt.Errorf("PostgresRepo - CreateWallet - r.Builder: %w", err)
	}

	row := r.postgres.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&wallet.ID)
	if err != nil {
		return fmt.Errorf("PostgresRepo - CreateWallet - r.Pool.Scan: %w", err)
	}

	return nil
}

// GetWalletByField  Get a single wallet based of the sent filter
func (r *Repo) GetWalletByField(ctx context.Context, filter map[string]interface{}) (*dtos.Wallet, error) {
	sql, args, err := r.postgres.Builder.
		Select("w.id,w.balance_id,w.currency,w.status,w.owner_type,w.owner_id,b.available_balance,b.change_amount,b.locked_amount").
		From("wallets as w").
		InnerJoin("balances as b ON w.balance_id = b.id").
		Where(squirrel.Eq(filter)).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("unable to build query: %w", err)
	}

	row := r.postgres.Pool.QueryRow(ctx, sql, args...)

	wallet := dtos.Wallet{}
	//var currency string
	err = row.Scan(
		&wallet.ID,
		&wallet.BalanceID,
		&wallet.Currency,
		&wallet.Status,
		&wallet.OwnerType,
		&wallet.OwnerId,
		&wallet.AvailableBalance,
		&wallet.ChangeAmount,
		&wallet.LockedAmount,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, messages.ErrWalletNotFound
		}
		r.log.Debug("GetWalletByField : row.Scan: %w", err)
		return nil, errors.New("something went wrong, try again later")
	}

	return &wallet, nil
}
