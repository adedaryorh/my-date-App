package usecase

import (
	"context"
	"fmt"

	"backend.app/internal/models"
	"backend.app/internal/repo"
	//"github.com/evrone/go-clean-template/internal/entity"
)

// Industry -.
type Industry interface {
	Create(context.Context, *models.Industry) error
	Industries(context.Context) ([]models.Industry, error)
}

// IndustryUseCase -.
type IndustryUseCase struct {
	repo repo.Industry
}

// NewIndustryUseCase -.
func NewIndustryUseCase(r repo.Industry) *IndustryUseCase {
	return &IndustryUseCase{
		repo: r,
	}
}

// Industries - get list of industries from database.
func (uc *IndustryUseCase) Industries(ctx context.Context) ([]models.Industry, error) {
	industries, err := uc.repo.GetIndustries(ctx)

	if err != nil {
		return nil, fmt.Errorf("IndustryUseCase - Industries - s.userRepo.GetIndustries: %w", err)
	}

	return industries, nil
}

// Create -.
func (uc *IndustryUseCase) Create(ctx context.Context, i *models.Industry) error {
	err := uc.repo.CreateIndustry(context.Background(), i)

	if err != nil {
		return fmt.Errorf("IndustryUseCase - Industry - s.userRepo.CreateIndustry: %w", err)
	}

	return nil
}
