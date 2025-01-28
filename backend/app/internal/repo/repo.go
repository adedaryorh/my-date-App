package repo

import (
	"context"
	"strings"
	"time"

	"backend.app/common/helpers"
	"backend.app/database"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"github.com/Masterminds/squirrel"
)

type Repo struct {
	db *database.DB
}

type Operations interface {
	CreateUser(ctx context.Context, c *models.User) error
	GetUserByField(ctx context.Context, filter map[string]interface{}) (*models.User, error)
}

func NewRepo(db *database.DB) Operations {
	repo := Repo{db}
	op := Operations(&repo)

	return op
}

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
