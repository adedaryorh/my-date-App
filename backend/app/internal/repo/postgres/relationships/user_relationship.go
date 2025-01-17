package relationships

import (
	"context"
	"fmt"

	"backend.app/internal/models"
	"backend.app/pkg/postgres"
)

// RelationshipPostgresRepo -.
type RelationshipPostgresRepo struct {
	*postgres.Postgres
}

// NewRelationshipsRepo -.
func NewRelationshipsRepo(pg *postgres.Postgres) *RelationshipPostgresRepo {
	return &RelationshipPostgresRepo{pg}
}

// Create -.
func (r *RelationshipPostgresRepo) Create(ctx context.Context, p *models.Relationship) error {
	sql, args, err := r.Builder.
		Insert("user_relationships").
		Columns("sender_user_id, receiver_user_id, status").
		Values(p.SenderUserID, p.ReceiverUserID, p.Status).
		Suffix("RETURNING \"id\", \"created_at\", \"updated_at\"").
		ToSql()

	if err != nil {
		return fmt.Errorf("RelationshipPostgresRepo - Create - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("RelationshipPostgresRepo - Create - r.Pool.Scan: %w", err)
	}

	return nil
}

// GetUserRelationships -.
func (r *RelationshipPostgresRepo) GetUserRelationships(ctx context.Context, userID int) ([]models.Relationship, error) {
	builder := r.Builder.
		Select("ur.id, s.user_id, s.first_name, s.last_name, ur.sender_user_id, r.user_id, r.first_name, r.last_name, ur.receiver_user_id, ur.status, ur.created_at, ur.updated_at").
		From("user_relationships ur").
		InnerJoin("users s ON s.id = sender_user_id").
		InnerJoin("users r ON r.id = receiver_user_id").
		Where("sender_user_id = ? OR receiver_user_id = ?", userID, userID).
		OrderBy("created_at DESC")

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("RelationshipPostgresRepo - GetUserRelationships - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("RelationshipPostgresRepo - GetUserRelationships - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	posts := make([]models.Relationship, 0)
	for rows.Next() {
		p := models.Relationship{
			Sender:   models.User{},
			Receiver: models.User{},
		}

		err = rows.Scan(
			&p.ID,
			&p.Sender.UserID,
			&p.Sender.FirstName,
			&p.Sender.LastName,
			&p.SenderUserID,
			&p.Receiver.UserID,
			&p.Receiver.FirstName,
			&p.Receiver.LastName,
			&p.ReceiverUserID,
			&p.Status,
			&p.CreatedAt,
			&p.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("RelationshipPostgresRepo - GetUserRelationships - rows.Scan: %w", err)
		}

		posts = append(posts, p)
	}

	return posts, nil
}

func (r *RelationshipPostgresRepo) GetRelationship(ctx context.Context, senderUserID int, receiverUserID int) (*models.Relationship, error) {
	builder := r.Builder.
		Select("ur.id, s.user_id, s.first_name, s.last_name, sender_user_id, r.user_id, r.first_name, r.last_name, ur.receiver_user_id, ur.status, ur.created_at, ur.updated_at").
		From("user_relationships ur").
		InnerJoin("users s ON s.id = sender_user_id").
		InnerJoin("users r ON r.id = receiver_user_id").
		Where("sender_user_id = ? AND receiver_user_id = ?", senderUserID, receiverUserID).
		OrderBy("created_at DESC")

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("RelationshipPostgresRepo - GetUserRelationships - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("RelationshipPostgresRepo - Get - r.Pool.Query: %w", err)
	}

	p := models.Relationship{}
	err = row.Scan(
		&p.ID,
		&p.Sender.UserID,
		&p.Sender.FirstName,
		&p.Sender.LastName,
		&p.SenderUserID,
		&p.Receiver.UserID,
		&p.Receiver.FirstName,
		&p.Receiver.LastName,
		&p.ReceiverUserID,
		&p.Status,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("RelationshipPostgresRepo - Get - rows.Scan: %w", err)
	}

	return &p, nil
}

// Delete -.
func (r *RelationshipPostgresRepo) Delete(ctx context.Context, relationshipID int) error {
	sql, args, err := r.Builder.
		Delete("user_relationships").
		Where("id = ?", relationshipID).
		ToSql()

	if err != nil {
		return fmt.Errorf("RelationshipPostgresRepo - Delete - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("RelationshipPostgresRepo - Delete - r.Pool.Exec: %w", err)
	}

	return nil

}
