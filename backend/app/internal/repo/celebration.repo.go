package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"

	"backend.app/common/helpers"
	"backend.app/common/messages"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
)

// CreateCelebration adds a new celebration for a user
func (r *Repo) CreateCelebration(ctx context.Context, celebration *models.Celebration) (*models.Celebration, error) {
	sql, args, err := r.postgres.Builder.
		Insert(Celebrations).
		Columns("audience,celebration_type,celebration_kind,owner,owner_identifier,owner_id,status,frequency,created_by,expires_at,notes,selected_friends,longitude,latitude,caption,celebration_date,created_at").
		Values(celebration.Audience,
			celebration.CelebrationType,
			celebration.CelebrationKind,
			celebration.Owner,
			celebration.OwnerIdentifier,
			celebration.OwnerId,
			celebration.Status,
			celebration.Frequency,
			celebration.CreatedBy,
			celebration.ExpiresAt,
			celebration.Notes,
			celebration.SelectedFriends,
			celebration.Long,
			celebration.Lat,
			celebration.Caption,
			celebration.CelebrationDate,
			celebration.CreatedAt).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		r.log.Error("Repo - CreateCelebration - r.Builder: %w", err)
		return nil, errors.New("something went wrong")
	}
	row := r.postgres.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&celebration.ID)
	if err != nil {
		r.log.Error("Repo - CreateCelebration - r.Pool.Scan: %w", err)
		return nil, errors.New("something went wrong")
	}

	return celebration, nil
}

// GetCelebrationByField gets a celebration by set fields
func (r *Repo) GetCelebrationByField(ctx context.Context, filter map[string]interface{}) (*dtos.Celebration, error) {
	sql, args, err := r.postgres.Builder.
		Select("c.id,c.celebration_type,c.celebration_kind,c.audience,c.owner,c.owner_identifier,c.owner_id,c.status,c.frequency,c.created_by,c.expires_at,c.longitude,c.latitude,c.notes,c.selected_friends,c.caption,c.celebration_date,c.created_at, json_agg(json_build_object('id',m.id,'object_ref',m.object_ref,'object_id',m.object_id,'media_type',m.media_type,'media_url',m.media_url)) as media").
		From("celebrations c").
		Join("media m ON m.owner_id = c.id").
		Where(squirrel.Eq(filter)).
		ToSql()

	if err != nil {
		r.log.Error("unable to build query: %w", err)
		return nil, errors.New("something went wrong")
	}
	row := r.postgres.Pool.QueryRow(ctx, sql, args...)

	celebration := dtos.Celebration{}
	var longitude, latitude pgtype.Float4
	err = row.Scan(
		&celebration.ID,
		&celebration.CelebrationType,
		&celebration.CelebrationKind,
		&celebration.Audience,
		&celebration.Owner,
		&celebration.OwnerIdentifier,
		&celebration.OwnerId,
		&celebration.Status,
		&celebration.Frequency,
		&celebration.CreatedBy,
		&celebration.ExpiresAt,
		&longitude,
		&latitude,
		&celebration.Notes,
		&celebration.SelectedFriends,
		&celebration.Caption,
		&celebration.CelebrationDate,
		&celebration.CreatedAt,
		&celebration.Media,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, messages.ErrCelebrationNotFound
		}
		r.log.Error("row.Scan: %w", err)
		return nil, errors.New("something went wrong")
	}
	return &celebration, nil
}

// DeleteCelebration deletes a celebration based on a passed id
func (r *Repo) DeleteCelebration(ctx context.Context, fields helpers.Map) error {
	sql, args, err := r.postgres.Builder.
		Delete(Celebrations).
		Where(squirrel.Eq(fields)).
		ToSql()

	if err != nil {
		r.log.Debug("Repo - DeleteCelebration - r.Builder: %w", err)
		return errors.New("something went wrong")
	}

	_, err = r.postgres.Pool.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Debug("Repo - DeleteCelebration - r.Pool.Exec: %w", err)
		return errors.New("something went wrong")
	}
	return nil
}

// UpdateCelebration updates a celebration by id
func (r *Repo) UpdateCelebration(ctx context.Context, Id uuid.UUID, user *models.User, fields map[string]interface{}) error {
	sql, args, err := r.postgres.Builder.
		Update(Celebrations).
		SetMap(squirrel.Eq(fields)).
		Where(squirrel.Eq{"id": Id, "owner_id": user.ID}).
		ToSql()

	if err != nil {
		r.log.Debug("Repo - UpdateCelebration - r.Builder: %w", err)
		return errors.New("something went wrong")
	}

	_, err = r.postgres.Pool.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Debug("Repo - UpdateCelebration - r.Pool.Exec: %w", err)
		return errors.New("something went wrong")
	}

	return nil
}

