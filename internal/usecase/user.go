package usecase

import (
	"celebut-api/internal/models"
	"celebut-api/internal/repo"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

// User -.
type User interface {
	Register(context.Context, *models.User) error
	Login(context.Context, string, string, int) (*models.User, error)
	ValidateExistenceByField(context.Context, string, string) error
}

// UserUseCase -.
type UserUseCase struct {
	repo repo.User
}

// NewUserUseCase -.
func NewUserUseCase(r repo.User) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

// Register - create a new user.
func (uc *UserUseCase) Register(ctx context.Context, user *models.User) error {
	userId, err := uc.generateUserID(10)
	if err != nil {
		return fmt.Errorf("unable to generate user ID: %w", err)
	}

	user.UserID = userId

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("unable to generate password hash: %w", err)
	}

	user.Password = string(hashedPassword)

	err = uc.repo.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("unable to create user: %w", err)
	}

	return nil
}

//Login -.
func (uc *UserUseCase) Login(ctx context.Context, email string, password string, accountTypeId int) (*models.User, error) {

	user, err := uc.repo.GetUserByField(ctx, "email", email)

	if err != nil {
		return nil, fmt.Errorf("user - login - s.repo.GetUserByEmail: %w", err)
	}

	if user.AccountType.ID != accountTypeId {
		return nil, errors.New("invalid credentials: user authentication failed")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials: user authentication failed")
	}

	return user, nil
}

func (uc *UserUseCase) ValidateExistenceByField(ctx context.Context, field string, value string) error {
	user, err := uc.repo.GetUserByField(ctx, field, value)

	if err != nil {
		return err
	}

	if user != nil {
		return fmt.Errorf("%s already exists", field)
	}

	return nil
}

func (uc *UserUseCase) generateUserID(size int) (string, error) {
	// First create a slice of bytes
	b := make([]byte, size)
	// Read size number of bytes into b
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// Encode our bytes as a base64 encoded string using URLEncoding
	encoded := base64.URLEncoding.EncodeToString(b)

	user, _ := uc.repo.GetUserByField(context.Background(), "user_id", encoded)

	if user != nil {
		return uc.generateUserID(size)
	}

	return encoded, nil

}
