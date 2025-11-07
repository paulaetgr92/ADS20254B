package service

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilioService struct {
	client     *twilio.RestClient
	from       string
	serviceSID string
}

func NewTwilioService() *TwilioService {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	serviceSID := os.Getenv("TWILIO_VERIFY_SERVICE_SID") // ⚠️ precisa estar no .env
	fromNumber := os.Getenv("TWILIO_PHONE_NUMBER")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return &TwilioService{
		client:     client,
		from:       fromNumber,
		serviceSID: serviceSID,
	}
}

func (t *TwilioService) SendActivationCode(phone string) (string, error) {
	params := &verify.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")

	resp, err := t.client.VerifyV2.CreateVerification(t.serviceSID, params)
	if err != nil {
		return "", fmt.Errorf("erro ao enviar código de ativação: %v", err)
	}

	return *resp.Sid, nil
}

func (t *TwilioService) CheckActivationCode(phone, code string) (bool, error) {
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phone)
	params.SetCode(code)

	resp, err := t.client.VerifyV2.CreateVerificationCheck(t.serviceSID, params)
	if err != nil {
		return false, fmt.Errorf("erro ao verificar código: %v", err)
	}

	return *resp.Status == "approved", nil
}
