package models

var RedisKeys = struct {
	DataAuthStateTokens         string
	AccessToken                 string
	RefreshToken                string
	ConfirmEmail                string
	ConfirmPhone                string
	PasswordReset               string
	ConfirmBeneficiaryEmail     string
	PasswordRetries             string
	ValidateBvn                 string
	ScholarshipProgramFavorites string
	DefaultSchoolProgram        string
	BalanceLock                 string
	GeneralCelebration          string
}{
	DataAuthStateTokens:         "data:auth:state-tokens",
	AccessToken:                 "auth:user:access:token",
	RefreshToken:                "auth:user:refresh:token",
	ConfirmEmail:                "auth:user:confirm:email:token",
	ConfirmPhone:                "auth:user:confirm:phone:token",
	PasswordReset:               "auth:user:password-reset:email:token",
	ConfirmBeneficiaryEmail:     "beneficiary:confirm:email",
	PasswordRetries:             "password:retries",
	ValidateBvn:                 "bvn:validate",
	ScholarshipProgramFavorites: "scholarship-program:favorites",
	DefaultSchoolProgram:        "default.school-program",
	BalanceLock:                 "balance:lock",
	GeneralCelebration:          "general:celebration",
}
