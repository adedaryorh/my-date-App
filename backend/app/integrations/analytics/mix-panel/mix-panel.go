package mixpanel

import (
	"context"
	"fmt"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/configs"
	"backend.app/internal/models"
	"backend.app/pkg/logger"
	mixpanel "github.com/mixpanel/mixpanel-go"
)

type MixPanel struct {
	config *configs.Config
	log    *logger.Logger
	client *mixpanel.ApiClient
}

func NewMixPanel(config *configs.Config, log *logger.Logger) *MixPanel {
	return &MixPanel{
		config: config,
		log:    log,
		client: mixpanel.NewApiClient(config.MixPanelProjectToken, mixpanel.ServiceAccount(config.MixPanelProjectId, config.MixPanelUsername, config.MixPanelApiSecret)),
	}
}

func (m *MixPanel) SetProfile(ctx context.Context, user *models.User) error {
	data := helpers.Map{
		"name":         fmt.Sprintf("%s %s", user.FirstName, *user.LastName),
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
		"country_code": user.CountryCode,
		"username":     user.Username,
		"account_type": user.AccountType,
		"interests":    user.Interests,
	}

	if user.AccountType == string(constants.AccountTypeBusiness) {
		data["industry"] = user.IndustryType
	}
	userProfile := mixpanel.NewPeopleProperties(user.ID.String(), data)

	// pass the profile payload in the PeopleSet function
	if err := m.client.PeopleSet(ctx,
		[]*mixpanel.PeopleProperties{
			userProfile,
		},
	); err != nil {
		m.log.Debug(">>>>> MIX PANEL SetProfile error : %v", err)
	}
	return nil
}

func (m *MixPanel) TrackEvent(ctx context.Context, user *models.User, event constants.AppEvent) error {
	if err := m.client.Track(ctx, []*mixpanel.Event{
		m.client.NewEvent(string(event), user.ID.String(), helpers.Map{
			"interests":    user.Interests,
			"country_code": user.CountryCode,
			"account_type": user.AccountType,
		}),
	}); err != nil {
		m.log.Debug(">>>>> MIX PANEL TrackEvent error : %v", err)
	}
	return nil
}
