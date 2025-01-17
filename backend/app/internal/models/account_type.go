package models

import "time"

type AccountType struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"  example:"admin"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
