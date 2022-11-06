package posts

import (
	"celebut-api/internal/models"
	"celebut-api/pkg/postgres"
	"context"
	"fmt"
)

// PostPostgresRepo -.
type PostPostgresRepo struct {
	*postgres.Postgres
}

// NewPostsRepo -.
func NewPostsRepo(pg *postgres.Postgres) *PostPostgresRepo {
	return &PostPostgresRepo{pg}
}

// Create -.
func (r *PostPostgresRepo) Create(ctx context.Context, c *models.Post) error {
	sql, args, err := r.Builder.
		Insert("posts").
		Columns("post_id, message").
		Values(c.PostID, c.Message).
		ToSql()

	if err != nil {
		return fmt.Errorf("PostPostgresRepo - Create - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PostPostgresRepo - Create - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *PostPostgresRepo) Get(ctx context.Context, cID string) (*models.Post, error) {
	sql, _, err := r.Builder.
		Select("post_id, message").
		From("posts").
		Where("post_id = ?", cID).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("PostPostgresRepo - Get - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("PostPostgresRepo - Get - r.Pool.Query: %w", err)
	}

	c := models.Post{}

	err = row.Scan(&c.PostID, &c.Message)
	if err != nil {
		return nil, fmt.Errorf("PostPostgresRepo - Get - rows.Scan: %w", err)
	}

	return &c, nil

}

func (r *PostPostgresRepo) Delete(ctx context.Context, c *models.Post) error {
	sql, args, err := r.Builder.
		Delete("posts").
		Where("post_id = ?", c.PostID).
		ToSql()

	if err != nil {
		return fmt.Errorf("PostPostgresRepo - Delete - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PostPostgresRepo - Delete - r.Pool.Exec: %w", err)
	}

	return nil

}
