package models

type InviteUserDto struct {
	Email    string `json:"email" validate:"required,email"`
	RoleSlug string `json:"role_slug" validate:"required"`
}

type SignInDto struct {
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,is_password"`
	FcmToken *string `json:"fcm_token" validate:"omitempty"`
}

type SignUpDto struct {
	FirstName    string `json:"first_name" validate:"required,min=3"`
	LastName     string `json:"last_name" validate:"required,min=3"`
	Password     string `json:"password" validate:"required,is_password"`
	Email        string `json:"email" validate:"required,email"`
	ReferralCode string `json:"referral_code" validate:"omitempty"`
}

type AuthenticatedUser struct {
	AccessToken  string `json:"accessToken"`
	User         User   `json:"user"`
	RefreshToken string `json:"refreshToken"`
}

type SetPasscodeInput struct {
	Passcode string `json:"passcode" validate:"required,min=6,max=6"`
}

type AcceptInviteDto struct {
	FirstName string `json:"firstName" validate:"required,min=3"`
	LastName  string `json:"lastName" validate:"required,min=3"`
	Password  string `json:"password" validate:"required,min=8"`
	Gender    string `json:"gender"`
}

type ResetPasswordDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,is_password"`
	Token    string `json:"token" validate:"required,len=6"`
}

type ChangePasswordDto struct {
	OldPassword string `json:"old_password" validate:"required,is_password"`
	NewPassword string `json:"new_password" validate:"required,is_password"`
}

type PasswordDto struct {
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type RefreshTokenRequest struct {
	Pin          string `json:"pin" validate:"required,len=4"`
	RefreshToken string `json:"refresh_token" validate:"required,min=10"`
}

type AuthTokens struct {
	AccessToken  string
	RefreshToken string
}

type DecodedStateToken struct {
	UserUUID string
	Code     string
}
