package repo

import (
	"celebut-api/internal/models"
	"context"
)

type (
	//Industry -.
	Industry interface {
		GetIndustries(context.Context) ([]models.Industry, error)
		CreateIndustry(context.Context, *models.Industry) error
	}

	//Client -.
	Client interface {
		GetClient(context.Context, string) (*models.Client, error)
		CreateClient(context.Context, *models.Client) error

		CreateToken(context.Context, *models.ClientToken) error
		GetClientToken(context.Context, string) (*models.ClientToken, error)
		UpdateClientToken(context.Context, *models.ClientToken) error
	}

	//User -.
	User interface {
		GetUserByField(ctx context.Context, field string, value string) (*models.User, error)
		CreateUser(context.Context, *models.User) error
		UpdateUser(context.Context, *models.User, bool) error
	}

	//UserOTP -.
	UserOTP interface {
		CreateOTP(context.Context, *models.UserOTP) error
		GetOTPByUserIDAndMode(ctx context.Context, userID int, mode string) (*models.UserOTP, error)
		UseOTP(context.Context, *models.UserOTP) error
	}

	//Post -.
	Post interface {
		Create(context.Context, *models.Post) error
		Get(context.Context, string) (*models.Post, error)
		Update(context.Context, *models.Post) error
		GetUserPosts(ctx context.Context, userID int, offset *int, limit *int) ([]models.Post, error)
		GetPostComments(ctx context.Context, postID int, offset *int, limit *int) ([]models.Post, error)
		Delete(ctx context.Context, userID int, postID string) error
	}

	//PostMedia -.
	PostMedia interface {
		Create(context.Context, int, []models.PostMedia) error
		GetPostMedia(ctx context.Context, postID int) ([]models.PostMedia, error)
		GetMedia(ctx context.Context, mediaID int) (*models.PostMedia, error)
		Delete(context.Context, models.PostMedia) error
	}

	//UserRelationship -.
	UserRelationship interface {
		Create(context.Context, *models.Relationship) error
		GetUserRelationships(ctx context.Context, userID int) ([]models.Relationship, error)
		GetRelationship(ctx context.Context, senderUserID int, receiverUserID int) (*models.Relationship, error)
		Delete(context.Context, int) error
	}

	//UserReaction -.
	UserReaction interface {
		Create(context.Context, *models.UserReaction) error
		GetPostReactions(ctx context.Context, postID int) ([]models.UserReaction, error)
		GetReaction(ctx context.Context, userID int, postID int) (*models.UserReaction, error)
		Delete(context.Context, int) error
	}
)
