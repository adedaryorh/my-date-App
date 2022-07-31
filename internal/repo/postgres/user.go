package postgres

import (
	"celebut-api/internal/models"
	"celebut-api/pkg/postgres"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

// UserPostgresRepo -.
type UserPostgresRepo struct {
	*postgres.Postgres
}

// NewUserRepo -.
func NewUserRepo(pg *postgres.Postgres) *UserPostgresRepo {
	return &UserPostgresRepo{pg}
}

// CreateUser -.
func (r *UserPostgresRepo) CreateUser(ctx context.Context, c *models.User) error {
	var industryId *int
	if c.Industry != nil {
		industryId = &c.Industry.ID
	}

	var dob *string
	if c.DateOfBirth != nil {
		formattedDob := c.DateOfBirth.Format("2006-01-02")
		dob = &formattedDob
	}
	sql, args, err := r.Builder.
		Insert("users").
		Columns("user_id, first_name, last_name, username, country_code, phone, email, dob, gender, relationship_status, business_name, industry_id, account_type_id, password_hash, status").
		Values(c.UserID, c.FirstName, c.LastName, c.Username, c.CountryCode, c.PhoneNumber, c.Email, dob, c.Gender, c.RelationshipStatus, c.BusinessName, industryId, c.AccountType.ID, c.Password, "enabled").
		ToSql()

	if err != nil {
		return fmt.Errorf("UserPostgresRepo - CreateUser - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserPostgresRepo - CreateUser - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetUserByField -.
func (r *UserPostgresRepo) GetUserByField(ctx context.Context, field string, value string) (*models.User, error) {
	sql, _, err := r.Builder.
		Select("user_id, first_name, last_name, username, country_code, phone, email, dob, gender, relationship_status, business_name, industry_id, account_type_id, password_hash").
		From("users").
		Where(squirrel.Eq{field: value}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("unable to build query: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, value)
	if err != nil {
		return nil, fmt.Errorf("r.Pool.Query: %w", err)
	}

	u := models.User{
		Industry:    &models.Industry{},
		AccountType: models.AccountType{},
	}

	var industryID *int
	err = row.Scan(
		&u.UserID,
		&u.FirstName,
		&u.LastName,
		&u.Username,
		&u.CountryCode,
		&u.PhoneNumber,
		&u.Email,
		&u.DateOfBirth,
		&u.Gender,
		&u.RelationshipStatus,
		&u.BusinessName,
		&industryID,
		&u.AccountType.ID,
		&u.Password)

	if industryID != nil {
		u.Industry.ID = *industryID
	}

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("row.Scan: %w", err)
	}

	return &u, nil
}
