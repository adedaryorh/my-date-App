package repo

import (
	"backend.app/common/helpers"
	"backend.app/internal/dtos"
	"backend.app/internal/mappers"
	"context"
	_ "encoding/json"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"backend.app/common/messages"
	"backend.app/internal/models"
)

var notificationDtos []dtos.NotificationDto

// CreateNotification
func (r *Repo) CreateNotification(ctx context.Context, notification *models.Notification) (*models.Notification, error) {
	sql, args, err := r.postgres.Builder.
		Insert("notifications").
		Columns("owner, owner_id, status, template, title, notification_type, content, meta_data").
		Values(notification.Owner, notification.OwnerId, notification.Status, notification.Template, notification.Title, notification.NotificationType, notification.Content, notification.MetaData).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		r.log.Error("NotificationPostgresRepo - CreateNotification - r.Builder: %w", err)
		return nil, errors.New("something went wrong")
	}

	row := r.postgres.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&notification.Id)
	if err != nil {
		r.log.Error("NotificationPostgresRepo - CreateNotification - r.Pool.Scan: %w", err)
		return nil, errors.New("something went wrong")
	}
	return notification, nil
}

// GetNotificationByID
func (r *Repo) GetNotificationById(ctx context.Context, notificationID uuid.UUID) (*models.Notification, error) {
	sql, args, err := r.postgres.Builder.
		Select("id, owner, owner_id, status, template, title, notification_type, content, meta_data, created_at, updated_at").
		From("notifications").
		Where(squirrel.Eq{"id": notificationID}).
		ToSql()

	if err != nil {
		r.log.Debug("unable to build query: %w", err)
		return nil, errors.New("something went wrong")
	}

	row := r.postgres.Pool.QueryRow(ctx, sql, args...)

	n := models.Notification{}
	err = row.Scan(
		&n.Id,
		&n.Owner,
		&n.OwnerId,
		&n.Status,
		&n.Template,
		&n.Title,
		&n.NotificationType,
		&n.Content,
		&n.MetaData,
		&n.CreatedAt,
		&n.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		r.log.Debug("row.Scan: %w", err)
		return nil, errors.New("something went wrong")
	}

	return &n, nil
}

func (r *Repo) GetAllNotifications(ctx context.Context, query *dtos.APIPagingDto) (*dtos.NotificationsResponse, error) {
	isFirstPage := query.Cursor == ""
	pointsNext := false
	builder := r.postgres.Builder.
		Select("id, owner, owner_id, status, template, title, notification_type, content, meta_data, created_at, updated_at").
		From("notifications")

	whereMap := getFilterFromQuery(query.Filter)
	builder = buildWhere(builder, whereMap)

	if query.Cursor != "" {
		decodedCursor, err := helpers.DecodeCursor(query.Cursor)
		if err != nil {
			r.log.Debug("GetAllNotifications: DecodeCursor error : %v", err)
			return nil, err
		}
		pointsNext = decodedCursor["points_next"] == true
		operator, order := getPaginationOperator(pointsNext, query.Direction)
		whereStr := fmt.Sprintf("(created_at %s ? OR (created_at = ? AND id %s ?))", operator, operator)
		builder = builder.Where(whereStr, decodedCursor["created_at"], decodedCursor["created_at"], decodedCursor["id"])
		if order != "" {
			query.Direction = order
		}
	}
	builder = builder.Limit(uint64(query.Limit + 1))
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

	notifications := make([]*models.Notification, 0)
	for rows.Next() {
		n := models.Notification{}
		err = rows.Scan(
			&n.Id,
			&n.Owner,
			&n.OwnerId,
			&n.Status,
			&n.Template,
			&n.Title,
			&n.NotificationType,
			&n.Content,
			&n.MetaData,
			&n.CreatedAt,
			&n.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		notifications = append(notifications, &n)
	}

	hasPagination := len(notifications) > query.Limit
	if hasPagination {
		notifications = notifications[:query.Limit]
	}
	if !isFirstPage && !pointsNext {
		notifications = helpers.Reverse(notifications)
	}

	var cursorData CursorData
	query.Limit = len(notifications)
	if len(notifications) > 0 {
		cursorData.FirstId = notifications[0].Id.String()
		cursorData.FirstCreatedAt = notifications[0].CreatedAt
		cursorData.LastId = notifications[query.Limit-1].Id.String()
		cursorData.LastCreatedAt = notifications[query.Limit-1].CreatedAt
	}

	pageInfo := calculatePagination(isFirstPage, hasPagination, cursorData, pointsNext)
	var not mappers.DtoNotificationMapper
	var notificationDtos []dtos.NotificationDto

	for _, notification := range notifications {
		u := not.MapNotificationToDto(notification)     // u is *dtos.NotificationDto
		notificationDtos = append(notificationDtos, *u) // Dereference the pointer
	}
	return &dtos.NotificationsResponse{
		Notifications: notificationDtos,
		PagingInfo: dtos.PagingInfo{
			NextCursor: pageInfo.NextCursor,
			PrevCursor: pageInfo.PrevCursor,
		},
	}, nil
}

// UpdateNotification(this will update dynamically with any value)
func (r *Repo) UpdateNotification(ctx context.Context, id uuid.UUID, fields map[string]interface{}) error {
	sql, args, err := r.postgres.Builder.
		Update("notifications").
		SetMap(squirrel.Eq(fields)).
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		r.log.Debug("Notification PostgresRepo - UpdateNotification - r.Builder: %w", err)
		return errors.New("something went wrong")
	}

	_, err = r.postgres.Pool.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Debug("Notification PostgresRepo - UpdateNotification - r.Pool.Exec: %w", err)
		return errors.New("something went wrong")
	}

	return nil
}

// GetSingleNotification retrieves a single notification by Either its ID or another unique field.
func (r *Repo) GetSingleNotification(ctx context.Context, filter map[string]interface{}) (*models.Notification, error) {
	sql, args, err := r.postgres.Builder.
		Select("id, owner, owner_id, status, template, title, notification_type, content, meta_data, created_at, updated_at").
		From("notifications").
		Where(squirrel.Eq(filter)).
		ToSql()

	if err != nil {
		r.log.Error("unable to build query: %w", err)
		return nil, errors.New("something went wrong")
	}

	row := r.postgres.Pool.QueryRow(ctx, sql, args...)

	n := models.Notification{}
	err = row.Scan(
		&n.Id,
		&n.Owner,
		&n.OwnerId,
		&n.Status,
		&n.Template,
		&n.Title,
		&n.NotificationType,
		&n.Content,
		&n.MetaData,
		&n.CreatedAt,
		&n.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, messages.ErrNotificationNotFound
		}
		r.log.Error("row.Scan: %w", err)
		return nil, errors.New("something went wrong")
	}

	return &n, nil
}
