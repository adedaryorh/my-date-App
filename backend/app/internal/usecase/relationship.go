package usecase

import (
	"context"
	"fmt"

	"backend.app/internal/models"
	"backend.app/internal/repo"
)

type Relationship interface {
	Create(context.Context, *models.Relationship) error
	GetUserRelationships(ctx context.Context, userID int) ([]models.Relationship, error)
	GetRelationship(ctx context.Context, senderUserID int, receiverUserID int) (*models.Relationship, error)
	Delete(context.Context, models.Relationship) error
}

// RelationshipUseCase -.
type RelationshipUseCase struct {
	relationshipsRepo repo.UserRelationship
}

// NewRelationshipUseCase -.
func NewRelationshipUseCase(r repo.UserRelationship) *RelationshipUseCase {
	return &RelationshipUseCase{
		relationshipsRepo: r,
	}
}

// Create -.
func (uc *RelationshipUseCase) Create(ctx context.Context, p *models.Relationship) error {
	err := uc.relationshipsRepo.Create(ctx, p)

	if err != nil {
		return fmt.Errorf("RelationshipUseCase - Relationships - s.postsRepo.Create: %w", err)
	}

	return nil
}

// GetUserRelationships -.
func (uc *RelationshipUseCase) GetUserRelationships(ctx context.Context, userID int) ([]models.Relationship, error) {
	relationships, err := uc.relationshipsRepo.GetUserRelationships(ctx, userID)

	if err != nil {
		return nil, fmt.Errorf("RelationshipUseCase - Relationships - unable to get: %w", err)
	}

	return relationships, nil
}

// GetRelationship -.
func (uc *RelationshipUseCase) GetRelationship(ctx context.Context, senderUserID int, receiverUserID int) (*models.Relationship, error) {
	relationship, err := uc.relationshipsRepo.GetRelationship(ctx, senderUserID, receiverUserID)
	if err != nil {
		return nil, fmt.Errorf("RelationshipUseCase - Relationships - unable to get: %w", err)
	}

	return relationship, nil
}

// Delete -.
func (uc *RelationshipUseCase) Delete(ctx context.Context, r models.Relationship) error {
	err := uc.relationshipsRepo.Delete(ctx, r.ID)
	if err != nil {
		return fmt.Errorf("RelationshipUseCase -Relationships - unable to delete: %w", err)
	}

	return nil
}
