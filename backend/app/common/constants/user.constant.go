package constants

type (
	Status             string
	AccountType        string
	VerificationStatus string
	AccountStatus      string
	IndustryType       string
	CompletionState    int
)

const (
	StatusActive      Status = "active"
	StatusDeActivated Status = "deactivated"
	StatusInActive    Status = "inactive"

	AccountTypePersonal AccountType = "personal"
	AccountTypeBusiness AccountType = "business"

	VerificationStatusNotVerified VerificationStatus = "not-verified"
	VerificationStatusVerified    VerificationStatus = "verified"

	AccountStatusPendingConfirmation AccountStatus = "pending-confirmation"
	AccountStatusActive              AccountStatus = "active"
	AccountStatusDeactivated         AccountStatus = "deactivated"

	CompletionStateOne CompletionState = 1 // signed up
	CompletionStateTwo CompletionState = 2 // added profile

	IndustryTypeEntertainment         IndustryType = "entertainment"
	IndustryTypeAgriculture           IndustryType = "agriculture"
	IndustryTypeFinance               IndustryType = "finance"
	IndustryTypeMining                IndustryType = "mining"
	IndustryTypeInformationTechnology IndustryType = "information-technology"
	IndustryTypeFurniture             IndustryType = "furniture"
)

func (i IndustryType) IsValid() bool {
	switch i {
	case IndustryTypeAgriculture,
		IndustryTypeFinance,
		IndustryTypeInformationTechnology,
		IndustryTypeMining,
		IndustryTypeFurniture,
		IndustryTypeEntertainment:
		return true
	}
	return false
}
