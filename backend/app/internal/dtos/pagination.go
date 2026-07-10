package dtos

type Paged[T any] struct {
	Page  *int `json:"page"`
	Limit *int `json:"limit"`
	Posts []T  `json:"data"`
}

type (
	APIPagingDto struct {
		Limit     int      `json:"limit,omitempty"`
		Sort      string   `json:"sort,omitempty"`
		Direction string   `json:"direction,omitempty"`
		Select    []string `json:"select,omitempty"`
		Filter    string   `json:"filter,omitempty"`
		Math      string   `json:"math,omitempty"`
		Range     string   `json:"range,omitempty"`
		Page      int      `json:"page,omitempty"`
		Cursor    string   `json:"cursor,omitempty"`
	}

	PagingInfo struct {
		TotalCount  int64  `json:"total_count"`
		Page        int    `json:"page"`
		HasNextPage bool   `json:"has_next_page"`
		Count       int    `json:"count"`
		NextCursor  string `json:"next_cursor"`
		PrevCursor  string `json:"prev_cursor"`
	}
)
