package models

import "time"

type Client struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	ClientID  string    `json:"client_id"`
	Secret    string    `json:"secret"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClientToken struct {
	ID        uint      `json:"id"`
	Token     string    `json:"token"`
	Expiry    time.Time `json:"expiry"`
	CreatedAt time.Time `json:"created_at"`
}
