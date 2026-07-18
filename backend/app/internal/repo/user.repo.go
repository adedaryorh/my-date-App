package repo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"

	"backend.app/common/helpers"
	"backend.app/common/messages"
	"backend.app/internal/dtos"
	"backend.app/internal/mappers"
	"backend.app/internal/models"
)

// CreateUser -.
func (r *Repo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if user.Role == "" {
		user.Role = "user"
	}
	sql, args, err := r.postgres.Builder.
		Insert("users").
		Columns("first_name, last_name, username, country_code, phone_number,completion_state,verification_status ,email, date_of_birth, account_type, password_hash, status,business_name,industry_type,banned_words, role").
		Values(user.FirstName, user.LastName, user.Username, user.CountryCode, user.PhoneNumber, user.CompletionState, user.VerificationStatus, user.Email, user.DateOfBirth, user.AccountType, user.PasswordHash, user.Status, user.BusinessName, user.IndustryType, user.BannedWords, user.Role).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		r.log.Error("UserPostgresRepo - CreateUser - r.Builder: %v", err)
		return nil, errors.New("something went wrong")
	}

	row := r.postgres.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&user.ID)
	if err != nil {
		r.log.Error("UserPostgresRepo - CreateUser - r.Pool.Scan: %v", err)
		return nil, errors.New("something went wrong")
	}
	return user, nil
}

// GetUserByField -.
func (r *Repo) GetUserByField(ctx context.Context, filter map[string]interface{}) (*models.User, error) {
	sql, args, err := r.postgres.Builder.
		Select("u.id,u.first_name, u.last_name, u.username,u.email,u.country_code,u.longitude,u.latitude, u.phone_number, u.completion_state, u.ip_address, u.device_type,u.date_of_birth,u.account_type,u.interests,u.notification_preference,u.language,u.profile_image_url,u.verification_status,u.password_hash, u.status,u.business_name,u.industry_type,u.created_at,u.updated_at,next_login_at,push_notification_settings,banned_words, u.role, u.google_id").
		From("users u").
		Where(squirrel.Eq(filter)).
		ToSql()

	if err != nil {
		r.log.Error("unable to build query: %v", err)
		return nil, errors.New("something went wrong")
	}
	row := r.postgres.Pool.QueryRow(ctx, sql, args...)

	u := models.User{}
	var longitude, latitude pgtype.Float4
	err = row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Username,
		&u.Email,
		&u.CountryCode,
		&longitude,
		&latitude,
		&u.PhoneNumber,
		&u.CompletionState,
		&u.IpAddress,
		&u.DeviceType,
		&u.DateOfBirth,
		&u.AccountType,
		&u.Interests,
		&u.NotificationPreference,
		&u.Language,
		&u.ProfileImageURL,
		&u.VerificationStatus,
		&u.PasswordHash,
		&u.Status,
		&u.BusinessName,
		&u.IndustryType,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.NextLoginAt,
		&u.PushNotificationSettings,
		&u.BannedWords,
		&u.Role,
		&u.GoogleID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, messages.ErrUserNotFound
		}
		r.log.Error("row.Scan: %v", err)
		return nil, errors.New("something went wrong")
	}
	return &u, nil
}

// GetAllUsers Gets all Users with pagination from DB
func (r *Repo) GetAllUsers(ctx context.Context, user *models.User, query *dtos.APIPagingDto) (*dtos.UsersResponse, error) {
	isFirstPage := query.Cursor == ""
	pointsNext := false
	builder := r.postgres.Builder.
		Select("u.id, u.first_name, u.last_name, u.username, u.country_code, u.phone, u.email, u.dob, u.gender, u.relationship_status, u.business_name,  u.account_type_id, u.password_hash, u.status, u.role").
		From("users u").
		Join("followers f ON f.follower_id = u.id")

	whereMap := getFilterFromQuery(query.Filter)
	builder = buildWhere(builder, whereMap)

	if query.Cursor != "" {
		decodedCursor, err := helpers.DecodeCursor(query.Cursor)
		if err != nil {
			r.log.Debug("GetAllUsers: DecodeCursor error : %v", err)
			return nil, err
		}
		pointsNext = decodedCursor["points_next"] == true

		operator, order := getPaginationOperator(pointsNext, query.Direction)
		whereStr := fmt.Sprintf("(u.created_at %s ? OR (u.created_at = ? AND u.id %s ?))", operator, operator)
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
		r.log.Debug("unable to build query: %w", err)
		return nil, errors.New("something went wrong")
	}

	rows, err := r.postgres.Pool.Query(ctx, sql, args...)
	if err != nil {
		r.log.Debug("r.Pool.Query: %w", err)
		return nil, errors.New("something went wrong")
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {

		u := models.User{}
		err = rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Username,
			&u.CountryCode,
			&u.PhoneNumber,
			&u.Email,
			&u.DateOfBirth,
			&u.PasswordHash,
			&u.Status,
			&u.Role,
		)

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
		cursorData.FirstId = users[0].ID.String()
		cursorData.FirstCreatedAt = users[0].CreatedAt
		cursorData.LastId = users[query.Limit-1].ID.String()
		cursorData.LastCreatedAt = users[query.Limit-1].CreatedAt
	}

	pageInfo := calculatePagination(isFirstPage, hasPagination, cursorData, pointsNext)
	var profile mappers.DtoUserMapper
	var userProfiles []*dtos.UserProfile

	for _, user := range users {
		u := profile.MapUserProfileDto(user)
		userProfiles = append(userProfiles, u)
	}
	return &dtos.UsersResponse{
		Users: userProfiles,
		PagingInfo: dtos.PagingInfo{
			NextCursor: pageInfo.NextCursor,
			PrevCursor: pageInfo.PrevCursor,
		},
	}, nil
}

