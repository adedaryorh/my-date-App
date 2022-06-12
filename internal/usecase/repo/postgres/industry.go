package postgres

import (
	"celebut-api/internal/models/business"
	"celebut-api/pkg/postgres"
	"context"
	"fmt"
)

// IndustryPostgresRepo -.
type IndustryPostgresRepo struct {
	*postgres.Postgres
}

// NewIndustryRepo -.
func NewIndustryRepo(pg *postgres.Postgres) *IndustryPostgresRepo {
	return &IndustryPostgresRepo{pg}
}

// CreateIndustry -.
func (r *IndustryPostgresRepo) CreateIndustry(ctx context.Context, i *business.Industry) error {
	sql, args, err := r.Builder.
		Insert("industries").
		Columns("name, description").
		Values(i.Name, i.Description).
		ToSql()

	if err != nil {
		return fmt.Errorf("IndustryPostgresRepo - CreateIndustry - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("IndustryPostgresRepo - CreateIndustry - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetIndustries -.
func (r *IndustryPostgresRepo) GetIndustries(ctx context.Context) ([]business.Industry, error) {
	sql, _, err := r.Builder.
		Select("id, name, description").
		From("industries").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("ClientPostgresRepo - GetIndustries - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("ClientPostgresRepo - GetIndustries - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	industries := make([]business.Industry, 0, _defaultEntityCap)

	for rows.Next() {
		i := business.Industry{}

		err = rows.Scan(&i.ID, &i.Name, &i.Description)
		if err != nil {
			return nil, fmt.Errorf("ClientPostgresRepo - GetIndustries - rows.Scan: %w", err)
		}

		industries = append(industries, i)
	}

	return industries, nil
}
