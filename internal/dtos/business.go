package dtos

type Business struct {
	UserID       string `json:"user_id"`
	BusinessName string `json:"business_name"`
	IndustryType string `json:"industry_type"`
	CountryCode  string `json:"country_code"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
}
