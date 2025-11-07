package service

import (
	"awesomeProject/Internal_temp/model"
	Repository "awesomeProject/Internal_temp/repository"
	db "awesomeProject/db/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SellerService struct {
	Repo Repository.SellersRepositoryInterface
	AC   *Repository.ActivationNewRepository
}

// Novo service
func NewSellerService(
	sellerRepo Repository.SellersRepositoryInterface,
	activationRepo *Repository.ActivationNewRepository,
) *SellerService {
	return &SellerService{
		Repo: sellerRepo,
		AC:   activationRepo,
	}
}

// Gera código de ativação aleatório de 6 dígitos
func generateActivationCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := 100000 + rnd.Intn(900000)
	return strconv.Itoa(code)
}

func (s *SellerService) SendActivationCode(to string, code string) (string, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER")) // número Twilio habilitado para SMS
	params.SetBody(fmt.Sprintf("Seu código de ativação é: %s", code))

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return "", fmt.Errorf("erro ao enviar SMS: %w", err)
	}

	// ✅ Retorna o SID da mensagem (identificador único do envio)
	return *resp.Sid, nil
}

func (s *SellerService) CreateSeller(ctx context.Context, seller model.SellerRequest) (model.SellerResponse, error) {

	arg := db.CreateSellerParams{
		Name:           seller.Name,
		Email:          seller.Email,
		ActivationCode: sql.NullString{}, // será preenchido após o envio via Twilio
		Password:       seller.Password,
		Cpf: sql.NullString{
			String: seller.CPF,
			Valid:  seller.CPF != "",
		},
		Cnpj: sql.NullString{
			String: seller.CNPJ,
			Valid:  seller.CNPJ != "",
		},
		Celular: seller.Celular,
		Status:  "pendente",
	}

	// 1️⃣ Cria o vendedor no banco
	dbSeller, err := s.Repo.CreateSeller(ctx, arg)
	if err != nil {
		return model.SellerResponse{}, err
	}

	code := generateActivationCode()
	codeSID, err := s.SendActivationCode(seller.Celular, code)
	if err != nil {
		return model.SellerResponse{}, fmt.Errorf("erro ao enviar código via Twilio: %w", err)
	}

	// 3️⃣ Salva o código e o SID no banco
	params := db.SaveActivationCodeParams{
		CadastroID:     dbSeller.CadastroID,
		ActivationCode: code,
		Code:           codeSID,
		ExpiresAt:      time.Now().Add(10 * time.Minute),
		Status: sql.NullString{
			String: "pendente",
			Valid:  true,
		},
	}

	if _, err := s.AC.SaveActivationCode(ctx, params); err != nil {
		return model.SellerResponse{}, errors.New("não foi possível salvar o código de ativação")
	}

	createdSeller := model.SellerResponse{
		CadastroID:     int(dbSeller.CadastroID),
		ActivationCode: "",
	}

	return createdSeller, nil
}
