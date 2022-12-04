package dtos

type UserRelationship struct {
}

type PagedRelationships struct {
	Page  int                `json:"page"`
	Limit int                `json:"limit"`
	Posts []UserRelationship `json:"relationships"`
}
