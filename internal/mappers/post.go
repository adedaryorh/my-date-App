package mappers

import (
	"celebut-api/internal/dtos"
	"celebut-api/internal/models"
)

type PostMapper interface {
	MapToPostDto(models.Post) dtos.Post
	MapToPostListDto([]models.Post) []dtos.Post
}

type DtoPostMapper struct {
}

func (um *DtoPostMapper) MapToPostListDto(posts []models.Post) []dtos.Post {
	postDtos := make([]dtos.Post, 0)

	for _, post := range posts {
		postDtos = append(postDtos, um.MapToPostDto(post))
	}

	return postDtos
}

func (um *DtoPostMapper) MapToPostDto(p models.Post) dtos.Post {
	return dtos.Post{
		PostID:  p.PostID,
		Message: *p.Message,
		Author: dtos.UserInfo{
			UserID:       p.User.UserID,
			FirstName:    p.User.FirstName,
			LastName:     p.User.LastName,
			Username:     p.User.Username,
			BusinessName: p.User.BusinessName,
		},
		ReactionsCount: len(p.UserReactions),
		CommentsCount:  len(p.Comments),
		Media:          um.MapToPostMediaListDto(p.Media),
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
}

func (um *DtoPostMapper) MapToPostMediaListDto(m []models.PostMedia) []dtos.PostMedia {
	pm := make([]dtos.PostMedia, 0)
	for _, media := range m {
		dto := um.MapToPostMediaDto(media)
		pm = append(pm, dto)
	}

	return pm
}

func (um *DtoPostMapper) MapToPostMediaDto(m models.PostMedia) dtos.PostMedia {
	return dtos.PostMedia{
		URL: m.Source,
	}
}
