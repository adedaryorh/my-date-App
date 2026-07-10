package mappers

import (
	"backend.app/internal/dtos"
	"backend.app/internal/models"
)

type UserReactionMapper interface {
	MapToUserReactionDto(reaction models.UserReaction) dtos.UserReaction
	MapToUserReactionListDto([]models.UserReaction) []dtos.UserReaction
}

type DtoUserReactionMapper struct {
}

func (um *DtoUserReactionMapper) MapToUserReactionListDto(posts []models.UserReaction) []dtos.UserReaction {
	reactionDtos := make([]dtos.UserReaction, 0)

	for _, post := range posts {
		reactionDtos = append(reactionDtos, um.MapToUserReactionDto(post))
	}

	return reactionDtos
}

func (um *DtoUserReactionMapper) MapToUserReactionDto(r models.UserReaction) dtos.UserReaction {
	return dtos.UserReaction{
		Reaction: r.Reaction,
	}
}
