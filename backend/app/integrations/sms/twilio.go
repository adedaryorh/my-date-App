package sms

import (
	"fmt"
	"strings"

	"backend.app/configs"
	"backend.app/internal/models"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

var discardErrors = []string{"20003", "21408"}

type Twilio struct {
	config *configs.Config
	client *twilio.RestClient
}

func NewTwilioService(c *configs.Config) PhoneService {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.TwilioAccountSid,
		Password: c.TwilioAuthToken,
	})
	t := Twilio{
		config: c,
		client: client,
	}
	phoneService := PhoneService(&t)
	return phoneService
}
func (t *Twilio) Send(job models.NotificationJob, channel string) error {

	if channel == models.NotificationChannels.Whatsapp {
		return t.sendWhatsapp(channel, job)
	}
	if channel == models.NotificationChannels.SMS {
		return t.sendSMS(channel, job)

	}
	return nil
}
func (t *Twilio) sendWhatsapp(entity string, job models.NotificationJob) error {
	message, err := getTextFromTemplate(job)
	if err != nil {
		return err
	}
	to := fmt.Sprintf("whatsapp:%s", job.To[0])

	params := &openapi.CreateMessageParams{
		To:                  &to,
		MessagingServiceSid: &t.config.TwilioMessagingServiceId,
		Body:                message,
	}

	_, err = t.client.Api.CreateMessage(params)

	if err != nil {
		return err
	}

	return nil
}

func (t *Twilio) sendSMS(entity string, job models.NotificationJob) error {

	message, err := getTextFromTemplate(job)
	if err != nil {
		return err
	}

	params := &openapi.CreateMessageParams{
		To:   &job.To[0],
		From: &job.From,
		Body: message,
	}

	_, err = t.client.Api.CreateMessage(params)
	if err != nil {
		// If the error is 21612, it means that the network provider does not support the from alphanumeric sender ID.
		// In this case, we will use the Twilio messaging service ID.
		if strings.Contains(err.Error(), "21612") {
			return t.sendSMSWithMessagingServiceId(entity, job)
		}
		fmt.Println(err)
		for _, discardError := range discardErrors {
			if strings.Contains(err.Error(), discardError) {
				return nil
			}
		}
		return err
	}
	return nil
}

func (t *Twilio) sendSMSWithMessagingServiceId(entity string, job models.NotificationJob) error {

	message, err := getTextFromTemplate(job)
	if err != nil {
		return err
	}

	params := &openapi.CreateMessageParams{
		To:                  &job.To[0],
		Body:                message,
		MessagingServiceSid: &t.config.TwilioMessagingServiceId,
	}

	_, err = t.client.Api.CreateMessage(params)

	if err != nil {
		// if error has one of the discard errors, we will discard it
		for _, discardError := range discardErrors {
			if strings.Contains(err.Error(), discardError) {
				return nil
			}
		}
		return err
	}

	return nil
}
