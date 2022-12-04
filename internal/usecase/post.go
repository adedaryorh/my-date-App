package usecase

import (
	"celebut-api/internal/models"
	"celebut-api/internal/repo"
	"context"
	"fmt"
)

type Post interface {
	Create(context.Context, *models.Post) error
	GetUserPosts(ctx context.Context, userID int, page int, limit int) ([]models.Post, error)
	Get(context.Context, string) (*models.Post, error)
	Delete(ctx context.Context, userID int, postID string) error
}

// PostUseCase -.
type PostUseCase struct {
	postsRepo repo.Post
	mediaRepo repo.PostMedia
}

// NewPostUseCase -.
func NewPostUseCase(r repo.Post, m repo.PostMedia) *PostUseCase {
	return &PostUseCase{
		postsRepo: r,
		mediaRepo: m,
	}
}

//Create -.
func (uc *PostUseCase) Create(ctx context.Context, p *models.Post) error {
	err := uc.postsRepo.Create(ctx, p)

	if err != nil {
		return fmt.Errorf("PostUseCase - Posts - s.postsRepo.Create: %w", err)
	}

	if len(p.Media) == 0 {
		return nil
	}

	err = uc.mediaRepo.Create(ctx, p.ID, p.Media)
	if err != nil {
		return fmt.Errorf("PostUseCase - Posts - s.mediaRepo.Create: %w", err)
	}

	return nil
}

//Get -.
func (uc *PostUseCase) Get(ctx context.Context, celebrationID string) (*models.Post, error) {
	c, err := uc.postsRepo.Get(ctx, celebrationID)

	if err != nil {
		return nil, fmt.Errorf("PostUseCase - Posts - unable to get celebration: %w", err)
	}

	return c, nil
}

//GetUserPosts -.
func (uc *PostUseCase) GetUserPosts(ctx context.Context, userID int, page int, limit int) ([]models.Post, error) {
	offset := (limit * page) - limit
	posts, err := uc.postsRepo.GetUserPosts(ctx, userID, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("PostUseCase - Posts - unable to get posts: %w", err)
	}

	postsWithMedia := make([]models.Post, 0)
	for _, post := range posts {
		media, err := uc.mediaRepo.GetPostMedia(ctx, post.ID)
		if err != nil {
			return nil, fmt.Errorf("PostUseCase - Posts - unable to get post media: %w", err)
		}

		post.Media = media
		postsWithMedia = append(postsWithMedia, post)
	}

	return postsWithMedia, nil
}

//Delete -.
func (uc *PostUseCase) Delete(ctx context.Context, userID int, postID string) error {
	post, err := uc.postsRepo.Get(ctx, postID)
	if err != nil {
		return fmt.Errorf("PostUseCase - Posts - Delete - unable to get celebration: %w", err)
	}

	media, err := uc.mediaRepo.GetPostMedia(ctx, post.ID)

	for _, postMedia := range media {
		err = uc.mediaRepo.Delete(ctx, postMedia)

	}

	err = uc.postsRepo.Delete(ctx, userID, postID)
	if err != nil {
		return fmt.Errorf("PostUseCase - Posts - unable to delete celebration: %w", err)
	}

	return nil
}
