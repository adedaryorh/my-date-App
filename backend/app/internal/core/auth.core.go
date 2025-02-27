package core

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/common/messages"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/pkg/response"
)

// SignUpUser method that handles user sign up
func (c *Core) SignUpUser(ctx context.Context, data *dtos.UserSignUp) *dtos.ResponseObject {
	existingPhoneUser, err := c.repo.GetUserByField(ctx, helpers.Map{"phone_number": data.PhoneNumber})
	if err != nil && err != messages.ErrUserNotFound {
		return response.ServerErrorResponse(err)
	}
	// phone already confirmed
	if existingPhoneUser != nil && existingPhoneUser.CompletionState > 1 {
		return response.BadRequestResponse(errors.New("phone number already confirmed"))
	}

	// phone number previously submitted but not confirmed yet
	if existingPhoneUser != nil && existingPhoneUser.CompletionState < 2 {
		if err = c.sendConfirmPhoneToken(ctx, models.RedisKeys.ConfirmPhone, existingPhoneUser); err != nil {
			c.log.Debug(">>>>> SEND TOKEN ERROR %v", err)
			return response.ServerErrorResponse(err, "user token not successfully sent")
		}
		return response.SuccessResponse(constants.UserTokenSuccessfullySent, nil)
	}

	err = c.isUnique(ctx, data.Email, data.PhoneNumber, data.Username)
	if err != nil {
		return response.BadRequestResponse(err, constants.HttpStatus(err.Error()))
	}

	dob, _ := time.Parse(constants.DATE_LAYOUT, data.DateOfBirth)
	// build user object
	firstName, LastName := splitFullName(data.FullName)
	newUser := models.User{
		FirstName:          firstName,
		Username:           data.Username,
		Email:              data.Email,
		CountryCode:        data.CountryCode,
		PhoneNumber:        data.PhoneNumber,
		CompletionState:    int(constants.CompletionStateOne),
		DateOfBirth:        &dob,
		AccountType:        string(constants.AccountTypePersonal),
		VerificationStatus: string(constants.VerificationStatusNotVerified),
		Status:             string(constants.VerificationStatusVerified),
		PasswordHash:       helpers.Hash(data.Password),
		CreatedAt:          time.Now().UTC(),
	}

	if LastName != nil {
		newUser.LastName = LastName
	}
	if existingPhoneUser == nil {
		// create user
		if err = c.repo.CreateUser(ctx, &newUser); err != nil {
			c.log.Debug(">>>>> SEND TOKEN ERROR %v", err)
			return response.ServerErrorResponse(err, constants.UserTokenNotSuccessfullySent)
		}

		// generate and send otp
		if err = c.sendConfirmPhoneToken(ctx, models.RedisKeys.ConfirmPhone, &newUser); err != nil {
			c.log.Debug(">>>>> SEND TOKEN ERROR %v", err)
			return response.ServerErrorResponse(err, "user token not successfully sent")
		}

		return response.CreatedSuccessResponse(constants.UserTokenSuccessfullySent, nil)
	}

	// update user
	fields := c.getUpdateFields(existingPhoneUser, data)
	if len(fields) > 0 {
		if err = c.repo.UpdateUser(ctx, existingPhoneUser.ID, fields); err != nil {
			return response.ServerErrorResponse(err, constants.UserTokenNotSuccessfullySent)
		}
	}
	//generate and send otp
	if err = c.sendConfirmPhoneToken(ctx, models.RedisKeys.ConfirmPhone, existingPhoneUser); err != nil {
		c.log.Debug(">>>>> SEND TOKEN ERROR %v", err)
		return response.ServerErrorResponse(err, "user token not successfully sent")
	}

	return response.CreatedSuccessResponse(constants.UserTokenSuccessfullySent, nil)

}

