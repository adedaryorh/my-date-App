package repo

import (
	"context"
	"errors"
	"fmt"

	"backend.app/common/helpers"
	"backend.app/common/messages"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

// CreateBlock adds a new blocked user
func (r *Repo) CreateBlock(ctx context.Context, blocked *models.Blocked) (*models.Blocked, error) {
	sql, args, err := r.postgres.Builder.
		Insert(Followers).
		Columns("user_id, blocked_user_id,created_at").
		Values(blocked.UserId, blocked.BlockedUserId, blocked.CreatedAt).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		r.log.Error("Repo - CreateBlockList - r.Builder: %w", err)
		return nil, errors.New("something went wrong")
	}
	row := r.postgres.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&blocked.ID)
	if err != nil {
		r.log.Error("Repo - CreateBlockList - r.Pool.Scan: %w", err)
		return nil, errors.New("something went wrong")
	}

	return blocked, nil
}

// GetBlockedUserByField gets a blocked user by set fields
func (r *Repo) GetBlockedUserByField(ctx context.Context, filter map[string]interface{}) (*dtos.Blocked, error) {
	sql, args, err := r.postgres.Builder.
		Select("b.id,b.user_id,b.blocked_user_id,u.first_name,u.last_name,u.username,u.profile_image_url").
		From("blocked b").
		Join("users u ON u.id = b.blocked_user_id").
		Where(squirrel.Eq(filter)).
		ToSql()

	if err != nil {
		r.log.Error("unable to build query: %w", err)
		return nil, errors.New("something went wrong")
	}
	row := r.postgres.Pool.QueryRow(ctx, sql, args...)

	blocked := dtos.Blocked{}

	err = row.Scan(
		&blocked.ID,
		&blocked.BlockedUserId,
		&blocked.FirstName,
		&blocked.LastName,
		&blocked.Username,
		&blocked.ProfileImage,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, messages.ErrBlockedUserNotFound
		}
		r.log.Error("row.Scan: %w", err)
		return nil, errors.New("something went wrong")
	}
	return &blocked, nil
}

// DeleteBlocked deletes a block list row based on a passed id
func (r *Repo) DeleteBlocked(ctx context.Context, fields helpers.Map) error {
	sql, args, err := r.postgres.Builder.
		Delete(Blocked).
		Where(squirrel.Eq(fields)).
		ToSql()

	if err != nil {
		r.log.Debug("Repo - DeleteBlocked - r.Builder: %w", err)
		return errors.New("something went wrong")
	}

	_, err = r.postgres.Pool.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Debug("Repo - DeleteBlocked - r.Pool.Exec: %w", err)
		return errors.New("something went wrong")
	}

	return nil
}

// GetAllBlocked Gets all blocked Users with pagination from DB
func (r *Repo) GetAllBlocked(ctx context.Context, user *models.User, query *dtos.APIPagingDto) (*dtos.BlockedResponse, error) {
	isFirstPage := query.Cursor == ""
	pointsNext := false
	builder := r.postgres.Builder.
		Select("b.id, b.blocked_user_id,u.first_name, u.last_name, u.username,u.profile_image_url,b.created_at").
		From("blocked b").
		Join("users u ON b.blocked_user_id = u.id").
		Where("b.user_id = ?", user.ID)

	whereMap := getFilterFromQuery(query.Filter)
	builder = buildWhere(builder, whereMap)

	if query.Cursor != "" {
		decodedCursor, err := helpers.DecodeCursor(query.Cursor)
		if err != nil {
			r.log.Debug("GetAllUsers: DecodeCursor error : %v", err)
			return nil, err
		}
		pointsNext = decodedCursor["points_next"] == true

		operator, order := getPaginationOperator(pointsNext, query.Direction)
		whereStr := fmt.Sprintf("(b.created_at %s ? OR (b.created_at = ? AND b.id %s ?))", operator, operator)
		builder = builder.Where(whereStr, decodedCursor["created_at"], decodedCursor["created_at"], decodedCursor["id"])
		if order != "" {
			query.Direction = order
		}
	}

	// add limit
	builder = builder.Limit(uint64(query.Limit + 1))
	// add sort
	builder = builder.OrderBy(query.Sort, query.Direction)

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

	allBlocked := make([]*dtos.Blocked, 0)
	for rows.Next() {

		blocked := dtos.Blocked{}
		err = rows.Scan(
			&blocked.ID,
			&blocked.BlockedUserId,
			&blocked.FirstName,
			&blocked.LastName,
			&blocked.Username,
			&blocked.ProfileImage,
			&blocked.BlockedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		allBlocked = append(allBlocked, &blocked)
	}

	hasPagination := len(allBlocked) > query.Limit
	if hasPagination {
		allBlocked = allBlocked[:query.Limit]
	}
	if !isFirstPage && !pointsNext {
		allBlocked = helpers.Reverse(allBlocked)
	}

	var cursorData CursorData
	query.Limit = len(allBlocked)
	if len(allBlocked) > 0 {
		cursorData.FirstId = allBlocked[0].ID.String()
		cursorData.FirstCreatedAt = allBlocked[0].BlockedAt
		cursorData.LastId = allBlocked[query.Limit-1].ID.String()
		cursorData.LastCreatedAt = allBlocked[query.Limit-1].BlockedAt
	}

	pageInfo := calculatePagination(isFirstPage, hasPagination, cursorData, pointsNext)
	return &dtos.BlockedResponse{
		Blocked: allBlocked,
		PagingInfo: dtos.PagingInfo{
			NextCursor: pageInfo.NextCursor,
			PrevCursor: pageInfo.PrevCursor,
		},
	}, nil
}