// GetAllCelebrations Gets all Celebrations with pagination from DB
func (r *Repo) GetAllCelebrations(ctx context.Context, user *models.User, query *dtos.APIPagingDto) (*dtos.CelebrationsResponse, error) {
	query, _ = getPaginationInfo(query)

	isFirstPage := query.Cursor == ""
	pointsNext := false
	builder := r.postgres.Builder.
		Select("c.id,c.celebration_type,c.celebration_kind,c.audience,c.owner,c.owner_identifier,c.owner_id,c.status,c.frequency,c.created_by,c.expires_at,c.longitude,c.latitude,c.notes,c.selected_friends,c.caption,c.celebration_date,c.created_at,json_agg(json_build_object('id',m.id,'object_ref',m.object_ref,'object_id',m.object_id,'media_type',m.media_type,'media_url',m.media_url)) as media").
		From("celebrations c").
		Join("media m ON m.object_id = c.id").
		GroupBy("c.id").
		Where("c.status = ?", "active")

	whereMap := getFilterFromQuery(query.Filter)
	builder = buildWhere(builder, whereMap)

	if query.Cursor != "" {
		decodedCursor, err := helpers.DecodeCursor(query.Cursor)
		if err != nil {
			r.log.Debug("GetAllCelebrations: DecodeCursor error : %v", err)
			return nil, err
		}
		pointsNext = decodedCursor["points_next"] == true

		operator, order := getPaginationOperator(pointsNext, query.Direction)
		whereStr := fmt.Sprintf("(c.created_at %s ? OR (c.created_at = ? AND c.id %s ?))", operator, operator)
		builder = builder.Where(whereStr, decodedCursor["created_at"], decodedCursor["created_at"], decodedCursor["id"])
		if order != "" {
			query.Direction = order
		}
	}

	// add limit
	builder = builder.Limit(uint64(query.Limit + 1))
	// add sort
	builder = builder.OrderBy(fmt.Sprintf("%s %s", query.Sort, query.Direction))

	sql, args, err := builder.ToSql()
	if err != nil {
		r.log.Debug("unable to build query: %w", err)
		return nil, errors.New("something went wrong")
	}

	rows, err := r.postgres.Pool.Query(ctx, sql, args...)
	if err != nil {
		r.log.Debug("r.Pool.Query: %w", err)
		return nil, errors.New("something went wrong")
	}
	defer rows.Close()

	celebrations := make([]*dtos.Celebration, 0)
	for rows.Next() {
		celebration := dtos.Celebration{}
		var longitude, latitude pgtype.Float4
		err = rows.Scan(
			&celebration.ID,
			&celebration.CelebrationType,
			&celebration.CelebrationKind,
			&celebration.Audience,
			&celebration.Owner,
			&celebration.OwnerIdentifier,
			&celebration.OwnerId,
			&celebration.Status,
			&celebration.Frequency,
			&celebration.CreatedBy,
			&celebration.ExpiresAt,
			&longitude,
			&latitude,
			&celebration.Notes,
			&celebration.SelectedFriends,
			&celebration.Caption,
			&celebration.CelebrationDate,
			&celebration.CreatedAt,
			&celebration.Media,
		)

		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		celebrations = append(celebrations, &celebration)
	}

	hasPagination := len(celebrations) > query.Limit
	if hasPagination {
		celebrations = celebrations[:query.Limit]
	}
	if !isFirstPage && !pointsNext {
		celebrations = helpers.Reverse(celebrations)
	}

	var cursorData CursorData
	query.Limit = len(celebrations)
	if len(celebrations) > 0 {
		cursorData.FirstId = celebrations[0].ID.String()
		cursorData.FirstCreatedAt = celebrations[0].CreatedAt
		cursorData.LastId = celebrations[query.Limit-1].ID.String()
		cursorData.LastCreatedAt = celebrations[query.Limit-1].CreatedAt
	}

	pageInfo := calculatePagination(isFirstPage, hasPagination, cursorData, pointsNext)
	return &dtos.CelebrationsResponse{
		Celebrations: celebrations,
		PagingInfo: dtos.PagingInfo{
			NextCursor: pageInfo.NextCursor,
			PrevCursor: pageInfo.PrevCursor,
		},
	}, nil
}
