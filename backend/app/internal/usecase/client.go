package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"backend.app/internal/models"
	"backend.app/internal/repo"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

const tokenExpiry = 3600

type ClientUseCase interface {
	CreateClient(context.Context, *models.Client) error
	AuthenticateClient(context.Context, string, string) (*string, error)
	ValidateClientToken(context.Context, string) error
}

// Client -.
type Client struct {
	repo repo.Client
}

// NewClientUseCase -.
func NewClientUseCase(r repo.Client) *Client {
	return &Client{
		repo: r,
	}
}

// AuthenticateClient - get client by client ID.
func (uc *Client) AuthenticateClient(ctx context.Context, clientId string, secret string) (*string, error) {
	client, err := uc.repo.GetClient(ctx, clientId)

	if err != nil {
		return nil, fmt.Errorf("Client - AuthenticateClient - s.userRepo.AuthenticateClient: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(client.Secret), []byte(secret))
	if err != nil {
		return nil, errors.New("invalid credentials: client authentication failed")
	}

	tokenUuid, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("Client - AuthenticateClient - Generate UUID: %w", err)
	}

	ct := &models.ClientToken{
		Token:  tokenUuid.String(),
		Expiry: time.Now().Add(time.Second * tokenExpiry),
	}

	err = uc.repo.CreateToken(ctx, ct)
	if err != nil {
		return nil, fmt.Errorf("client - AuthenticateClient - Create Token: %w", err)
	}

	return &(ct.Token), nil
}

func (uc *Client) ValidateClientToken(ctx context.Context, token string) error {
	ct, err := uc.repo.GetClientToken(ctx, token)

	if err != nil {
		return fmt.Errorf("client - AuthenticateClient - uc.userRepo.GetClientToken: %w", err)
	}

	// Check Expiry
	if ct.Expiry.Before(time.Now()) {
		return errors.New("client token has expired")
	}

	// If time is half-close to expiry
	if time.Now().Sub(ct.Expiry) > (time.Second * (0.5 * tokenExpiry)) {
		return nil
	}

	ct.Expiry = time.Now().Add(time.Second * tokenExpiry)
	err = uc.repo.UpdateClientToken(ctx, ct)

	if err != nil {
		return err
	}

	return nil

}

// CreateClient -.
func (uc *Client) CreateClient(ctx context.Context, c *models.Client) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Secret), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("client - create - s.userRepo.CreateClient: %w", err)
	}

	c.Secret = string(hashedPassword)

	err = uc.repo.CreateClient(ctx, c)
	if err != nil {
		return fmt.Errorf("client - create - s.userRepo.CreateClient: %w", err)
	}

	return nil
}