// GetUserByID -.
func (r *Repo) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	sql, _, err := r.postgres.Builder.
		Select("u.id, u.user_id, u.first_name, u.last_name, u.username, u.country_code, u.phone, u.email, u.dob, u.gender, u.relationship_status, u.business_name, u.industry_id, u.account_type_id, u.password_hash, u.status, u.role, i.name").
		From("users u").
		LeftJoin("industries i ON i.id = industry_id").
		Where(squirrel.Eq{"u.id": userID}).
		ToSql()

	if err != nil {
		r.log.Debug("unable to build query: %w", err)
		return nil, errors.New("something went wrong")
	}

	row := r.postgres.Pool.QueryRow(ctx, sql, userID)

	u := models.User{}

	var industryID *int
	var industryName *string
	err = row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Username,
		&u.CountryCode,
		&u.PhoneNumber,
		&u.Email,
		&u.DateOfBirth,
		&industryID,
		&u.PasswordHash,
		&u.Status,
		&u.Role,
		&industryName)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		r.log.Debug("row.Scan: %w", err)
		return nil, errors.New("something went wrong")
	}
	if industryID != nil {
	}
	return &u, nil
}

// UpdateUser -.
func (r *Repo) UpdateUser(ctx context.Context, Id uuid.UUID, fields map[string]interface{}) error {
	sql, args, err := r.postgres.Builder.
		Update("users").
		SetMap(squirrel.Eq(fields)).
		Where(squirrel.Eq{"id": Id}).
		ToSql()

	if err != nil {
		r.log.Debug("UserPostgresRepo - UpdateUser - r.Builder: %w", err)
		return errors.New("something went wrong")
	}

	_, err = r.postgres.Pool.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Debug("UserPostgresRepo - UpdateUser - r.Pool.Exec: %w", err)
		return errors.New("something went wrong")
	}

	return nil
}

// IncrementUserFields increases or decreases fields passed
func (r *Repo) IncrementUserFields(ctx context.Context, Id uuid.UUID, fields []*models.Incrementor) error {
	builder := r.postgres.Builder.Update("users")
	for _, column := range fields {
		expr := fmt.Sprintf("%s%s%d", column.Field, column.Operator, column.Value)
		builder = builder.Set(column.Field, squirrel.Expr(expr))
	}
	builder = builder.Where(squirrel.Eq{"id": Id})
	sql, args, err := builder.ToSql()

	if err != nil {
		r.log.Debug("IncrementUserFields - UpdateUser - r.Builder: %w", err)
		return errors.New("something went wrong")
	}

	_, err = r.postgres.Pool.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Debug("IncrementUserFields - UpdateUser - r.Pool.Exec: %w", err)
		return errors.New("something went wrong")
	}

	return nil
}

// UpdateUserRole updates a user's role by user_id
func (r *Repo) UpdateUserRole(ctx context.Context, userID string, role string) error {
	sql, args, err := r.postgres.Builder.
		Update("users").
		Set("role", role).
		Where(squirrel.Eq{"id": userID}).
		ToSql()

	if err != nil {
		r.log.Debug("UserPostgresRepo - UpdateUserRole - r.Builder: %w", err)
		return errors.New("something went wrong")
	}

	_, err = r.postgres.Pool.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Debug("UserPostgresRepo - UpdateUserRole - r.Pool.Exec: %w", err)
		return errors.New("something went wrong")
	}

	return nil
}

// DeleteUser deletes a user by user_id
func (r *Repo) DeleteUser(ctx context.Context, userID string) error {
	sql, args, err := r.postgres.Builder.
		Delete("users").
		Where(squirrel.Eq{"id": userID}).
		ToSql()

	if err != nil {
		r.log.Debug("UserPostgresRepo - DeleteUser - r.Builder: %w", err)
		return errors.New("something went wrong")
	}

	_, err = r.postgres.Pool.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Debug("UserPostgresRepo - DeleteUser - r.Pool.Exec: %w", err)
		return errors.New("something went wrong")
	}

	return nil
}

// GetUserByEmail -.
func (r *Repo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return r.GetUserByField(ctx, map[string]interface{}{"email": email})
}

// CreateUserFromGoogle -.
func (r *Repo) CreateUserFromGoogle(ctx context.Context, email string, firstName string, lastName string, picture string, googleID string) (*models.User, error) {
	// Generate username from email (part before @)
	username := strings.Split(email, "@")[0]
	// If username is empty, use a fallback
	if username == "" {
		username = "user"
	}

	user := &models.User{
		Email:           email,
		FirstName:       firstName,
		LastName:        &lastName,
		Username:        username,
		ProfileImageURL: &picture,
		GoogleID:        googleID,
	}

	// Set default role if empty
	if user.Role == "" {
		user.Role = "user"
	}

	// Insert the user into the database
	sql, args, err := r.postgres.Builder.
		Insert("users").
		Columns("first_name, last_name, username, country_code, phone_number, completion_state, verification_status, email, date_of_birth, account_type, password_hash, status, business_name, industry_type, banned_words, role, google_id").
		Values(user.FirstName, user.LastName, user.Username, user.CountryCode, user.PhoneNumber, user.CompletionState, user.VerificationStatus, user.Email, user.DateOfBirth, user.AccountType, user.PasswordHash, user.Status, user.BusinessName, user.IndustryType, user.BannedWords, user.Role, user.GoogleID).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		r.log.Error("UserPostgresRepo - CreateUserFromGoogle - r.Builder: %w", err)
		return nil, errors.New("something went wrong")
	}

	row := r.postgres.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&user.ID)
	if err != nil {
		r.log.Error("UserPostgresRepo - CreateUserFromGoogle - r.Pool.Scan: %w", err)
		return nil, errors.New("something went wrong")
	}
	return user, nil
}
