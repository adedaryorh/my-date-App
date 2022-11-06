package dtos

import "time"

type Post struct {
	Message   string    `json:"message"`
	Author    User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
