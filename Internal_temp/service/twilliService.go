package service

import (
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioService struct {
	client *twilio.RestClient
	from   string
}

func NewTwilioService() *TwilioService {
	accountSid := "ACfc9dddaf8282a7bad0be9ac46fc22830"
	authToken := "2bb497173c660b533c0f1d527ceb97f2"
	fromWhatsApp := "+14155238886"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	return &TwilioService{
		client: client,
		from:   fromWhatsApp,
	}
}

func (t *TwilioService) SendActivationCodeTemplate(to string, code string) error {
	params := &openapi.CreateMessageParams{}
	params.SetFrom("whatsapp:" + t.from)
	params.SetTo("whatsapp:" + "+5511990177807")
	templateSID := "HXb5b62575e6e4ff6129ad7c8efe1f983e"

	variables := fmt.Sprintf(`{"1":"%s"}`, code)
	params.SetContentSid(templateSID)
	params.SetContentVariables(variables)

	_, err := t.client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("erro ao enviar código de ativação: %w", err)
	}

	return nil
}
