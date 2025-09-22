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
	fromWhatsApp := os.Getenv("TWILIO_WHATSAPP_FROM")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return &TwilioService{
		client: client,
		from:   fromWhatsApp,
	}
}

func (t *TwilioService) SendActivationCode(to string, code string) error {
	if len(to) > 0 && to[:9] != "whatsapp:" {
		if to[0] != '+' {
			to = "+55" + to
		}
		to = "whatsapp:" + to
	}

	params := &openapi.CreateMessageParams{}
	params.SetFrom(t.from)
	params.SetTo(to)
	params.SetBody(fmt.Sprintf("Seu código de ativação é: %s", code))

	_, err := t.client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("erro ao enviar código de ativação: %w", err)
	}

	return nil
}
