package service

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioService struct {
	client *twilio.RestClient
	from   string
}

func NewTwilioService() *TwilioService {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromNumber := os.Getenv("TWILIO_PHONE_NUMBER")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return &TwilioService{
		client: client,
		from:   fromNumber,
	}
}

func (t *TwilioService) SendActivationCode(to, code string) error {
	params := &openapi.CreateMessageParams{}
	params.SetFrom(t.from)
	params.SetTo(to)
	params.SetBody(fmt.Sprintf("Seu código de ativação é: %s", code))

	_, err := t.client.Api.CreateMessage(params)
	return err
}
