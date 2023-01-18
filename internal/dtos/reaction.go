package dtos

import (
	"celebut-api/internal/models"
)

type NewReaction struct {
	Reaction *string
	Author   models.User
}
