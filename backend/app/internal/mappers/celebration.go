package mappers

import (
	"backend.app/internal/dtos"
	"backend.app/internal/models"
)

type CelebrationMapper interface {
	MapToCelebrationDto(models.Celebration) dtos.Celebration
	MapToCelebrationListDto([]models.Celebration) []dtos.Celebration
}
