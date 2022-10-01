package accounts

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
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return fmt.Errorf("UserPostgresRepo - CreateUser - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&c.ID)
	if err != nil {
		return fmt.Errorf("UserPostgresRepo - CreateUser - r.Pool.Scan: %w", err)
	}

	return nil
}

// GetUserByField -.
func (r *UserPostgresRepo) GetUserByField(ctx context.Context, field string, value string) (*models.User, error) {
	sql, _, err := r.Builder.
		Select("id, user_id, first_name, last_name, username, country_code, phone, email, dob, gender, relationship_status, business_name, industry_id, account_type_id, password_hash, status").
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
		&u.ID,
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
		&u.Password,
		&u.Status)

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

// UpdateUser -.
func (r *UserPostgresRepo) UpdateUser(ctx context.Context, c *models.User, updatePassword bool) error {
	var industryId *int
	if c.Industry != nil && c.Industry.ID != 0 {
		industryId = &c.Industry.ID
	}

	var dob *string
	if c.DateOfBirth != nil {
		formattedDob := c.DateOfBirth.Format("2006-01-02")
		dob = &formattedDob
	}

	sql, args, err := r.Builder.
		Update("users").
		SetMap(squirrel.Eq{
			"first_name":          c.FirstName,
			"last_name":           c.LastName,
			"username":            c.Username,
			"country_code":        c.CountryCode,
			"phone":               c.PhoneNumber,
			"email":               c.Email,
			"dob":                 dob,
			"gender":              c.Gender,
			"relationship_status": c.RelationshipStatus,
			"business_name":       c.BusinessName,
			"industry_id":         industryId,
			"account_type_id":     c.AccountType.ID,
			"password_hash":       c.Password,
			"status":              c.Status,
		}).
		Where(squirrel.Eq{"user_id": c.UserID}).
		ToSql()

	if err != nil {
		return fmt.Errorf("UserPostgresRepo - UpdateUser - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserPostgresRepo - UpdateUser - r.Pool.Exec: %w", err)
	}

	return nil
}
