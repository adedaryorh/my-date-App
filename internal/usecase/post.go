package usecase

import (
	"celebut-api/internal/models"
	"celebut-api/internal/repo"
	"context"
	"fmt"
)

type Post interface {
	Create(context.Context, *models.Post) error
	Get(context.Context, string) (*models.Post, error)
	Delete(context.Context, *models.Post) error
}

// PostUseCase -.
type PostUseCase struct {
	repo repo.Post
}

// NewPostUseCase -.
func NewPostUseCase(r repo.Post) *PostUseCase {
	return &PostUseCase{
		repo: r,
	}
}

//Create -.
func (uc *PostUseCase) Create(ctx context.Context, c *models.Post) error {
	err := uc.repo.Create(ctx, c)

	if err != nil {
		return fmt.Errorf("PostUseCase - Posts - s.userRepo.Create: %w", err)
	}

	return nil
}

//Get -.
func (uc *PostUseCase) Get(ctx context.Context, celebrationID string) (*models.Post, error) {
	c, err := uc.repo.Get(ctx, celebrationID)

	if err != nil {
		return nil, fmt.Errorf("PostUseCase - Posts - unable to get celebration: %w", err)
	}

	return c, nil
}

//Delete -.
func (uc *PostUseCase) Delete(ctx context.Context, c *models.Post) error {
	err := uc.repo.Delete(ctx, c)
	if err != nil {
		return fmt.Errorf("PostUseCase - Posts - unable to delete celebration: %w", err)
	}

	return nil
}
