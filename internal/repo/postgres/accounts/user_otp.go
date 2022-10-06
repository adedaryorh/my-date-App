package accounts

import (
	"celebut-api/internal/models"
	"celebut-api/pkg/postgres"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

// UserOTPRepo -.
type UserOTPRepo struct {
	*postgres.Postgres
}

// NewUserOTPRepo -.
func NewUserOTPRepo(pg *postgres.Postgres) *UserOTPRepo {
	return &UserOTPRepo{pg}
}

// CreateOTP -.
func (r *UserOTPRepo) CreateOTP(ctx context.Context, c *models.UserOTP) error {
	sql, args, err := r.Builder.
		Insert("user_otps").
		Columns("user_id, otp, mode, used, created_at").
		Values(c.UserID, c.OTP, c.Mode, false, c.CreatedAt).
		ToSql()

	if err != nil {
		return fmt.Errorf("UserOTPRepo - CreateOTP - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserOTPRepo - CreateOTP - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *UserOTPRepo) UseOTP(ctx context.Context, o *models.UserOTP) error {
	sql, args, err := r.Builder.
		Update("user_otps").
		Set("used", true).
		Set("updated_at", time.Now().UTC()).
		Where("id = ?", o.ID).
		ToSql()

	if err != nil {
		return fmt.Errorf("UserOTPRepo - UseOTP - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserOTPRepo - UseOTP - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *UserOTPRepo) GetOTPByUserIDAndMode(ctx context.Context, userID int, mode string) (*models.UserOTP, error) {
	sql, args, err := r.Builder.
		Select("id, user_id, otp, created_at").
		From("user_otps").
		Where("user_id = ? AND mode = ? AND used = false", userID, mode).
		OrderBy("created_at DESC").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("unable to build query: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("r.Pool.Query: %w", err)
	}

	otp := models.UserOTP{}

	err = row.Scan(
		&otp.ID,
		&otp.UserID,
		&otp.OTP,
		&otp.CreatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("row.Scan: %w", err)
	}

	return &otp, nil
}
