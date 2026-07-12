package repo

import (
	"backend.app/database"
	"context"
	"strings"
	"time"

	"backend.app/common/helpers"
	"backend.app/database/postgres"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/pkg/logger"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type Repo struct {
	postgres *postgres.Postgres
	log      *logger.Logger
}

const (
	WalletsTable  = "wallets"
	BalancesTable = "balances"
	UsersTable    = "users"
	Followers     = "followers"
	Blocked       = "blocked"
	Celebrations  = "celebrations"
	Media         = "media"
)

type Operations interface {
	// balance
	CreateBalanceWithHistory(ctx context.Context, balance models.Balance) (*models.Balance, error)
	GetBalanceByField(ctx context.Context, filter map[string]interface{}) (*models.Balance, error)
	GetBalanceHistory(ctx context.Context, filter map[string]interface{}, sort string, limit int) ([]*models.BalanceHistory, error)

	// blocked
	CreateBlock(ctx context.Context, blocked *models.Blocked) (*models.Blocked, error)
	GetBlockedUserByField(ctx context.Context, filter map[string]interface{}) (*dtos.Blocked, error)
	DeleteBlocked(ctx context.Context, fields helpers.Map) error
	GetAllBlocked(ctx context.Context, user *models.User, query *dtos.APIPagingDto) (*dtos.BlockedResponse, error)

	// celebration
	CreateCelebration(ctx context.Context, celebration *models.Celebration) (*models.Celebration, error)
	GetCelebrationByField(ctx context.Context, filter map[string]interface{}) (*dtos.Celebration, error)
	DeleteCelebration(ctx context.Context, fields helpers.Map) error
	GetAllCelebrations(ctx context.Context, user *models.User, query *dtos.APIPagingDto) (*dtos.CelebrationsResponse, error)
	UpdateCelebration(ctx context.Context, Id uuid.UUID, user *models.User, fields map[string]interface{}) error

	// follower
	CreateFollower(ctx context.Context, follower *models.Follower) (*models.Follower, error)
	GetFollowerByField(ctx context.Context, filter map[string]interface{}) (*dtos.Follower, error)
	DeleteFollower(ctx context.Context, fields helpers.Map) error
	GetAllFollowers(ctx context.Context, user *models.User, query *dtos.APIPagingDto) (*dtos.FollowersResponse, error)
	GetAllFriends(ctx context.Context, user *models.User, query *dtos.APIPagingDto) (*dtos.FollowersResponse, error)

	// media
	CreateMedia(ctx context.Context, media *models.Media) (*models.Media, error)
	DeleteMedia(ctx context.Context, fields helpers.Map) error

	// notification
	CreateNotification(ctx context.Context, notification *models.Notification) (*models.Notification, error)
	GetNotificationById(ctx context.Context, notificationId uuid.UUID) (*models.Notification, error)
	GetAllNotifications(ctx context.Context, query *dtos.APIPagingDto) (*dtos.NotificationsResponse, error)
	UpdateNotification(ctx context.Context, Id uuid.UUID, fields map[string]interface{}) error
	DeleteNotification(ctx context.Context, notificationId uuid.UUID) error
	GetSingleNotification(ctx context.Context, filter map[string]interface{}) (*models.Notification, error)

	// user
	CreateUser(ctx context.Context, c *models.User) (*models.User, error)
	GetUserByField(ctx context.Context, filter map[string]interface{}) (*models.User, error)
	UpdateUser(ctx context.Context, Id uuid.UUID, fields map[string]interface{}) error
	GetAllUsers(ctx context.Context, user *models.User, query *dtos.APIPagingDto) (*dtos.UsersResponse, error)
	IncrementUserFields(ctx context.Context, Id uuid.UUID, fields []*models.Incrementor) error
	// UpdateUserRole updates user role by user_id (UUID string)
	UpdateUserRole(ctx context.Context, userID string, role string) error
	// DeleteUser deletes user by user_id (UUID string)
	DeleteUser(ctx context.Context, userID string) error

	// OAuth
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	CreateUserFromGoogle(ctx context.Context, email string, firstName string, lastName string, picture string, googleID string) (*models.User, error)

	// wallet
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
	GetWalletByField(ctx context.Context, filter map[string]interface{}) (*dtos.Wallet, error)
}

func NewRepo(db *database.DB, log *logger.Logger) Operations {
	repo := Repo{db.Postgres, log}
	op := Operations(&repo)
	return op
}

//WHy are we casting Opn

type WhereObj struct {
	field     string
	condition string
	value     interface{}
}

const (
	DEFAULTPAGE                  = 1
	DEFAULTLIMIT                 = 10
	PageDefaultSortBy            = "created_at"
	PageDefaultSortDirectionDesc = "desc"
)

func getPaginationInfo(query *dtos.APIPagingDto) (*dtos.APIPagingDto, int) {
	var offset int
	// load defaults
	if query.Page == 0 {
		query.Page = DEFAULTPAGE
	}
	if query.Limit == 0 {
		query.Limit = DEFAULTLIMIT
	}
	if query.Sort == "" {
		query.Sort = PageDefaultSortBy
	}
	if query.Direction == "" {
		query.Direction = PageDefaultSortDirectionDesc
	}

	if query.Page > 1 {
		offset = query.Limit * (query.Page - 1)
	}
	return query, offset
}

func getPagingInfo(query *dtos.APIPagingDto, count int) dtos.PagingInfo {
	var hasNextPage bool

	next := int64((query.Page * query.Limit) - count)
	if next < 0 && query.Limit > 0 {
		hasNextPage = true
	}

	pagingInfo := dtos.PagingInfo{
		TotalCount:  int64(count),
		HasNextPage: hasNextPage,
		Page:        int(query.Page),
	}

	return pagingInfo
}

