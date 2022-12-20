package posts

import (
	"celebut-api/internal/controller/response"
	"celebut-api/internal/dtos"
	"celebut-api/internal/mappers"
	"celebut-api/internal/models"
	"celebut-api/internal/services/file"
	"celebut-api/internal/usecase"
	"celebut-api/internal/validators"
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

type Post interface {
	Create(ctx context.Context, post dtos.NewPost) (*dtos.Post, error)
	Comment(ctx context.Context, postID int, comment dtos.NewPost) (*dtos.Post, error)
	Delete(ctx context.Context, userID int, postId string) error
	Report(ctx context.Context, userID int, postID string) error
	GetAll(ctx context.Context, userID int, page *int, limit *int) (*dtos.PagedPosts, error)
	GetComments(ctx context.Context, postID string, page *int, limit *int) (*dtos.Post, *dtos.PagedPosts, error)
	Get(ctx context.Context, postId string) (*dtos.Post, error)
	ValidatePost(ctx context.Context, postId string) (*models.Post, error)
}

type PostService struct {
	usecase      usecase.Post
	uploadClient file.UploadFileClient
	mapper       mappers.PostMapper
}

func NewPostService(pc usecase.Post, uploadClient file.UploadFileClient, mapper mappers.PostMapper) *PostService {
	return &PostService{usecase: pc, uploadClient: uploadClient, mapper: mapper}
}

func (ps *PostService) Create(ctx context.Context, p dtos.NewPost) (*dtos.Post, error) {
	postID, err := generateRandomID()
	if err != nil {
		return nil, err
	}

	post := &models.Post{
		Message: p.Message,
		PostID:  postID,
		User:    p.Author,
	}

	if p.Message == nil && len(p.Media) == 0 {
		return nil, &response.ServiceErrorResponse{
			Err:        errors.New("content/file is required for a post"),
			StatusCode: 400,
		}
	}

	postsMedia := make([]models.PostMedia, 0)
	for _, media := range p.Media {
		url, err := ps.UploadPostMedia(ctx, media)

		postsMedia = append(postsMedia, models.PostMedia{
			Source: url,
		})

		if err != nil {
			return nil, fmt.Errorf("user service - error updating user: %w", err)
		}
	}

	post.Media = postsMedia

	err = ps.usecase.Create(ctx, post)
	if err != nil {
		return nil, fmt.Errorf("unable to create a post: %w", err)
	}

	postDto := ps.mapper.MapToPostDto(*post)

	return &postDto, nil
}

// Delete .
func (ps *PostService) Delete(ctx context.Context, userID int, postID string) error {
	err := ps.usecase.Delete(ctx, userID, postID)

	if err != nil {
		return err
	}

	return nil
}

// Report .
func (ps *PostService) Report(ctx context.Context, userID int, postID string) error {
	post, err := ps.usecase.Get(ctx, postID)
	if err != nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get post: %w", err),
			StatusCode: http.StatusNotFound,
		}
	}

	//TODO: Check that user ID cannot flag their post

	err = ps.usecase.FlagPost(ctx, post)
	if err != nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to flag post: %w", err),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}

func (ps *PostService) GetAll(ctx context.Context, userID int, page *int, limit *int) (*dtos.PagedPosts, error) {
	posts, err := ps.usecase.GetUserPosts(ctx, userID, page, limit)

	if err != nil {
		return nil, &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get user's posts: %w", err),
			StatusCode: 400,
		}
	}

	pagedPosts := &dtos.PagedPosts{
		Page:  page,
		Limit: limit,
		Posts: ps.mapper.MapToPostListDto(posts),
	}

	return pagedPosts, nil
}

func (ps *PostService) GetComments(ctx context.Context, postID string, page *int, limit *int) (*dtos.Post, *dtos.PagedPosts, error) {
	post, err := ps.usecase.Get(ctx, postID)
	if err != nil {
		return nil, nil, &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get post: %w", err),
			StatusCode: http.StatusNotFound,
		}
	}

	postDto := ps.mapper.MapToPostDto(*post)

	posts, err := ps.usecase.GetPostComments(ctx, post.ID, page, limit)
	if err != nil {
		return nil, nil, &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get user's posts: %w", err),
			StatusCode: http.StatusBadRequest,
		}
	}

	pagedPosts := &dtos.PagedPosts{
		Page:  page,
		Limit: limit,
		Posts: ps.mapper.MapToPostListDto(posts),
	}

	return &postDto, pagedPosts, nil
}

func (ps *PostService) Get(ctx context.Context, postID string) (*dtos.Post, error) {
	post, err := ps.usecase.Get(ctx, postID)

	if err != nil {
		return nil, fmt.Errorf("unable to get post: %w", err)
	}

	dto := ps.mapper.MapToPostDto(*post)

	return &dto, nil
}

func (ps *PostService) ValidatePost(ctx context.Context, postID string) (*models.Post, error) {
	post, err := ps.usecase.Get(ctx, postID)

	if err != nil {
		return nil, fmt.Errorf("unable to get post: %w", err)
	}

	return post, nil
}

func (ps *PostService) Comment(ctx context.Context, parentID int, comment dtos.NewPost) (*dtos.Post, error) {
	postID, err := generateRandomID()
	if err != nil {
		return nil, err
	}

	post := &models.Post{
		Message:  comment.Message,
		PostID:   postID,
		ParentID: &parentID,
		User:     comment.Author,
	}

	if comment.Message == nil && len(comment.Media) == 0 {
		return nil, &response.ServiceErrorResponse{
			Err:        errors.New("content/file is required for a post"),
			StatusCode: 400,
		}
	}

	postsMedia := make([]models.PostMedia, 0)
	for _, media := range comment.Media {
		url, err := ps.UploadPostMedia(ctx, media)

		postsMedia = append(postsMedia, models.PostMedia{
			Source: url,
		})

		if err != nil {
			return nil, fmt.Errorf("user service - error updating user: %w", err)
		}
	}

	post.Media = postsMedia

	err = ps.usecase.Create(ctx, post)
	if err != nil {
		return nil, fmt.Errorf("unable to create a post: %w", err)
	}

	postDto := ps.mapper.MapToPostDto(*post)

	return &postDto, nil
}

func (ps *PostService) UploadPostMedia(ctx context.Context, image multipart.FileHeader) (string, error) {
	imgFile, err := image.Open()
	defer imgFile.Close()

	if err != nil {
		return "", fmt.Errorf("unable to open file: %w", err)
	}

	// We only accept PNG, JPEG and JPG for profile images for now...
	_, err = validators.ValidateFileExtension(imgFile, []string{"image/png", "image/jpeg", "image/jpg"})
	if err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}

	fileName, err := generateRandomID()
	if err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}

	ext := filepath.Ext(image.Filename)
	fileName = fmt.Sprintf("profile-images/%s.%s", fileName, ext)
	profileImageURL, err := ps.uploadClient.UploadToBucket(ctx, "celebut", fileName, imgFile)

	if err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}

	return *profileImageURL, nil
}

func generateRandomID() (string, error) {
	newUUID, err := uuid.DefaultGenerator.NewV4()

	if err != nil {
		return "", fmt.Errorf("unable to generate UUID: %s", err)
	}

	fileName := strings.ReplaceAll(newUUID.String(), "-", "")

	return fileName, nil
}
