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
	}

	//Celebration -.
	Celebration interface {
		Create(context.Context, *models.Celebration) error
		Get(context.Context, string) (*models.Celebration, error)
		//GetAll(context.Context, string) []models.Celebration
		Delete(context.Context, *models.Celebration) error
	}
)
