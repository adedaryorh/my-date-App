package posts

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"backend.app/internal/controller/response"
	"backend.app/internal/dtos"
	"backend.app/internal/mappers"
	"backend.app/internal/models"
	"backend.app/internal/services/file"
	"backend.app/internal/usecase"
	"backend.app/internal/validators"
	"github.com/gofrs/uuid"
)

type Post interface {
	Create(ctx context.Context, post dtos.NewPost) (*dtos.Post, error)
	Comment(ctx context.Context, postID int, comment dtos.NewPost) (*dtos.Post, error)
	Delete(ctx context.Context, userID int, postId string) error
	Report(ctx context.Context, userID int, postID string) error
	GetAll(ctx context.Context, userID int, page *int, limit *int) (*dtos.PagedPosts, error)
	GetFeed(ctx context.Context, userID int, page *int, limit *int) (*dtos.PagedPosts, error)
	GetComments(ctx context.Context, postID string, page *int, limit *int) (*dtos.Post, *dtos.PagedPosts, error)
	Get(ctx context.Context, postId string) (*dtos.Post, error)
	ValidatePost(ctx context.Context, postId string) (*models.Post, error)
	AddReaction(ctx context.Context, postID string, reaction dtos.NewReaction) (*dtos.Post, error)
	DeleteReaction(ctx context.Context, userID int, postID string) error
}

type PostService struct {
	postsUc         usecase.Post
	reactions       usecase.UserReaction
	relationshipsUc usecase.Relationship
	uploadClient    file.UploadFileClient
	mapper          mappers.PostMapper
}

func NewPostService(pc usecase.Post, r usecase.UserReaction, rel usecase.Relationship, uploadClient file.UploadFileClient, mapper mappers.PostMapper) *PostService {
	return &PostService{postsUc: pc, reactions: r, relationshipsUc: rel, uploadClient: uploadClient, mapper: mapper}
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
		url, err := ps.uploadPostMedia(ctx, media)

		postsMedia = append(postsMedia, models.PostMedia{
			Source: url,
		})

		if err != nil {
			return nil, fmt.Errorf("user service - error updating user: %w", err)
		}
	}

	post.Media = postsMedia

	err = ps.postsUc.Create(ctx, post)
	if err != nil {
		return nil, fmt.Errorf("unable to create a post: %w", err)
	}

	postDto := ps.mapper.MapToPostDto(*post)

	return &postDto, nil
}

// Delete .
func (ps *PostService) Delete(ctx context.Context, userID int, postID string) error {
	err := ps.postsUc.Delete(ctx, userID, postID)

	if err != nil {
		return err
	}

	return nil
}

// Report .
func (ps *PostService) Report(ctx context.Context, userID int, postID string) error {
	post, err := ps.postsUc.Get(ctx, postID)
	if err != nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get post: %w", err),
			StatusCode: http.StatusNotFound,
		}
	}

	//TODO: Check that user ID cannot flag their post

	err = ps.postsUc.FlagPost(ctx, post)
	if err != nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to flag post: %w", err),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}

func (ps *PostService) GetAll(ctx context.Context, userID int, page *int, limit *int) (*dtos.PagedPosts, error) {
	posts, err := ps.postsUc.GetUserPosts(ctx, userID, page, limit)

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
	post, err := ps.postsUc.Get(ctx, postID)
	if err != nil {
		return nil, nil, &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get post: %w", err),
			StatusCode: http.StatusNotFound,
		}
	}

	postDto := ps.mapper.MapToPostDto(*post)

	posts, err := ps.postsUc.GetPostComments(ctx, post.ID, page, limit)
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
	post, err := ps.postsUc.Get(ctx, postID)

	if err != nil {
		return nil, fmt.Errorf("unable to get post: %w", err)
	}

	dto := ps.mapper.MapToPostDto(*post)

	return &dto, nil
}

func (ps *PostService) ValidatePost(ctx context.Context, postID string) (*models.Post, error) {
	post, err := ps.postsUc.Get(ctx, postID)

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
		url, err := ps.uploadPostMedia(ctx, media)

		postsMedia = append(postsMedia, models.PostMedia{
			Source: url,
		})

		if err != nil {
			return nil, fmt.Errorf("user service - error updating user: %w", err)
		}
	}

	post.Media = postsMedia

	err = ps.postsUc.Create(ctx, post)
	if err != nil {
		return nil, fmt.Errorf("unable to create a post: %w", err)
	}

	postDto := ps.mapper.MapToPostDto(*post)

	return &postDto, nil
}

