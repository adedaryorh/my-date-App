package users

type User struct {
	ID                 uint   `json:"id"`
	UserID             string `json:"user_id"`
	FirstName          string
	LastName           string
	Username           string
	CountryCode        string
	PhoneNumber        string
	Email              string
	DateOfBirth        string
	Gender             string
	RelationshipStatus string
	Interests          string
	PasswordHash       string
	CreatedAt          string
	UpdatedAt          string
}
