package celebrations

import (
	"celebut-api/internal/models"
	"celebut-api/pkg/postgres"
	"context"
	"fmt"
)

// CelebrationPostgresRepo -.
type CelebrationPostgresRepo struct {
	*postgres.Postgres
}

// NewCelebrationRepo -.
func NewCelebrationRepo(pg *postgres.Postgres) *CelebrationPostgresRepo {
	return &CelebrationPostgresRepo{pg}
}

// Create -.
func (r *CelebrationPostgresRepo) Create(ctx context.Context, c *models.Celebration) error {
	sql, args, err := r.Builder.
		Insert("celebrations").
		Columns("celebration_id, message").
		Values(c.CelebrationID, c.Message).
		ToSql()

	if err != nil {
		return fmt.Errorf("CelebrationPostgresRepo - Create - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("CelebrationPostgresRepo - Create - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *CelebrationPostgresRepo) Get(ctx context.Context, cID string) (*models.Celebration, error) {
	sql, _, err := r.Builder.
		Select("celebration_id, message").
		From("celebrations").
		Where("celebration_id = ?", cID).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("CelebrationPostgresRepo - Get - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("CelebrationPostgresRepo - Get - r.Pool.Query: %w", err)
	}

	c := models.Celebration{}

	err = row.Scan(&c.CelebrationID, &c.Message)
	if err != nil {
		return nil, fmt.Errorf("CelebrationPostgresRepo - Get - rows.Scan: %w", err)
	}

	return &c, nil

}

func (r *CelebrationPostgresRepo) Delete(ctx context.Context, c *models.Celebration) error {
	sql, args, err := r.Builder.
		Delete("celebrations").
		Where("celebration_id = ?", c.CelebrationID).
		ToSql()

	if err != nil {
		return fmt.Errorf("CelebrationPostgresRepo - Delete - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("CelebrationPostgresRepo - Delete - r.Pool.Exec: %w", err)
	}

	return nil

}
