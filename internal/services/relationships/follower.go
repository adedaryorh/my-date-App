package relationships

import (
	"celebut-api/internal/config"
	"celebut-api/internal/controller/response"
	"celebut-api/internal/dtos"
	"celebut-api/internal/mappers"
	"celebut-api/internal/models"
	"celebut-api/internal/usecase"
	"context"
	"errors"
	"fmt"
	"net/http"
)

type Customer interface {
	FollowBusiness(ctx context.Context, sessionUserID int, userID string) error
	UnfollowBusiness(ctx context.Context, sessionUserID int, userID string) error
	GetFollowers(ctx context.Context, sessionUserID int) (*dtos.PagedRelationships, error)
}

type FollowerService struct {
	relationship usecase.Relationship
	user         usecase.User
	userMapper   mappers.UserMapper
}

func NewFollowerService(r usecase.Relationship, u usecase.User, m mappers.UserMapper) *FollowerService {
	return &FollowerService{relationship: r, user: u, userMapper: m}
}

// FollowBusiness -.
func (fs *FollowerService) FollowBusiness(ctx context.Context, sessionUserID int, userID string) error {
	user, err := fs.user.UserByField(ctx, "user_id", userID)
	if err != nil || user == nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get user: %w", err),
			StatusCode: http.StatusNotFound,
		}
	}

	if user.AccountType.ID != config.ACCOUNT_BUSINESS_ID {
		return &response.ServiceErrorResponse{
			Message:    "user is not a business",
			Err:        errors.New("unable to follow business: user is not a business"),
			StatusCode: http.StatusBadRequest,
		}
	}

	rel, _ := fs.relationship.GetRelationship(ctx, sessionUserID, user.ID)
	if rel != nil {
		return &response.ServiceErrorResponse{
			Message:    "you already follow this business",
			Err:        errors.New("unable to add user: user is already added"),
			StatusCode: http.StatusBadRequest,
		}
	}

	nr := &models.Relationship{
		SenderUserID:   sessionUserID,
		ReceiverUserID: user.ID,
		Status:         config.StatusEnabled,
	}

	err = fs.relationship.Create(ctx, nr)
	if err != nil {
		return fmt.Errorf("unable to create new relationship: %w", err)
	}

	return nil
}

// UnfollowBusiness -.
func (fs *FollowerService) UnfollowBusiness(ctx context.Context, sessionUserID int, userID string) error {
	user, err := fs.user.UserByField(ctx, "user_id", userID)
	if err != nil || user == nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get user: %w", err),
			StatusCode: http.StatusNotFound,
		}
	}

	rel, err := fs.relationship.GetRelationship(ctx, sessionUserID, user.ID)
	if err != nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get relationship: %w", err),
			StatusCode: http.StatusNotFound,
		}
	}

	err = fs.relationship.Delete(ctx, *rel)
	if err != nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to delete relationship: %w", err),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}

// GetFollowers -.
func (fs *FollowerService) GetFollowers(ctx context.Context, sessionUserID int) (*dtos.PagedRelationships, error) {
	rels, err := fs.relationship.GetUserRelationships(ctx, sessionUserID)
	if err != nil {
		return nil, &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get relationship: %w", err),
			StatusCode: http.StatusNotFound,
		}
	}

	relationships := make([]dtos.UserInfo, 0)
	for _, rel := range rels {
		if rel.SenderUserID == sessionUserID {
			receiverInfo := fs.userMapper.MapToUserInfoDto(rel.Receiver)
			relationships = append(relationships, receiverInfo)

			continue
		}

		senderInfo := fs.userMapper.MapToUserInfoDto(rel.Sender)
		relationships = append(relationships, senderInfo)
	}

	pagedRelationships := &dtos.PagedRelationships{
		Users: relationships,
	}

	return pagedRelationships, nil
}