func (c *Core) getUpdateFields(user *models.User, data *dtos.UserSignUp) helpers.Map {
	fields := make(helpers.Map)
	firstName, lastName := splitFullName(data.FullName)

	if user.FirstName != firstName {
		fields["first_name"] = firstName
	}
	if user.FirstName != firstName {
		fields["username"] = data.Username
	}
	if user.Email != data.Email {
		fields["email"] = data.Email
	}
	if lastName != nil && *user.LastName != *lastName {
		fields["last_name"] = *lastName
	}

	return fields

}

func (c *Core) sendConfirmPhoneToken(ctx context.Context, redisKey string, user *models.User) error {
	key := fmt.Sprintf("%s:%s", models.RedisKeys.ConfirmPhone, user.PhoneNumber)
	duration := helpers.GetDurationFromTimeString(constants.AUTH_TOKEN_TTL)
	token := c.TokenService.SetToken(ctx, key, &duration)
	content := map[string]interface{}{
		"token":      token,
		"first_name": user.FirstName,
		"phone":      user.PhoneNumber,
		"subject":    "Confirm Phone",
	}
	return c.SendNotification(ctx, user, models.NotificationTemplate.ConfirmPhone, content)
}

func (c *Core) isUnique(ctx context.Context, email, phoneNumber, username string) error {
	// ensure email is unique
	existingEmailUser, err := c.repo.GetUserByField(ctx, helpers.Map{"email": email})
	if err != nil && err != messages.ErrUserNotFound {
		return err
	}
	if existingEmailUser != nil {
		return errors.New("email-already-exists")
	}

	// ensure username is unique
	existingUsernameUser, err := c.repo.GetUserByField(ctx, helpers.Map{"username": username})
	if err != nil && err != messages.ErrUserNotFound {
		return err
	}
	if existingUsernameUser != nil {
		return errors.New("username-already-exists")
	}

	return nil
}

// SignUpBusiness method that handles business sign up
func (c *Core) SignUpBusiness(ctx context.Context, data *dtos.BusinessSignUp) *dtos.ResponseObject {
	existingPhoneUser, err := c.repo.GetUserByField(ctx, helpers.Map{"phone_number": data.PhoneNumber})
	if err != nil && err != messages.ErrUserNotFound {
		return response.ServerErrorResponse(err)
	}
	// phone already confirmed
	if existingPhoneUser != nil && existingPhoneUser.CompletionState > 1 {
		return response.BadRequestResponse(errors.New("phone number already confirmed"))
	}
	// phone number previously submitted but not confirmed
	if existingPhoneUser != nil && existingPhoneUser.CompletionState < 2 {
		if err = c.sendConfirmPhoneToken(ctx, models.RedisKeys.ConfirmPhone, existingPhoneUser); err != nil {
			c.log.Debug(">>>>> SEND TOKEN ERROR %v", err)
			return response.ServerErrorResponse(err, "business token not successfully sent")
		}
		return response.SuccessResponse(constants.UserTokenSuccessfullySent, nil)
	}

	err = c.isUnique(ctx, data.Email, data.PhoneNumber, data.Username)
	if err != nil {
		return response.BadRequestResponse(err, constants.HttpStatus(err.Error()))
	}

	// build user object
	newUser := models.User{
		FirstName:          data.BusinessName,
		Username:           data.Username,
		BusinessName:       &data.BusinessName,
		IndustryType:       helpers.PointerString(string(data.Industry)),
		Email:              data.Email,
		CountryCode:        data.CountryCode,
		PhoneNumber:        data.PhoneNumber,
		CompletionState:    int(constants.CompletionStateOne),
		AccountType:        string(constants.AccountTypeBusiness),
		VerificationStatus: string(constants.VerificationStatusNotVerified),
		Status:             string(constants.VerificationStatusVerified),
		PasswordHash:       helpers.Hash(data.Password),
		CreatedAt:          time.Now().UTC(),
	}

	if existingPhoneUser == nil {
		// create user
		err = c.repo.CreateUser(ctx, &newUser)
		if err != nil {
			return response.ServerErrorResponse(err, constants.BusinessTokenNotSuccessfullySent)
		}
		if err = c.sendConfirmPhoneToken(ctx, models.RedisKeys.ConfirmPhone, &newUser); err != nil {
			c.log.Debug(">>>>> SEND TOKEN ERROR %v", err)
			return response.ServerErrorResponse(err, "user token not successfully sent")
		}
	}
	// generate and send otp
	if err = c.sendConfirmPhoneToken(ctx, models.RedisKeys.ConfirmPhone, existingPhoneUser); err != nil {
		c.log.Debug(">>>>> SEND TOKEN ERROR %v", err)
		return response.ServerErrorResponse(err, "user token not successfully sent")
	}

	return response.CreatedSuccessResponse(constants.BusinessTokenSuccessfullySent, nil)

}

