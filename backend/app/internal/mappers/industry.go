package mappers

import (
	"backend.app/internal/dtos"
	"backend.app/internal/models"
)

type IndustryMapper interface {
	MapToIndustryDto(industry models.Industry) dtos.Industry
	MapToIndustryListDto([]models.Industry) []dtos.Industry
}

type DtoIndustryMapper struct {
}

func (um *DtoIndustryMapper) MapToIndustryListDto(industries []models.Industry) []dtos.Industry {
	industryDtos := make([]dtos.Industry, 0)

	for _, post := range industries {
		industryDtos = append(industryDtos, um.MapToIndustryDto(post))
	}

	return industryDtos
}

func (um *DtoIndustryMapper) MapToIndustryDto(i models.Industry) dtos.Industry {
	return dtos.Industry{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
	}
}
