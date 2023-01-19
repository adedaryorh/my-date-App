package usecase

import (
	"celebut-api/internal/models"
	"celebut-api/internal/repo"
	"context"
	"fmt"
)

type UserReaction interface {
	Create(context.Context, *models.UserReaction) error
	GetPostReactions(ctx context.Context, postID int) ([]models.UserReaction, error)
	GetUserReaction(ctx context.Context, userID int, postID int) (*models.UserReaction, error)
	DeleteUserReaction(ctx context.Context, userID int, postID int) error
}

// UserReactionUseCase -.
type UserReactionUseCase struct {
	reactionsRepo repo.UserReaction
}

// NewUserReactionUseCase -.
func NewUserReactionUseCase(r repo.UserReaction) *UserReactionUseCase {
	return &UserReactionUseCase{
		reactionsRepo: r,
	}
}

// Create -.
func (uc *UserReactionUseCase) Create(ctx context.Context, r *models.UserReaction) error {
	err := uc.reactionsRepo.Create(ctx, r)

	if err != nil {
		return fmt.Errorf("UserReactionUseCase - UserReactions - s.reactionsRepo.Create: %w", err)
	}

	return nil
}

// GetPostReactions -.
func (uc *UserReactionUseCase) GetPostReactions(ctx context.Context, postID int) ([]models.UserReaction, error) {
	reactions, err := uc.reactionsRepo.GetPostReactions(ctx, postID)

	if err != nil {
		return nil, fmt.Errorf("UserReactionUseCase - UserReactions - unable to get: %w", err)
	}

	return reactions, nil
}

// GetUserReaction -.
func (uc *UserReactionUseCase) GetUserReaction(ctx context.Context, userID int, postID int) (*models.UserReaction, error) {
	reaction, err := uc.reactionsRepo.GetReaction(ctx, userID, postID)
	if err != nil {
		return nil, fmt.Errorf("UserReactionUseCase - UserReactions - unable to get: %w", err)
	}

	return reaction, nil
}

// DeleteUserReaction -.
func (uc *UserReactionUseCase) DeleteUserReaction(ctx context.Context, userID int, postID int) error {
	reaction, err := uc.GetUserReaction(ctx, userID, postID)
	if err != nil {
		return fmt.Errorf("UserReactionUseCase - UserReactions - unable to delete: %w", err)
	}

	err = uc.reactionsRepo.Delete(ctx, reaction.ID)
	if err != nil {
		return fmt.Errorf("UserReactionUseCase - UserReactions - unable to delete: %w", err)
	}

	return nil
}
