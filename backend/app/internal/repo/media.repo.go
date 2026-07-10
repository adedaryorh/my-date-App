package repo

import (
	"context"
	"errors"

	"backend.app/common/helpers"
	"backend.app/internal/models"
	"github.com/Masterminds/squirrel"
)

// CreateMedia adds a new celebration for a user
func (r *Repo) CreateMedia(ctx context.Context, media *models.Media) (*models.Media, error) {
	sql, args, err := r.postgres.Builder.
		Insert(Media).
		Columns("object_ref,object_id,media_type,media_url,owner_id,created_at").
		Values(media.ObjectRef, media.ObjectId, media.MediaType, media.MediaUrl, media.OwnerId, media.CreatedAt).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		r.log.Error("Repo - CreateMedia - r.Builder: %w", err)
		return nil, errors.New("something went wrong")
	}
	row := r.postgres.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&media.ID)
	if err != nil {
		r.log.Error("Repo - CreateMedia - r.Pool.Scan: %w", err)
		return nil, errors.New("something went wrong")
	}

	return media, nil
}

// DeleteMedia deletes a media based on a passed fields
func (r *Repo) DeleteMedia(ctx context.Context, fields helpers.Map) error {
	sql, args, err := r.postgres.Builder.
		Delete(Media).
		Where(squirrel.Eq(fields)).
		ToSql()

	if err != nil {
		r.log.Debug("Repo - DeleteMedia - r.Builder: %w", err)
		return errors.New("something went wrong")
	}

	_, err = r.postgres.Pool.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Debug("Repo - DeleteMedia - r.Pool.Exec: %w", err)
		return errors.New("something went wrong")
	}
	return nil
}
