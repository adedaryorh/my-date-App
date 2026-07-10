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

// CreateFollower adds a new follower for a user
func (r *Repo) CreateFollower(ctx context.Context, follower *models.Follower) (*models.Follower, error) {
	sql, args, err := r.postgres.Builder.
		Insert(Followers).
		Columns("user_id, follower_id,created_at").
		Values(follower.UserId, follower.FollowerId, follower.CreatedAt).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		r.log.Error("Repo - CreateFollower - r.Builder: %w", err)
		return nil, errors.New("something went wrong")
	}
	row := r.postgres.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&follower.ID)
	if err != nil {
		r.log.Error("Repo - CreateFollower - r.Pool.Scan: %w", err)
		return nil, errors.New("something went wrong")
	}

	return follower, nil
}

// GetFollowerByField gets a follower by set fields
func (r *Repo) GetFollowerByField(ctx context.Context, filter map[string]interface{}) (*dtos.Follower, error) {
	sql, args, err := r.postgres.Builder.
		Select("f.id,f.follower_id,u.first_name,u.last_name,u.username,u.profile_image_url").
		From("followers f").
		Join("users u  ON u.id = f.follower_id").
		Where(squirrel.Eq(filter)).
		ToSql()

	if err != nil {
		r.log.Error("unable to build query: %w", err)
		return nil, errors.New("something went wrong")
	}
	row := r.postgres.Pool.QueryRow(ctx, sql, args...)

	follower := dtos.Follower{}

	err = row.Scan(
		&follower.ID,
		&follower.UserId,
		&follower.FirstName,
		&follower.LastName,
		&follower.Username,
		&follower.ProfileImage,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, messages.ErrFollowerNotFound
		}
		r.log.Error("row.Scan: %w", err)
		return nil, errors.New("something went wrong")
	}
	return &follower, nil
}

// DeleteFollower deletes a follower based on a passed id
func (r *Repo) DeleteFollower(ctx context.Context, fields helpers.Map) error {
	sql, args, err := r.postgres.Builder.
		Delete(Followers).
		Where(squirrel.Eq(fields)).
		ToSql()

	if err != nil {
		r.log.Debug("Repo - DeleteFollower - r.Builder: %w", err)
		return errors.New("something went wrong")
	}

	_, err = r.postgres.Pool.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Debug("Repo - DeleteFollower - r.Pool.Exec: %w", err)
		return errors.New("something went wrong")
	}
	return nil
}

// GetAllFollowers Gets all Users followers with pagination from DB
func (r *Repo) GetAllFollowers(ctx context.Context, user *models.User, query *dtos.APIPagingDto) (*dtos.FollowersResponse, error) {
	isFirstPage := query.Cursor == ""
	pointsNext := false
	builder := r.postgres.Builder.
		Select("f.id, u.id,u.first_name, u.last_name, u.username,u.profile_image_url,f.created_at").
		From("followers f").
		Join("users u ON f.follower_id = u.id").
		Join("blocked b ON b.user_id=f.user_id").
		Where("f.user_id = ? AND b.blocked_user_id <> f.follower_id", user.ID)

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
		whereStr := fmt.Sprintf("(f.created_at %s ? OR (f.created_at = ? AND f.id %s ?))", operator, operator)
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

	followers := make([]*dtos.Follower, 0)
	for rows.Next() {

		f := dtos.Follower{}
		err = rows.Scan(
			&f.ID,
			&f.UserId,
			&f.FirstName,
			&f.LastName,
			&f.Username,
			&f.ProfileImage,
			&f.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		followers = append(followers, &f)
	}

	hasPagination := len(followers) > query.Limit
	if hasPagination {
		followers = followers[:query.Limit]
	}
	if !isFirstPage && !pointsNext {
		followers = helpers.Reverse(followers)
	}

	var cursorData CursorData
	query.Limit = len(followers)
	if len(followers) > 0 {
		cursorData.FirstId = followers[0].ID.String()
		cursorData.FirstCreatedAt = followers[0].CreatedAt
		cursorData.LastId = followers[query.Limit-1].ID.String()
		cursorData.LastCreatedAt = followers[query.Limit-1].CreatedAt
	}

	pageInfo := calculatePagination(isFirstPage, hasPagination, cursorData, pointsNext)
	return &dtos.FollowersResponse{
		Followers: followers,
		PagingInfo: dtos.PagingInfo{
			NextCursor: pageInfo.NextCursor,
			PrevCursor: pageInfo.PrevCursor,
		},
	}, nil
}

// GetAllFriends Gets all Users friends with pagination from DB
func (r *Repo) GetAllFriends(ctx context.Context, user *models.User, query *dtos.APIPagingDto) (*dtos.FollowersResponse, error) {
	isFirstPage := query.Cursor == ""
	pointsNext := false
	builder := r.postgres.Builder.
		Select("f.id, u.id, u.first_name, u.last_name, u.username,u.profile_image_url,f.created_at").
		From("followers f").
		Join("users u ON f.follower_id = u.id").
		Join("blocked b ON b.user_id=f.user_id").
		Where("(f.user_id = ? OR f.follower_id = ? )AND b.blocked_user_id <> f.follower_id AND b.blocked_user_id <> f.user_id", user.ID, user.ID)

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
		whereStr := fmt.Sprintf("(f.created_at %s ? OR (f.created_at = ? AND f.id %s ?))", operator, operator)
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

	followers := make([]*dtos.Follower, 0)
	for rows.Next() {

		f := dtos.Follower{}
		err = rows.Scan(
			&f.ID,
			&f.UserId,
			&f.FirstName,
			&f.LastName,
			&f.Username,
			&f.ProfileImage,
			&f.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		followers = append(followers, &f)
	}

	hasPagination := len(followers) > query.Limit
	if hasPagination {
		followers = followers[:query.Limit]
	}
	if !isFirstPage && !pointsNext {
		followers = helpers.Reverse(followers)
	}

	var cursorData CursorData
	query.Limit = len(followers)
	if len(followers) > 0 {
		cursorData.FirstId = followers[0].ID.String()
		cursorData.FirstCreatedAt = followers[0].CreatedAt
		cursorData.LastId = followers[query.Limit-1].ID.String()
		cursorData.LastCreatedAt = followers[query.Limit-1].CreatedAt
	}

	pageInfo := calculatePagination(isFirstPage, hasPagination, cursorData, pointsNext)
	return &dtos.FollowersResponse{
		Followers: followers,
		PagingInfo: dtos.PagingInfo{
			NextCursor: pageInfo.NextCursor,
			PrevCursor: pageInfo.PrevCursor,
		},
	}, nil
}
