package posts

import (
	"context"
	"fmt"
	"strconv"

	"backend.app/internal/models"
	"backend.app/pkg/postgres"
	"github.com/Masterminds/squirrel"
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
func (r *PostPostgresRepo) Create(ctx context.Context, p *models.Post) error {
	sql, args, err := r.Builder.
		Insert("posts").
		Columns("user_id, post_id, message, parent_id").
		Values(p.User.ID, p.PostID, p.Message, p.ParentID).
		Suffix("RETURNING \"id\", \"created_at\", \"updated_at\"").
		ToSql()

	if err != nil {
		return fmt.Errorf("PostPostgresRepo - Create - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("PostPostgresRepo - Create - r.Pool.Scan: %w", err)
	}

	return nil
}

// Update -.
func (r *PostPostgresRepo) Update(ctx context.Context, p *models.Post) error {
	sql, args, err := r.Builder.
		Update("posts").
		SetMap(squirrel.Eq{
			"flagged_counter": p.FlaggedCounter,
		}).
		Where(squirrel.Eq{"id": p.ID}).
		ToSql()

	if err != nil {
		return fmt.Errorf("PostPostgresRepo - Update - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PostPostgresRepo - Update - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *PostPostgresRepo) Get(ctx context.Context, postID string) (*models.Post, error) {
	sql, args, err := r.Builder.
		Select("id, user_id, post_id, message, parent_id, flagged_counter").
		From("posts").
		Where("post_id = ?", postID).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("PostPostgresRepo - Get - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("PostPostgresRepo - Get - r.Pool.Query: %w", err)
	}

	p := models.Post{}

	err = row.Scan(&p.ID, &p.UserID, &p.PostID, &p.Message, &p.ParentID, &p.FlaggedCounter)
	if err != nil {
		return nil, fmt.Errorf("PostPostgresRepo - Get - rows.Scan: %w", err)
	}

	return &p, nil
}

func (r *PostPostgresRepo) GetUserPosts(ctx context.Context, userID int, offset *int, limit *int) ([]models.Post, error) {
	builder := r.Builder.
		Select("id, user_id, post_id, message, created_at, updated_at").
		From("posts").
		Where("user_id = ? AND parent_id IS NULL", userID).
		OrderBy("created_at DESC")

	if offset != nil {
		builder = builder.Offset(uint64(*offset))
	}

	if limit != nil {
		builder = builder.Limit(uint64(*limit))
	}

	sql, args, err := builder.ToSql()
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

		err = rows.Scan(&p.ID, &p.UserID, &p.PostID, &p.Message, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("PostPostgresRepo - GetUserPosts - rows.Scan: %w", err)
		}

		posts = append(posts, p)
	}

	return posts, nil
}

func (r *PostPostgresRepo) GetMultiUsersPosts(ctx context.Context, userIDs []int, offset *int, limit *int) ([]models.Post, error) {
	userIDsToStr := make([]string, len(userIDs))
	for i, v := range userIDs {
		userIDsToStr[i] = strconv.Itoa(v)
	}

	builder := r.Builder.
		Select("id, user_id, post_id, message, created_at, updated_at").
		From("posts").
		Where(
			squirrel.Eq{
				"user_id":   userIDsToStr,
				"parent_id": nil,
			},
		).
		OrderBy("created_at DESC")

	if offset != nil {
		builder = builder.Offset(uint64(*offset))
	}

	if limit != nil {
		builder = builder.Limit(uint64(*limit))
	}

	sql, args, err := builder.ToSql()
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

		err = rows.Scan(&p.ID, &p.UserID, &p.PostID, &p.Message, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("PostPostgresRepo - GetUserPosts - rows.Scan: %w", err)
		}

		posts = append(posts, p)
	}

	return posts, nil
}

func (r *PostPostgresRepo) GetPostComments(ctx context.Context, postID int, offset *int, limit *int) ([]models.Post, error) {
	builder := r.Builder.
		Select("id, user_id, post_id, message, created_at, updated_at").
		From("posts").
		Where("parent_id = ?", postID).
		OrderBy("created_at ASC")

	if offset != nil {
		builder = builder.Offset(uint64(*offset))
	}

	if limit != nil {
		builder = builder.Limit(uint64(*limit))
	}

	sql, args, err := builder.ToSql()

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

		err = rows.Scan(&p.ID, &p.UserID, &p.PostID, &p.Message, &p.CreatedAt, &p.UpdatedAt)
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
