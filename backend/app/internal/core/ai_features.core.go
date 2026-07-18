package core

import (
	"context"
	"fmt"

	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/internal/services/aiclient"
	"backend.app/pkg/response"
	"github.com/google/uuid"
)

func (c *Core) GetNearbyCelebrations(ctx context.Context, latitude, longitude, radiusKM float64, limit int) *dtos.ResponseObject {
	if radiusKM <= 0 || radiusKM > 500 {
		radiusKM = 50
	}
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	rows, err := c.postgres.Pool.Query(ctx, `
		SELECT id, COALESCE(caption, ''), COALESCE(notes, ''), latitude, longitude,
		       6371 * acos(LEAST(1, cos(radians($1)) * cos(radians(latitude)) *
		       cos(radians(longitude) - radians($2)) + sin(radians($1)) * sin(radians(latitude)))) AS distance_km,
		       celebration_date, created_at
		FROM celebrations
		WHERE status = 'active' AND expires_at > CURRENT_TIMESTAMP
		  AND latitude IS NOT NULL AND longitude IS NOT NULL
		  AND 6371 * acos(LEAST(1, cos(radians($1)) * cos(radians(latitude)) *
		      cos(radians(longitude) - radians($2)) + sin(radians($1)) * sin(radians(latitude)))) <= $3
		ORDER BY distance_km ASC LIMIT $4`, latitude, longitude, radiusKM, limit)
	if err != nil {
		return response.ServerErrorResponse(err)
	}
	defer rows.Close()
	items := make([]dtos.NearbyCelebration, 0)
	for rows.Next() {
		var item dtos.NearbyCelebration
		if err := rows.Scan(&item.ID, &item.Caption, &item.Notes, &item.Latitude, &item.Longitude, &item.DistanceKM, &item.CelebrationDate, &item.CreatedAt); err != nil {
			return response.ServerErrorResponse(err)
		}
		items = append(items, item)
	}
	return response.SuccessResponse("Nearby celebrations retrieved successfully", items)
}

func canModerate(user *models.User) bool { return user.Role == "admin" || user.Role == "moderator" }

func (c *Core) GetModerationQueue(ctx context.Context, user *models.User, status string, limit int) *dtos.ResponseObject {
	if !canModerate(user) {
		return response.ForbiddenResponse(fmt.Errorf("insufficient permissions"))
	}
	if status == "" {
		status = "pending"
	}
	if status != "pending" && status != "approved" && status != "rejected" {
		return response.BadRequestResponse(fmt.Errorf("invalid moderation status"))
	}
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	rows, err := c.postgres.Pool.Query(ctx, `SELECT id, celebration_id, user_id, toxicity_score,
		flagged_content, status, moderated_by, moderated_at, created_at
		FROM moderation_queue WHERE status = $1 ORDER BY created_at ASC LIMIT $2`, status, limit)
	if err != nil {
		return response.ServerErrorResponse(err)
	}
	defer rows.Close()
	items := make([]dtos.ModerationQueueItem, 0)
	for rows.Next() {
		var item dtos.ModerationQueueItem
		if err := rows.Scan(&item.ID, &item.CelebrationID, &item.UserID, &item.ToxicityScore, &item.FlaggedContent, &item.Status, &item.ModeratorID, &item.ModeratedAt, &item.CreatedAt); err != nil {
			return response.ServerErrorResponse(err)
		}
		items = append(items, item)
	}
	return response.SuccessResponse("Moderation queue retrieved successfully", items)
}

func (c *Core) ResolveModeration(ctx context.Context, user *models.User, queueID uuid.UUID, decision string) *dtos.ResponseObject {
	if !canModerate(user) {
		return response.ForbiddenResponse(fmt.Errorf("insufficient permissions"))
	}
	if decision != "approved" && decision != "rejected" {
		return response.BadRequestResponse(fmt.Errorf("invalid decision"))
	}
	tx, err := c.postgres.Pool.Begin(ctx)
	if err != nil {
		return response.ServerErrorResponse(err)
	}
	defer tx.Rollback(ctx)
	var celebrationID uuid.UUID
	if err := tx.QueryRow(ctx, `UPDATE moderation_queue SET status=$1, moderated_by=$2, moderated_at=CURRENT_TIMESTAMP,
		updated_at=CURRENT_TIMESTAMP WHERE id=$3 AND status='pending' RETURNING celebration_id`, decision, user.ID, queueID).Scan(&celebrationID); err != nil {
		return response.BadRequestResponse(fmt.Errorf("pending moderation item not found"))
	}
	celebrationStatus := "active"
	if decision == "rejected" {
		celebrationStatus = "expired"
	}
	if _, err := tx.Exec(ctx, `UPDATE celebrations SET status=$1, updated_at=CURRENT_TIMESTAMP WHERE id=$2`, celebrationStatus, celebrationID); err != nil {
		return response.ServerErrorResponse(err)
	}
	if err := tx.Commit(ctx); err != nil {
		return response.ServerErrorResponse(err)
	}
	if decision == "approved" {
		var content string
		if err := c.postgres.Pool.QueryRow(ctx, `SELECT trim(concat_ws(' ', caption, notes)) FROM celebrations WHERE id=$1`, celebrationID).Scan(&content); err == nil && content != "" {
			_, _ = c.aiClient.IndexCelebration(ctx, &aiclient.IndexCelebrationRequest{CelebrationID: celebrationID.String(), Text: content})
		}
	}
	return response.SuccessResponse("Moderation item resolved successfully", map[string]interface{}{"id": queueID, "decision": decision})
}
