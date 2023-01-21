package usecase

import (
	"celebut-api/internal/models"
	"celebut-api/internal/repo"
	"context"
	"fmt"
)

type Post interface {
	Create(context.Context, *models.Post) error
	GetUserPosts(ctx context.Context, userID int, page *int, limit *int) ([]models.Post, error)
	GetMultiUserPosts(ctx context.Context, userIDs []int, page *int, limit *int) ([]models.Post, error)
	Get(context.Context, string) (*models.Post, error)
	GetPostComments(ctx context.Context, postID int, page *int, limit *int) ([]models.Post, error)
	Delete(ctx context.Context, userID int, postID string) error
	FlagPost(ctx context.Context, post *models.Post) error
}

// PostUseCase -.
type PostUseCase struct {
	postsRepo     repo.Post
	mediaRepo     repo.PostMedia
	reactionsRepo repo.UserReaction
	usersRepo     repo.User
}

// NewPostUseCase -.
func NewPostUseCase(r repo.Post, m repo.PostMedia, ur repo.UserReaction, u repo.User) *PostUseCase {
	return &PostUseCase{
		postsRepo:     r,
		mediaRepo:     m,
		reactionsRepo: ur,
		usersRepo:     u,
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
func (uc *PostUseCase) Get(ctx context.Context, postID string) (*models.Post, error) {
	c, err := uc.postsRepo.Get(ctx, postID)

	if err != nil {
		return nil, fmt.Errorf("PostUseCase - Posts - unable to get post: %w", err)
	}

	return c, nil
}

//GetUserPosts -.
func (uc *PostUseCase) GetUserPosts(ctx context.Context, userID int, page *int, limit *int) ([]models.Post, error) {
	var offset int
	if limit != nil && page != nil {
		offset = (*limit * *page) - *limit
	}

	posts, err := uc.postsRepo.GetUserPosts(ctx, userID, &offset, limit)
	if err != nil {
		return nil, fmt.Errorf("PostUseCase - Posts - unable to get posts: %w", err)
	}

	fullPosts := make([]models.Post, 0)
	for _, post := range posts {
		if err = uc.buildFullPost(ctx, &post); err != nil {
			return nil, err
		}

		fullPosts = append(fullPosts, post)
	}

	return fullPosts, nil
}

//GetPostComments -.
func (uc *PostUseCase) GetPostComments(ctx context.Context, postID int, page *int, limit *int) ([]models.Post, error) {
	var offset int
	if limit != nil && page != nil {
		offset = (*limit * *page) - *limit
	}

	posts, err := uc.postsRepo.GetPostComments(ctx, postID, &offset, limit)
	if err != nil {
		return nil, fmt.Errorf("PostUseCase - Posts - unable to get posts: %w", err)
	}

	fullPosts := make([]models.Post, 0)
	for _, post := range posts {
		if err = uc.buildFullPost(ctx, &post); err != nil {
			return nil, err
		}

		fullPosts = append(fullPosts, post)
	}

	return fullPosts, nil
}

// Delete -.
func (uc *PostUseCase) Delete(ctx context.Context, userID int, postID string) error {
	post, err := uc.postsRepo.Get(ctx, postID)
	if err != nil {
		return fmt.Errorf("PostUseCase - Posts - Delete - unable to get post: %w", err)
	}

	media, err := uc.mediaRepo.GetPostMedia(ctx, post.ID)

	for _, postMedia := range media {
		err = uc.mediaRepo.Delete(ctx, postMedia)
	}

	err = uc.postsRepo.Delete(ctx, userID, postID)
	if err != nil {
		return fmt.Errorf("PostUseCase - Posts - unable to delete post: %w", err)
	}

	return nil
}

// FlagPost - .
func (uc *PostUseCase) FlagPost(ctx context.Context, post *models.Post) error {
	post.FlaggedCounter = post.FlaggedCounter + 1

	err := uc.postsRepo.Update(ctx, post)
	if err != nil {
		return fmt.Errorf("PostUseCase - Posts - FlagPost - unable to update: %w", err)
	}

	return nil
}

// GetMultiUserPosts -.
func (uc *PostUseCase) GetMultiUserPosts(ctx context.Context, userIDs []int, page *int, limit *int) ([]models.Post, error) {
	var offset int
	if limit != nil && page != nil {
		offset = (*limit * *page) - *limit
	}

	posts, err := uc.postsRepo.GetMultiUsersPosts(ctx, userIDs, &offset, limit)
	if err != nil {
		return nil, fmt.Errorf("PostUseCase - Posts - unable to get posts: %w", err)
	}

	fullPosts := make([]models.Post, 0)
	for _, post := range posts {
		if err = uc.buildFullPost(ctx, &post); err != nil {
			return nil, err
		}

		fullPosts = append(fullPosts, post)
	}

	return fullPosts, nil
}

// buildFullPost builds a full post object by reference
func (uc *PostUseCase) buildFullPost(ctx context.Context, post *models.Post) error {
	// TODO: Improve performance by reducing amount of DB lookups
	media, err := uc.mediaRepo.GetPostMedia(ctx, post.ID)
	if err != nil {
		return fmt.Errorf("PostUseCase - Posts - unable to get post media: %w", err)
	}

	comments, err := uc.postsRepo.GetPostComments(ctx, post.ID, nil, nil)
	if err != nil {
		return fmt.Errorf("PostUseCase - Posts - unable to get post comments: %w", err)
	}

	reactions, err := uc.reactionsRepo.GetPostReactions(ctx, post.ID)
	if err != nil {
		return fmt.Errorf("PostUseCase - Posts - unable to get post reactions: %w", err)
	}

	user, err := uc.usersRepo.GetUserByID(ctx, post.UserID)
	if err != nil || user == nil {
		return fmt.Errorf("PostUseCase - Posts - unable to get post user: %w", err)
	}

	post.Comments = comments
	post.Media = media
	post.UserReactions = reactions
	post.User = *user

	return nil
}
