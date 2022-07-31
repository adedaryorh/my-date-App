package mappers

import (
	"celebut-api/internal/dtos"
	"celebut-api/internal/models"
)

type CelebrationMapper interface {
	MapToCelebrationDto(user models.Celebration) dtos.Celebration
}

type DtoCelebrationMapper struct {
}

func (um *DtoCelebrationMapper) MapToCelebrationDto(c models.Celebration) dtos.Celebration {
	return dtos.Celebration{
		Message: *c.Message,
	}
}
