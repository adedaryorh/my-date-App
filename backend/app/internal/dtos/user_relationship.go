package dtos

type PagedRelationships struct {
	Page  *int       `json:"page"`
	Limit *int       `json:"limit"`
	Users []UserInfo `json:"users"`
}