func genWhere(filtered WhereObj) *WhereObj {
	switch filtered.condition {
	case "eq":
		filtered.condition = "="
	case "like":
		filtered.condition = "like"
		filtered.value = "%" + filtered.value.(string) + "%"
	case "in":
		filtered.condition = "in"
		filtered.value = strings.Split(filtered.value.(string), ",")
	case "ne":
		filtered.condition = "<>"
	case "gt":
		filtered.condition = ">"
	case "lt":
		filtered.condition = "<"
	}
	return &filtered
}

func buildWhere(builder squirrel.SelectBuilder, whereMap map[string]map[string]interface{}) squirrel.SelectBuilder {
	for condition, filter := range whereMap {
		if condition == "=" || condition == "in" {
			builder = builder.Where(squirrel.Eq(filter))
		}
		if condition == "<>" {
			builder = builder.Where(squirrel.NotEq(filter))
		}
		if condition == ">" {
			builder = builder.Where(squirrel.Gt(filter))
		}
		if condition == ">=" {
			builder = builder.Where(squirrel.GtOrEq(filter))
		}
		if condition == "<" {
			builder = builder.Where(squirrel.Lt(filter))
		}
		if condition == "<=" {
			builder = builder.Where(squirrel.LtOrEq(filter))
		}
		if condition == "like" {
			builder = builder.Where(squirrel.Like(filter))
		}

	}

	return builder
}

func getFilterFromQuery(filterValue string) map[string]map[string]interface{} {
	filtered := parseFilterEntries(filterValue)
	for i := 0; i < len(filtered); i++ {
		present := filtered[i]
		filtered[i] = genWhere(*present)
	}
	whereMap := make(map[string]map[string]interface{})
	for _, filter := range filtered {

		m, ok := whereMap[filter.condition]
		if !ok {
			whereMap[filter.condition] = map[string]interface{}{
				filter.field: filter.value,
			}
		} else {
			m[filter.field] = filter.value
			whereMap[filter.condition] = m
		}
	}

	return whereMap
}

func parseFilterEntries(filter string) []*WhereObj {

	if filter == "" {
		return nil
	}
	splitFilter := strings.Split(filter, " ")
	var allWhereObj []*WhereObj
	for i := 0; i < len(splitFilter); i++ {
		data := strings.Split(splitFilter[i], "|")
		var obj WhereObj
		if len(data) > 2 {
			obj.field = data[0]
			obj.condition = data[1]
			obj.value = data[2]

		} else {
			obj = filterConditions(splitFilter[i])
		}
		allWhereObj = append(allWhereObj, &obj)
	}

	return allWhereObj
}

func filterConditions(filter string) WhereObj {
	var obj WhereObj

	var contains bool
	switch contains {
	case strings.Contains(filter, ":"):
		splitted := strings.Split(filter, ":")
		if len(splitted) == 2 {
			if string(splitted[1][0]) == "-" {
				obj = WhereObj{
					field:     splitted[0],
					condition: "ne",
					value:     splitted[1][1:len(splitted[1])],
				}
			} else {
				obj = WhereObj{
					field:     splitted[0],
					condition: "eq",
					value:     splitted[1],
				}
			}

		}
	}

	return obj
}

func generatePager(next, prev helpers.Map) dtos.PagingInfo {
	return dtos.PagingInfo{
		NextCursor: helpers.EncodeCursor(next),
		PrevCursor: helpers.EncodeCursor(prev),
	}
}

func getPaginationOperator(pointsNext bool, sortOrder string) (string, string) {
	if pointsNext && sortOrder == "asc" {
		return ">", ""
	}
	if pointsNext && sortOrder == "desc" {
		return "<", ""
	}
	if !pointsNext && sortOrder == "asc" {
		return "<", "desc"
	}
	if !pointsNext && sortOrder == "desc" {
		return ">", "asc"
	}

	return "", ""
}

type CursorData struct {
	FirstId        string
	FirstCreatedAt time.Time
	LastId         string
	LastCreatedAt  time.Time
}

func calculatePagination(isFirstPage bool, hasPagination bool, data CursorData, pointsNext bool) dtos.PagingInfo {
	pagination := dtos.PagingInfo{}
	nextCur := helpers.Map{}
	prevCur := helpers.Map{}
	if isFirstPage {
		if hasPagination {
			// nextCur := helpers.CreateCursor(books[limit-1].ID, books[limit-1].CreatedAt, true)
			nextCur := helpers.CreateCursor(data.LastId, data.LastCreatedAt, true)
			pagination = generatePager(nextCur, nil)
		}
	} else {
		if pointsNext {
			// if pointing next, it always has prev but it might not have next
			if hasPagination {
				// nextCur = helpers.CreateCursor(books[limit-1].ID, books[limit-1].CreatedAt, true)
				nextCur = helpers.CreateCursor(data.LastId, data.LastCreatedAt, true)
			}
			//prevCur = helpers.CreateCursor(books[0].ID, books[0].CreatedAt, false)
			prevCur = helpers.CreateCursor(data.FirstId, data.FirstCreatedAt, false)
			pagination = generatePager(nextCur, prevCur)
		} else {
			// this is case of prev, there will always be nest, but prev needs to be calculated
			//nextCur = helpers.CreateCursor(books[limit-1].ID, books[limit-1].CreatedAt, true)
			nextCur = helpers.CreateCursor(data.LastId, data.LastCreatedAt, true)
			if hasPagination {
				//prevCur = helpers.CreateCursor(books[0].ID, books[0].CreatedAt, false)
				prevCur = helpers.CreateCursor(data.FirstId, data.FirstCreatedAt, false)
			}
			pagination = generatePager(nextCur, prevCur)
		}
	}
	return pagination
}
