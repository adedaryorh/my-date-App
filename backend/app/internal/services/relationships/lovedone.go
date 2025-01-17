package relationships

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"backend.app/internal/config"
	"backend.app/internal/controller/response"
	"backend.app/internal/dtos"
	"backend.app/internal/mappers"
	"backend.app/internal/models"
	"backend.app/internal/usecase"
)

type LovedOne interface {
	AddLovedOne(ctx context.Context, sessionUserID int, userID string) error
	RemoveLovedOne(ctx context.Context, sessionUserID int, userID string) error
	GetLovedOnes(ctx context.Context, sessionUserID int) (*dtos.PagedRelationships, error)
	DiscoverLovedOnes(ctx context.Context, contacts []string) ([]dtos.UserInfo, error)
}

type LovedOneService struct {
	relationship usecase.Relationship
	user         usecase.User
	userMapper   mappers.UserMapper
}

func NewLovedOneService(r usecase.Relationship, u usecase.User, m mappers.UserMapper) *LovedOneService {
	return &LovedOneService{relationship: r, user: u, userMapper: m}
}

// AddLovedOne -.
func (ls *LovedOneService) AddLovedOne(ctx context.Context, sessionUserID int, userID string) error {
	user, err := ls.user.UserByField(ctx, "user_id", userID)
	if err != nil || user == nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get user: %w", err),
			StatusCode: http.StatusNotFound,
		}
	}

	if user.AccountType.ID == config.ACCOUNT_BUSINESS_ID {
		return &response.ServiceErrorResponse{
			Message:    "business can not be added as a loved one",
			Err:        errors.New("unable to add user: user is a business"),
			StatusCode: http.StatusBadRequest,
		}
	}

	rel, _ := ls.relationship.GetRelationship(ctx, sessionUserID, user.ID)
	if rel != nil {
		return &response.ServiceErrorResponse{
			Message:    "user already added as a loved one",
			Err:        errors.New("unable to add user: user is already added"),
			StatusCode: http.StatusBadRequest,
		}
	}

	nr := &models.Relationship{
		SenderUserID:   sessionUserID,
		ReceiverUserID: user.ID,
		Status:         config.StatusEnabled,
	}

	err = ls.relationship.Create(ctx, nr)
	if err != nil {
		return fmt.Errorf("unable to create new relationship: %w", err)
	}

	return nil
}

// RemoveLovedOne -.
func (ls *LovedOneService) RemoveLovedOne(ctx context.Context, sessionUserID int, userID string) error {
	user, err := ls.user.UserByField(ctx, "user_id", userID)
	if err != nil || user == nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get user: %w", err),
			StatusCode: http.StatusNotFound,
		}
	}

	rel, err := ls.relationship.GetRelationship(ctx, sessionUserID, user.ID)
	if err != nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get relationship: %w", err),
			StatusCode: http.StatusNotFound,
		}
	}

	err = ls.relationship.Delete(ctx, *rel)
	if err != nil {
		return &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to delete relationship: %w", err),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}

// GetLovedOnes -.
func (ls *LovedOneService) GetLovedOnes(ctx context.Context, sessionUserID int) (*dtos.PagedRelationships, error) {
	rels, err := ls.relationship.GetUserRelationships(ctx, sessionUserID)
	if err != nil {
		return nil, &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get loved ones: %w", err),
			StatusCode: http.StatusNotFound,
		}
	}

	relationships := make([]dtos.UserInfo, 0)
	for _, rel := range rels {
		if rel.SenderUserID == sessionUserID {
			receiverInfo := ls.userMapper.MapToUserInfoDto(rel.Receiver)
			relationships = append(relationships, receiverInfo)

			continue
		}

		senderInfo := ls.userMapper.MapToUserInfoDto(rel.Sender)
		relationships = append(relationships, senderInfo)
	}

	pagedRelationships := &dtos.PagedRelationships{
		Users: relationships,
	}

	return pagedRelationships, nil
}

func (ls *LovedOneService) DiscoverLovedOnes(ctx context.Context, contacts []string) ([]dtos.UserInfo, error) {
	users, err := ls.user.UsersByField(ctx, "phone", contacts)
	if err != nil {
		return nil, &response.ServiceErrorResponse{
			Err:        fmt.Errorf("unable to get loved ones: %w", err),
			StatusCode: http.StatusNotFound,
		}
	}

	lo := make([]dtos.UserInfo, 0)
	for _, user := range users {
		lo = append(lo, ls.userMapper.MapToUserInfoDto(user))
	}

	return lo, nil
}
