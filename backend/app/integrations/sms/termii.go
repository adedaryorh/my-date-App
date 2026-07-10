package sms

import (
	"fmt"
	"os"
	"path/filepath"

	"backend.app/common/helpers"
	"backend.app/configs"
	"backend.app/internal/models"
	"backend.app/pkg/rest"
)

type Termii struct {
	config *configs.Config
}

func NewTermii(config *configs.Config) PhoneService {
	termii := &Termii{config: config}
	return PhoneService(termii)
}
func (t *Termii) Send(job models.NotificationJob, channel string) error {
	message, err := getTextFromTemplate(job)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"from":    "Bamboo",
		"sms":     *message,
		"type":    "plain",
		"channel": "dnd",
		//"api_key": t.config.TermiiApiKey,
	}

	for _, to := range job.To {
		data["to"] = to
		err = t.sendRequest(data, "/sms/send")
		if err != nil {
			return err
		}
	}

	return nil
}

func getTextFromTemplate(job models.NotificationJob) (*string, error) {

	cwd, _ := os.Getwd()

	templateFileName := filepath.Join(cwd, "templates", "phone", job.Template+".txt")

	str, err := helpers.ParseTemplate(templateFileName, job.Content)
	if err != nil {
		return nil, err
	}
	return str, nil
}

func (t *Termii) sendRequest(payload map[string]interface{}, endpoint string) error {
	rest := rest.NewRestClient("https://api.ng.termii.com/api")
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := rest.Post(endpoint, payload, headers)
	fmt.Println("response", resp)

	if err != nil {
		return err
	}
	return nil
}