// ConfirmPhone method that handles phone number confirmation
func (c *Core) ConfirmPhone(ctx context.Context, data *dtos.ConfirmPhoneNumber) *dtos.ResponseObject {
	// ensure phone number is unique
	existingUser, err := c.repo.GetUserByField(ctx, helpers.Map{"phone_number": data.PhoneNumber})
	if err != nil && err != messages.ErrUserDoesNotExist {
		return response.ServerErrorResponse(err)
	}

	if existingUser == nil {
		return response.BadRequestResponse(messages.ErrUserDoesNotExist, constants.HttpStatusResourceNotFound)
	}
	if existingUser.CompletionState > 2 {
		return response.BadRequestResponse(fmt.Errorf("%s already confirmed", data.PhoneNumber))
	}

	//confirm user token
	valid := c.TokenService.ValidateToken(ctx, fmt.Sprintf("%s:%s", models.RedisKeys.ConfirmPhone, data.PhoneNumber), data.Token)
	if !valid {
		return response.BadRequestResponse(messages.ErrInvalidToken, constants.HttpStatusInvalidToken)
	}

	// create wallet for user
	if err = c.CreateWallet(ctx, existingUser, constants.CurrencyUSD); err != nil {
		return response.ServerErrorResponse(err)
	}

	if err = c.repo.UpdateUser(ctx, existingUser.ID, helpers.Map{
		"status":           constants.AccountStatusActive,
		"completion_state": constants.CompletionStateTwo,
		"updated_at":       time.Now().UTC(),
	}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.PhoneNumberConfirmedSuccessFully, nil)
}

// Login method that handles logs user in
func (c *Core) Login(ctx context.Context, data models.SignInDto) *dtos.ResponseObject {
	user, err := c.repo.GetUserByField(ctx, helpers.Map{"email": data.Email})
	if err != nil {
		return response.ServerErrorResponse(err)
	}
	// validate password
	if user == nil {
		return response.BadRequestResponse(messages.ErrIncorrectLogin)
	}

	if user.Status != string(constants.AccountStatusActive) {
		return response.BadRequestResponse(messages.ErrInactiveUser)
	}

	if user.NextLoginAt != nil && user.NextLoginAt.After(time.Now()) {
		return response.BadRequestResponse(errors.New("user account temporarily locked "))
	}

	//validate hashed password
	key := fmt.Sprintf("%s:%s", models.RedisKeys.PasswordRetries, user.ID.String())
	isValid := helpers.CompareHash(user.PasswordHash, data.Password)

	if !isValid {
		count := c.redisService.GetIntValue(ctx, key)
		if count >= 5 {
			// deactivate user
			err := c.setNextLogin(ctx, user.ID)
			if err != nil {
				return response.ServerErrorResponse(err)
			}
			return response.BadRequestResponse(errors.New("too many password retries, user account temporarily locked"))
		}
		count++
		// count password retries
		c.redisService.Set(ctx, key, count, helpers.GetDurationFromTimeString("15m"))
		return response.BadRequestResponse(messages.ErrIncorrectLogin)
	}
	c.redisService.Delete(ctx, key)

	// if data.FcmToken != nil && *data.FcmToken != user.Integrations.Firebase.Id {
	// 	integrationData := user.Integrations
	// 	integrationData.Firebase.Id = *data.FcmToken
	// 	err = c.repo.UpdateUserByID(ctx, user.Id, &models.User{Integrations: integrationData})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	result, err := c.generateTokens(ctx, user)
	if err != nil {
		return response.ServerErrorResponse(err)
	}

	return response.CreatedSuccessResponse(constants.LoginSuccessful, result)
}

