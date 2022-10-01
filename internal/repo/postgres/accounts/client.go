package accounts

import (
	"celebut-api/internal/models"
	"celebut-api/pkg/postgres"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

const _defaultEntityCap = 64

// ClientPostgresRepo -.
type ClientPostgresRepo struct {
	*postgres.Postgres
}

// NewClientRepo -.
func NewClientRepo(pg *postgres.Postgres) *ClientPostgresRepo {
	return &ClientPostgresRepo{pg}
}

// CreateClient -.
func (r *ClientPostgresRepo) CreateClient(ctx context.Context, c *models.Client) error {
	sql, args, err := r.Builder.
		Insert("clients").
		Columns("name, client_id, secret").
		Values(c.Name, c.ClientID, c.Secret).
		ToSql()

	if err != nil {
		return fmt.Errorf("ClientPostgresRepo - CreateClient - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ClientPostgresRepo - CreateClient - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetClient -.
func (r *ClientPostgresRepo) GetClient(ctx context.Context, clientId string) (*models.Client, error) {
	sql, _, err := r.Builder.
		Select("name, client_id, secret").
		From("clients").
		Where(squirrel.Eq{"client_id": clientId}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("ClientPostgresRepo - AuthenticateClient - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, clientId)
	c := models.Client{}

	err = row.Scan(&c.Name, &c.ClientID, &c.Secret)
	if err != nil {
		return nil, fmt.Errorf("ClientPostgresRepo - AuthenticateClient - row.Scan: %w", err)
	}

	return &c, nil
}

// GetClientToken -.
func (r *ClientPostgresRepo) GetClientToken(ctx context.Context, token string) (*models.ClientToken, error) {
	sql, args, err := r.Builder.
		Select("token, expiry").
		From("client_tokens").
		Where("token = ?", token).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("ClientPostgresRepo - GetClientToken - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("ClientPostgresRepo - GetClientToken - r.Pool.Query: %w", err)
	}

	ct := models.ClientToken{}

	err = row.Scan(&ct.Token, &ct.Expiry)
	if err != nil {
		return nil, fmt.Errorf("ClientPostgresRepo - GetClientToken - rows.Scan: %w", err)
	}

	return &ct, nil
}

// CreateToken -.
func (r *ClientPostgresRepo) CreateToken(ctx context.Context, ct *models.ClientToken) error {
	sql, args, err := r.Builder.
		Insert("client_tokens").
		Columns("token, expiry").
		Values(ct.Token, ct.Expiry).
		ToSql()

	if err != nil {
		return fmt.Errorf("ClientPostgresRepo - CreateToken - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ClientPostgresRepo - CreateClient - r.Pool.Exec: %w", err)
	}

	return nil
}

// UpdateClientToken -.
func (r *ClientPostgresRepo) UpdateClientToken(ctx context.Context, ct *models.ClientToken) error {
	sql, args, err := r.Builder.
		Update("client_tokens").
		Set("expiry", ct.Expiry).
		Where(squirrel.Eq{"token": ct.Token}).
		ToSql()

	if err != nil {
		return fmt.Errorf("ClientPostgresRepo - UpdateClientToken - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ClientPostgresRepo - UpdateClientToken - r.Pool.Exec: %w", err)
	}

	return nil
}
