package mappers

import (
	"celebut-api/internal/dtos"
	"celebut-api/internal/models"
)

type PostMapper interface {
	MapToPostDto(user models.Post) dtos.Post
}

type DtoPostMapper struct {
}

func (um *DtoPostMapper) MapToPostDto(c models.Post) dtos.Post {
	return dtos.Post{
		Message: *c.Message,
	}
}
