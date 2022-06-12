package business

import "time"

type Industry struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"  example:"entertainment"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at "`
	UpdatedAt   time.Time `json:"updated_at "`
}
