package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"

	"backend.app/common/messages"
	"backend.app/internal/models"
)

func (r *Repo) GetBalanceByField(ctx context.Context, filter map[string]interface{}) (*models.Balance, error) {
	var balance models.Balance
	sql, args, err := r.postgres.Builder.
		Select("id,available_balance, currency, status, change_amount, locked_amount,mode,hash ,sequence, previous_hash, hash_key, tnx_time").
		From("balances").
		Where(squirrel.Eq(filter)).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("unable to build query: %w", err)
	}

	row := r.postgres.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&balance)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, messages.ErrBalanceNotFound
		}
		return nil, fmt.Errorf("row.Scan: %w", err)
	}

	return &balance, nil
}

func (r *Repo) GetBalanceHistory(ctx context.Context, filter map[string]interface{}, sort string, limit int) ([]*models.BalanceHistory, error) {
	var field, order string
	builder := r.postgres.Builder.
		Select("id, balance_id,operation,available_balance, currency, status, change_amount, locked_amount,mode,hash ,previous_hash, hash_key,sequence, operation,transaction_id,tnx_time, created_at,updated_at").
		From("balance_history").
		Where(squirrel.Eq(filter))

	// created_at:asc
	if sort != "" {
		splittedSort := strings.Split(sort, ":")
		if len(splittedSort) < 2 {
			field = "created_at"
			order = "asc"
		} else {
			field = splittedSort[0]
			order = splittedSort[1]
		}

		builder = builder.OrderBy(fmt.Sprintf("balance_history.%s %s", field, order))
	}

	if limit > 0 {
		builder = builder.Limit(uint64(limit))
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("PostgresRepo - GetBalanceHistory - r.Builder: %w", err)
	}

	rows, err := r.postgres.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("PostgresRepo - GetBalanceHistory - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	balanceHistory := make([]*models.BalanceHistory, 0)
	for rows.Next() {
		history := models.BalanceHistory{}
		err = rows.Scan(
			&history.ID,
			&history.BalanceId,
			&history.AvailableBalance,
			&history.Currency,
			&history.Status,
			&history.ChangeAmount,
			&history.LockedAmount,
			&history.Mode,
			&history.Hash,
			&history.PreviousHash,
			&history.HashKey,
			&history.Sequence,
			&history.Operation,
			&history.TransactionId,
			&history.TxnTime,
			&history.CreatedAt,
			&history.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("PostPostgresRepo - GetUserPosts - rows.Scan: %w", err)
		}

		balanceHistory = append(balanceHistory, &history)
	}

	return balanceHistory, nil
}

func (r *Repo) CreateBalanceWithHistory(ctx context.Context, balance models.Balance) (*models.Balance, error) {
	// start session
	tx, err := r.postgres.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()
	// create balance
	sql, args, err := r.postgres.Builder.
		Insert("balances").
		Columns("id,available_balance, currency, status, change_amount, locked_amount,mode,hash ,sequence, hash_key, tnx_time").
		Values(balance.ID, balance.AvailableBalance, balance.Currency, balance.Status, balance.ChangeAmount, balance.LockedAmount, balance.Mode, balance.Hash, balance.Sequence, balance.HashKey, balance.TxnTime).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		tx.Rollback(ctx)
		return nil, fmt.Errorf("BalancePostgresRepo - CreateBalance - r.Builder: %w", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	// create balance history in same transaction
	historySql, historyArgs, err := r.postgres.Builder.
		Insert("balance_history").
		Columns("balance_id,operation,available_balance, currency, status, change_amount, locked_amount,mode,hash ,sequence, hash_key, tnx_time").
		Values(balance.ID, "insert", balance.AvailableBalance, balance.Currency, balance.Status, balance.ChangeAmount, balance.LockedAmount, balance.Mode, balance.Hash, balance.Sequence, balance.HashKey, balance.TxnTime).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		tx.Rollback(ctx)
		return nil, fmt.Errorf("BalancePostgresRepo - CreateBalanceHistory - r.Builder: %w", err)
	}

	_, err = tx.Exec(ctx, historySql, historyArgs...)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	return &balance, nil
}

// func (r *Repo) UpdateBalanceLedger(ctx context.Context, balance *models.Balance, balanceHistory *models.BalanceHistory) (*models.BalanceTxnResponse, error) {
// 	// start session
// 	tx := r.Repo.PostgresDb.Begin()

// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	if err := tx.Error; err != nil {
// 		return nil, err
// 	}
// 	// :TODO implement transaction locking instead
// 	// lock wallet
// 	//tx.Exec("LOCK TABLE wallets IN ACCESS EXCLUSIVE MODE")

// 	// update wallet
// 	balanceResponse := &models.BalanceTxnResponse{}

// 	if err := tx.WithContext(ctx).Model(&models.Balance{}).Where("id = ?", balance.Id).UpdateColumns(balance).Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	// insert wallet history
// 	if err := tx.WithContext(ctx).Model(&models.BalanceHistory{}).Create(balanceHistory).Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	if balance.Mode == models.TransactionMode.Credit {
// 		balanceResponse.CreditTransactionId = *balance.TransactionId
// 		//balanceResponse.CreditWalletHistoryId = balanceHistory.Id
// 	}

// 	return balanceResponse, tx.Commit().Error
// }
