package celebrations

import (
	"celebut-api/internal/dtos"
	"context"
)

type Celebration interface {
	Create(ctx context.Context, post dtos.NewCelebration) (*dtos.Celebration, error)
	Delete(ctx context.Context, userID int, celebrationID string) error
	Report(ctx context.Context, userID int, postID string) error
	Recent(ctx context.Context, userID int, page *int, limit *int) (*dtos.PagedCelebrations, error)
	AroundYou(ctx context.Context, userID int, page *int, limit *int) (*dtos.PagedCelebrations, error)
	LovedOnes(ctx context.Context, userID int, page *int, limit *int) (*dtos.PagedCelebrations, error)
}