func (ps *PostService) uploadPostMedia(ctx context.Context, imageStr string) (string, error) {

	// We only accept PNG, JPEG and JPG for profile images for now...
	contentType, err := validators.ValidateBase64Extension(imageStr, []string{"image/png", "image/jpeg", "image/jpg"})
	if err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}

	profileImgB64 := imageStr[strings.IndexByte(imageStr, ',')+1:]
	imgBytes, _ := base64.StdEncoding.DecodeString(profileImgB64)
	imgFile := bytes.NewReader(imgBytes)

	fileName, err := generateRandomID()
	if err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}

	ext := "png"
	if contentType == "image/jpeg" {
		ext = "jpeg"
	} else if contentType == "image/jpg" {
		ext = "jpg"
	}

	fileName = fmt.Sprintf("images/posts/%s.%s", fileName, ext)
	postImageURL, err := ps.uploadClient.UploadToBucket(ctx, "celebut", fileName, imgFile)

	if err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}

	return *postImageURL, nil
}

func (ps *PostService) AddReaction(ctx context.Context, postID string, r dtos.NewReaction) (*dtos.Post, error) {
	post, err := ps.ValidatePost(ctx, postID)
	if err != nil {
		return nil, &response.ServiceErrorResponse{
			Err:        err,
			StatusCode: 404,
		}
	}

	postDto := ps.mapper.MapToPostDto(*post)

	// Validate user has never created a reaction
	existingReaction, err := ps.reactions.GetUserReaction(ctx, r.User.ID, post.ID)
	if err != nil {
		return nil, &response.ServiceErrorResponse{
			Err:        err,
			StatusCode: 500,
		}
	}

	// Update if reactions exist instead
	if existingReaction != nil {
		existingReaction.Reaction = r.Reaction

		err = ps.reactions.Update(ctx, existingReaction)
		if err != nil {
			return nil, &response.ServiceErrorResponse{
				Err:        err,
				StatusCode: 500,
			}
		}

		return &postDto, nil
	}

	// Create reactions if not
	reaction := &models.UserReaction{
		UserID:   r.User.ID,
		PostID:   post.ID,
		Reaction: r.Reaction,
	}

	err = ps.reactions.Create(ctx, reaction)
	if err != nil {
		return nil, &response.ServiceErrorResponse{
			Err:        err,
			StatusCode: 500,
		}
	}

	return &postDto, nil
}

func (ps *PostService) DeleteReaction(ctx context.Context, userID int, postID string) error {
	post, err := ps.ValidatePost(ctx, postID)
	if err != nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to validate post: %w", err),
			StatusCode: 404,
		}
	}

	err = ps.reactions.DeleteUserReaction(ctx, userID, post.ID)
	if err != nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to delete reaction: %w", err),
			StatusCode: 404,
		}
	}

	return nil
}

func (ps *PostService) GetFeed(ctx context.Context, userID int, page *int, limit *int) (*dtos.PagedPosts, error) {
	// Get user's posts + friend's posts
	relationships, err := ps.relationshipsUc.GetUserRelationships(ctx, userID)
	if err != nil {
		return nil, &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get user's feed: %w", err),
			StatusCode: 400,
		}
	}

	friendIDs := make([]int, 0)
	for _, relationship := range relationships {
		if relationship.SenderUserID != userID {
			friendIDs = append(friendIDs, relationship.SenderUserID)
		} else {
			friendIDs = append(friendIDs, relationship.ReceiverUserID)
		}
	}

	friendIDs = append(friendIDs, userID)
	// Get user's relationships
	// TODO: Leverage Redis
	posts, err := ps.postsUc.GetMultiUserPosts(ctx, friendIDs, page, limit)

	if err != nil {
		return nil, &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get user's feed: %w", err),
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

func generateRandomID() (string, error) {
	newUUID, err := uuid.DefaultGenerator.NewV4()

	if err != nil {
		return "", fmt.Errorf("unable to generate UUID: %s", err)
	}

	fileName := strings.ReplaceAll(newUUID.String(), "-", "")

	return fileName, nil
}
