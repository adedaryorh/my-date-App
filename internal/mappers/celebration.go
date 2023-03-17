package mappers

import (
	"celebut-api/internal/dtos"
	"celebut-api/internal/models"
)

type CelebrationMapper interface {
	MapToCelebrationDto(models.Celebration) dtos.Celebration
	MapToCelebrationListDto([]models.Celebration) []dtos.Celebration
}
