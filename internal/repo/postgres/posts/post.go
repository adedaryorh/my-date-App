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
		Columns("user_id, post_id, message").
		Values(c.User.ID, c.PostID, c.Message).
		Suffix("RETURNING \"id\", \"created_at\", \"updated_at\"").
		ToSql()

	if err != nil {
		return fmt.Errorf("PostPostgresRepo - Create - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return fmt.Errorf("PostPostgresRepo - Create - r.Pool.Scan: %w", err)
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

func (r *PostPostgresRepo) GetUserPosts(ctx context.Context, userID int, offset int, limit int) ([]models.Post, error) {
	sql, args, err := r.Builder.
		Select("id, post_id, message, created_at, updated_at").
		From("posts").
		Where("user_id = ?", userID).
		Offset(uint64(offset)).
		Limit(uint64(limit)).
		OrderBy("created_at DESC").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("PostPostgresRepo - GetUserPosts - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("PostPostgresRepo - GetUserPosts - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	posts := make([]models.Post, 0)
	for rows.Next() {
		p := models.Post{}

		err = rows.Scan(&p.ID, &p.PostID, &p.Message, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("PostPostgresRepo - GetUserPosts - rows.Scan: %w", err)
		}

		posts = append(posts, p)
	}

	return posts, nil
}

func (r *PostPostgresRepo) Delete(ctx context.Context, userID int, postID string) error {
	sql, args, err := r.Builder.
		Delete("posts").
		Where("post_id = ? AND user_id = ?", postID, userID).
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
