package dtos

type Paged[T any] struct {
	Page  *int `json:"page"`
	Limit *int `json:"limit"`
	Posts []T  `json:"data"`
}
