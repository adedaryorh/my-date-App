package dtos

import (
	"backend.app/internal/models"
)

type NewReaction struct {
	Reaction string
	User     models.User
}

type UserReaction struct {
	Reaction string `json:"reaction"`
}
