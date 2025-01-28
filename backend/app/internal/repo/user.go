package repo

import (
	"context"
	"fmt"

	"backend.app/common/helpers"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

// CreateUser -.
func (r *Repo) CreateUser(ctx context.Context, c *models.User) error {
	var industryId *int
	if c.Industry != nil {
		industryId = &c.Industry.ID
	}

	var dob *string
	if c.DateOfBirth != nil {
		formattedDob := c.DateOfBirth.Format("2006-01-02")
		dob = &formattedDob
	}

	sql, args, err := r.db.Postgres.Builder.
		Insert("users").
		Columns("user_id, first_name, last_name, username, country_code, phone, email, dob, gender, relationship_status, business_name, industry_id, account_type_id, password_hash, status").
		Values(c.UserID, c.FirstName, c.LastName, c.Username, c.CountryCode, c.PhoneNumber, c.Email, dob, c.Gender, c.RelationshipStatus, industryId, c.PasswordHash, c.Status).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return fmt.Errorf("UserPostgresRepo - CreateUser - r.Builder: %w", err)
	}

	row := r.db.Postgres.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&c.ID)
	if err != nil {
		return fmt.Errorf("UserPostgresRepo - CreateUser - r.Pool.Scan: %w", err)
	}

	return nil
}

// GetUserByField -.
func (r *Repo) GetUserByField(ctx context.Context, filter map[string]interface{}) (*models.User, error) {
	sql, args, err := r.db.Postgres.Builder.
		Select("u.id,u.user_id, u.first_name, u.last_name, u.username, u.country_code, u.phone, u.email, u.dob, u.gender, u.relationship_status, u.business_name, u.industry_id, u.account_type_id, u.password_hash, u.status, i.name").
		From("users u").
		LeftJoin("industries i ON i.id = industry_id").
		Where(squirrel.Eq(filter)).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("unable to build query: %w", err)
	}

	row := r.db.Postgres.Pool.QueryRow(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("r.Pool.Query: %w", err)
	}

	u := models.User{
		Industry: &models.Industry{},
	}

	var industryID *int
	var industryName *string
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
		&industryID,
		&u.PasswordHash,
		&u.Status,
		&industryName)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("row.Scan: %w", err)
	}

	if industryID != nil {
		u.Industry.ID = *industryID
		u.Industry.Name = *industryName
	}

	return &u, nil
}

// GetAllUsers Gets all Users with pagination from DB
func (r *Repo) GetAllUsers(ctx context.Context, query *dtos.APIPagingDto) (*dtos.UsersResponse, error) {

	isFirstPage := query.Cursor == ""
	pointsNext := false

	builder := r.db.Postgres.Builder.
		Select("u.id, u.user_id, u.first_name, u.last_name, u.username, u.country_code, u.phone, u.email, u.dob, u.gender, u.relationship_status, u.business_name, u.industry_id, u.account_type_id, u.password_hash, u.status, i.name").
		From("users u").
		LeftJoin("industries i ON i.id = industry_id")

	whereMap := getFilterFromQuery(query.Filter)
	builder = buildWhere(builder, whereMap)

	if query.Cursor != "" {
		decodedCursor, err := helpers.DecodeCursor(query.Cursor)
		if err != nil {
			//
		}
		pointsNext = decodedCursor["points_next"] == true

		operator, order := getPaginationOperator(pointsNext, query.Direction)
		whereStr := fmt.Sprintf("(u.created_at %s ? OR (u.created_at = ? AND user_id %s ?))", operator, operator)
		builder = builder.Where(whereStr, decodedCursor["created_at"], decodedCursor["created_at"], decodedCursor["id"])
		if order != "" {
			query.Direction = order
		}
	}

	// add limit
	builder = builder.Limit(uint64(query.Limit + 1))
	// add sort
	builder = builder.OrderBy(query.Sort, query.Direction)

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("unable to build query: %w", err)
	}

	rows, err := r.db.Postgres.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("r.Pool.Query: %w", err)
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		var industryID *int
		var industryName *string

		u := models.User{
			Industry: &models.Industry{},
		}
		err = rows.Scan(
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
			&industryID,
			&u.PasswordHash,
			&u.Status,
			&industryName)

		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		users = append(users, &u)
	}

	hasPagination := len(users) > query.Limit
	if hasPagination {
		users = users[:query.Limit]
	}
	if !isFirstPage && !pointsNext {
		users = helpers.Reverse(users)
	}

	var cursorData CursorData
	query.Limit = len(users)
	if len(users) > 0 {
		cursorData.FirstId = users[0].UserID.String()
		cursorData.FirstCreatedAt = users[0].CreatedAt
		cursorData.LastId = users[query.Limit-1].UserID.String()
		cursorData.LastCreatedAt = users[query.Limit-1].CreatedAt
	}

	pageInfo := calculatePagination(isFirstPage, hasPagination, cursorData, pointsNext)

	return &dtos.UsersResponse{
		Users: users,
		PagingInfo: dtos.PagingInfo{
			NextCursor: pageInfo.NextCursor,
			PrevCursor: pageInfo.PrevCursor,
		},
	}, nil
}

// GetUserByID -.
func (r *Repo) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	sql, _, err := r.db.Postgres.Builder.
		Select("u.id, u.user_id, u.first_name, u.last_name, u.username, u.country_code, u.phone, u.email, u.dob, u.gender, u.relationship_status, u.business_name, u.industry_id, u.account_type_id, u.password_hash, u.status, i.name").
		From("users u").
		LeftJoin("industries i ON i.id = industry_id").
		Where(squirrel.Eq{"u.id": userID}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("unable to build query: %w", err)
	}

	row := r.db.Postgres.Pool.QueryRow(ctx, sql, userID)

	u := models.User{
		Industry: &models.Industry{},
	}

	var industryID *int
	var industryName *string
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
		&industryID,
		&u.PasswordHash,
		&u.Status,
		&industryName)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("row.Scan: %w", err)
	}

	if industryID != nil {
		u.Industry.ID = *industryID
		u.Industry.Name = *industryName
	}

	return &u, nil
}

// UpdateUser -.
func (r *Repo) UpdateUser(ctx context.Context, c *models.User, updatePassword bool) error {
	var industryId *int
	if c.Industry != nil && c.Industry.ID != 0 {
		industryId = &c.Industry.ID
	}

	var dob *string
	if c.DateOfBirth != nil {
		formattedDob := c.DateOfBirth.Format("2006-01-02")
		dob = &formattedDob
	}

	sql, args, err := r.db.Postgres.Builder.
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
			"industry_id":         industryId,
			"status":              c.Status,
		}).
		Where(squirrel.Eq{"user_id": c.UserID}).
		ToSql()

	if err != nil {
		return fmt.Errorf("UserPostgresRepo - UpdateUser - r.Builder: %w", err)
	}

	_, err = r.db.Postgres.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserPostgresRepo - UpdateUser - r.Pool.Exec: %w", err)
	}

	return nil
}