func (c *Core) setNextLogin(ctx context.Context, userId uuid.UUID) error {
	now := time.Now().Add(helpers.GetDurationFromTimeString(constants.AUTH_TOKEN_TTL)).UTC()
	return c.repo.UpdateUser(ctx, userId, helpers.Map{
		"next_login_at": now,
	})
}

func (c *Core) deactivateUser(ctx context.Context, userId uuid.UUID) error {
	// invalidate all token for the user
	err := c.repo.UpdateUser(ctx, userId, helpers.Map{"status": constants.AccountStatusDeactivated})
	if err != nil {
		return err
	}
	return c.redisService.DeleteByPattern(ctx, fmt.Sprintf("*:%s", userId.String()))
}

// SendResetPasswordToken methods to send password reset token
func (c *Core) SendResetPasswordToken(ctx context.Context, phoneNumber string) *dtos.ResponseObject {
	user, err := c.repo.GetUserByField(ctx, helpers.Map{"phone_number": phoneNumber})
	if err != nil {
		return response.ServerErrorResponse(err)
	}
	if user.Status != string(constants.AccountStatusActive) {
		return response.BadRequestResponse(errors.New("inactive user"))
	}

	// generate and send otp
	if err = c.sendConfirmPhoneToken(ctx, models.RedisKeys.PasswordReset, user); err != nil {
		c.log.Debug(">>>>> SEND TOKEN ERROR %v", err)
		return response.ServerErrorResponse(err, "user token not successfully sent")
	}
	return response.SuccessResponse(constants.PasswordTokenSuccessfullySent, nil)
}

// ResetPassword methods to use reset  a user's password
func (c *Core) ResetPassword(ctx context.Context, data *dtos.ResetPassword) *dtos.ResponseObject {
	user, err := c.repo.GetUserByField(ctx, helpers.Map{"phone_number": data.PhoneNumber})
	if err != nil {
		return response.ServerErrorResponse(err)
	}
	valid := c.TokenService.ValidateToken(ctx, fmt.Sprintf("%s:%s", models.RedisKeys.PasswordReset, data.PhoneNumber), data.Token)
	if !valid {
		return response.BadRequestResponse(errors.New("invalid token"), constants.HttpStatusInvalidToken)
	}
	if err = c.repo.UpdateUser(ctx, user.ID, helpers.Map{"password_hash": helpers.Hash(data.Password)}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.PasswordResetSuccessful, nil)
}

func (c *Core) RefreshTokens(ctx context.Context, user *models.User) (*models.AuthenticatedUser, error) {
	return c.generateTokens(ctx, user)
}

func (c *Core) generateTokens(ctx context.Context, user *models.User) (*models.AuthenticatedUser, error) {
	// generate jwt tokens
	tokens, err := c.middleware.Jwt.CreateAuthRefreshTokens(ctx, *user)
	if err != nil {
		return nil, messages.ErrCouldNotGenerateToken
	}
	authAdmin := models.AuthenticatedUser{
		User:         *user,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	// return user with token
	return &authAdmin, nil
}

func splitFullName(fullname string) (string, *string) {
	splittedName := strings.Split(fullname, " ")
	if len(splittedName) < 2 {
		return splittedName[0], nil
	}
	return splittedName[0], &splittedName[1]
}
