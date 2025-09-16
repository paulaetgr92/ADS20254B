package service

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioService struct {
	client      *twilio.RestClient
	from        string
	templateSID string
}

func NewTwilioService() *TwilioService {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromWhatsApp := os.Getenv("TWILIO_WHATSAPP_FROM")
	templateSID := os.Getenv("TWILIO_TEMPLATE_SID")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return &TwilioService{
		client:      client,
		from:        fromWhatsApp,
		templateSID: templateSID,
	}
}

func (t *TwilioService) SendActivationCodeTemplate(to string, code string) error {
	params := &openapi.CreateMessageParams{}
	params.SetFrom("whatsapp:" + t.from)
	params.SetTo("whatsapp:" + to)
	params.SetContentSid(t.templateSID)

	variables := fmt.Sprintf(`{"1":"%s"}`, code)
	params.SetContentVariables(variables)

	_, err := t.client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("erro ao enviar código de ativação: %w", err)
	}

	return nil
}
