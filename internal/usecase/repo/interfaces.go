package repo

import (
	"celebut-api/internal/models"
	"celebut-api/internal/models/business"
	"context"
)

type (
	//Industry -.
	Industry interface {
		GetIndustries(context.Context) ([]business.Industry, error)
		CreateIndustry(context.Context, *business.Industry) error
	}

	//Client -.
	Client interface {
		GetClient(context.Context, string) (*models.Client, error)
		CreateClient(context.Context, *models.Client) error

		CreateToken(context.Context, *models.ClientToken) error
		GetClientToken(context.Context, string) (*models.ClientToken, error)
		UpdateClientToken(context.Context, *models.ClientToken) error
	}
)
