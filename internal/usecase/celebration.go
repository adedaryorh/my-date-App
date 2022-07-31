package usecase

import (
	"celebut-api/internal/models"
	"celebut-api/internal/repo"
	"context"
	"fmt"
)

type Celebration interface {
	Create(context.Context, *models.Celebration) error
	Get(context.Context, string) (*models.Celebration, error)
	//Get(context.Context) []*models.Celebration
	Delete(context.Context, *models.Celebration) error
}

// CelebrationUseCase -.
type CelebrationUseCase struct {
	repo repo.Celebration
}

// NewCelebrationUseCase -.
func NewCelebrationUseCase(r repo.Celebration) *CelebrationUseCase {
	return &CelebrationUseCase{
		repo: r,
	}
}

// Industries - get list of industries from database.
//func (uc *IndustryUseCase) Industries(ctx context.Context) ([]models.Industry, error) {
//	industries, err := uc.repo.GetIndustries(ctx)
//
//	if err != nil {
//		return nil, fmt.Errorf("IndustryUseCase - Industries - s.repo.GetIndustries: %w", err)
//	}
//
//	return industries, nil
//}

//Create -.
func (uc *CelebrationUseCase) Create(ctx context.Context, c *models.Celebration) error {
	err := uc.repo.Create(ctx, c)

	if err != nil {
		return fmt.Errorf("CelebrationUseCase - Celebrations - s.repo.Create: %w", err)
	}

	return nil
}

//Get -.
func (uc *CelebrationUseCase) Get(ctx context.Context, celebrationID string) (*models.Celebration, error) {
	c, err := uc.repo.Get(ctx, celebrationID)

	if err != nil {
		return nil, fmt.Errorf("CelebrationUseCase - Celebrations - unable to get celebration: %w", err)
	}

	return c, nil
}

//Delete -.
func (uc *CelebrationUseCase) Delete(ctx context.Context, c *models.Celebration) error {
	err := uc.repo.Delete(ctx, c)
	if err != nil {
		return fmt.Errorf("CelebrationUseCase - Celebrations - unable to delete celebration: %w", err)
	}

	return nil
}
