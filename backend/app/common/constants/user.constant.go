package constants

type (
	Status             string
	AccountType        string
	VerificationStatus string
	AccountStatus      string
	IndustryType       string
	AreaOfInterest     string
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

	AreaOfInterestNewsAndEvent        AreaOfInterest = "news-and-event"
	AreaOfInterestEntertainment       AreaOfInterest = "entertainment"
	AreaOfInterestLifestyle           AreaOfInterest = "lifestyle"
	AreaOfInterestPersonalDevelopment AreaOfInterest = "personal-development"
	AreaOfInterestHumourAndMemes      AreaOfInterest = "humour-and-memes"
	AreaOfInterestSports              AreaOfInterest = "sports"
	AreaOfInterestScience             AreaOfInterest = "science"
	AreaOfInterestHistory             AreaOfInterest = "history"
	AreaOfInterestAnimals             AreaOfInterest = "animals"
	AreaOfInterestEducation           AreaOfInterest = "education"
	AreaOfInterestTechnology          AreaOfInterest = "technology"
	AreaOfInterestProductAndBrand     AreaOfInterest = "product-and-brand"
	AreaOfInterestMarketing           AreaOfInterest = "marketing"
	AreaOfInterestScaryThings         AreaOfInterest = "scary-things"
	AreaOfInterestMovies              AreaOfInterest = "movies"
	AreaOfInterestMusic               AreaOfInterest = "music"
)

func (a AreaOfInterest) IsValid() bool {
	switch a {
	case AreaOfInterestNewsAndEvent,
		AreaOfInterestEntertainment,
		AreaOfInterestLifestyle,
		AreaOfInterestPersonalDevelopment,
		AreaOfInterestHumourAndMemes,
		AreaOfInterestScience,
		AreaOfInterestSports,
		AreaOfInterestHistory,
		AreaOfInterestAnimals,
		AreaOfInterestEducation,
		AreaOfInterestTechnology,
		AreaOfInterestProductAndBrand,
		AreaOfInterestMarketing,
		AreaOfInterestScaryThings,
		AreaOfInterestMovies,
		AreaOfInterestMusic:
		return true
	}
	return false
}

// IsValid is used to validate the industry type enum
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
