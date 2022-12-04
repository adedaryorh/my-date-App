package posts

import (
	"celebut-api/internal/models"
	"celebut-api/pkg/postgres"
	"context"
	"fmt"
)

// PostMediaPostgresRepo -.
type PostMediaPostgresRepo struct {
	*postgres.Postgres
}

// NewPostMediaRepo -.
func NewPostMediaRepo(pg *postgres.Postgres) *PostMediaPostgresRepo {
	return &PostMediaPostgresRepo{pg}
}

// Create -.
func (r *PostMediaPostgresRepo) Create(ctx context.Context, postID int, media []models.PostMedia) error {
	builder := r.Builder.
		Insert("posts_media").
		Columns("post_id", "type", "file_path", "file_extension")

	for _, m := range media {
		builder = builder.Values(postID, m.Type, m.Source, m.Extension)
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("PostMediaPostgresRepo - Create - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PostMediaPostgresRepo - Create - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *PostMediaPostgresRepo) GetPostMedia(ctx context.Context, postID int) ([]models.PostMedia, error) {
	sql, args, err := r.Builder.
		Select("id, type, file_path, file_extension, created_at").
		From("posts_media").
		Where("post_id = ?", postID).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("PostMediaPostgresRepo - GetPostMedia - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("PostMediaPostgresRepo - GetPostMedia - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	postMedia := make([]models.PostMedia, 0)
	for rows.Next() {
		pm := models.PostMedia{
			Post: models.Post{},
		}

		err = rows.Scan(&pm.ID, &pm.Type, &pm.Source, &pm.Extension, &pm.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("PostMediaPostgresRepo - GetPostMedia - rows.Scan: %w", err)
		}

		postMedia = append(postMedia, pm)
	}

	return postMedia, nil
}

func (r *PostMediaPostgresRepo) GetMedia(ctx context.Context, mediaID int) (*models.PostMedia, error) {
	sql, _, err := r.Builder.
		Select("id, post_id, type, file_path, file_extension, created_at").
		From("posts_media").
		Where("id = ?", mediaID).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("PostMediaPostgresRepo - GetMedia - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("PostMediaPostgresRepo - GetMedia - r.Pool.Query: %w", err)
	}

	pm := models.PostMedia{}

	err = row.Scan(&pm.ID, &pm.Post.PostID, &pm.Type, &pm.Source, &pm.Extension, &pm.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("PostMediaPostgresRepo - GetPostMedia - rows.Scan: %w", err)
	}

	return &pm, nil
}

func (r *PostMediaPostgresRepo) Delete(ctx context.Context, m models.PostMedia) error {
	sql, args, err := r.Builder.
		Delete("posts_media").
		Where("id = ?", m.ID).
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
