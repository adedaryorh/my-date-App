package models

import "time"

type Relationship struct {
	ID             int
	SenderUserID   int
	ReceiverUserID int
	Sender         User
	Receiver       User
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
