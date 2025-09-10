package main

import (
	"fmt"
	"log"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	accountSid := ""
	authToken := "" 

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetFrom("whatsapp:+14155238886")
	params.SetTo("whatsapp:+5511983007746")
	params.SetContentSid("HXb5b62575e6e4ff6129ad7c8efe1f983e")
	params.SetContentVariables(`{"1":"12/1","2":"3pm"}`)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Fatalf("Erro ao enviar mensagem: %v", err)
	}

	if resp.Sid != nil {
		fmt.Println("Mensagem enviada com sucesso! SID:", *resp.Sid)
	} else {
		fmt.Println("Mensagem enviada, mas sem SID retornado")
	}
}
