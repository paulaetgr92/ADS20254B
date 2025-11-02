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

// Envia código de ativação via Twilio
func (s *SellerService) SendActivationCode(to string, code string) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(os.Getenv("TWILIO_SMS_FROM")) // número Twilio habilitado para SMS
	params.SetBody(fmt.Sprintf("Seu código de ativação é: %s", code))

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("erro ao enviar SMS: %w", err)
	}
	return nil
}

// Cria seller, salva código de ativação e envia SMS
func (s *SellerService) CreateSeller(ctx context.Context, seller model.Seller) (model.Seller, error) {
	seller.Status = "pendente"

	// Cria seller no banco
	dbSeller, err := s.Repo.CreateSeller(ctx, seller)
	if err != nil {
		return seller, err
	}

	createdSeller := model.Seller{
		Name:       dbSeller.Name,
		Email:      dbSeller.Email,
		CadastroId: dbSeller.ID,
		Phone:      dbSeller.Celular,
		Status:     dbSeller.Status,
	}

	// Gera código de ativação
	codeStr := generateActivationCode()
	createdSeller.ActivationCodes = codeStr

	// Salva código no banco
	params := db.SaveActivationCodeParams{
		CadastroID:      createdSeller.CadastroId,
		ActivationCodes: codeStr,
		Code:            codeStr,
		ExpiresAt:       time.Now().Add(10 * time.Minute),
		Status: sql.NullString{
			String: "pendente",
			Valid:  true,
		},
	}

	if err := s.AC.SaveActivationCode(ctx, params); err != nil {
		return createdSeller, errors.New("não foi possível salvar o código de ativação")
	}

	// Envia código via SMS
	if err := s.SendActivationCode(createdSeller.Phone, codeStr); err != nil {
		return createdSeller, fmt.Errorf("erro ao enviar código de ativação: %w", err)
	}

	return createdSeller, nil
}

// Verifica código de ativação e ativa seller
func (s *SellerService) VerifySeller(ctx context.Context, cadastroId int64, code string) (string, error) {
	arg := db.GetCadastroByActivationCodeParams{
		ID: cadastroId,
		ActivationCode: sql.NullString{
			String: code,
			Valid:  true,
		},
	}

	cadastro, err := s.AC.GetActivationCode(ctx, arg)
	if err != nil {
		return "", errors.New("código de ativação inválido")
	}

	if cadastro.Status == "ativo" {
		return "", errors.New("conta já está ativa")
	}

	req := db.UpdateCadastroStatusParams{
		ID:     cadastro.ID,
		Status: "ativo",
	}

	if err := s.Repo.UpdateSellerStatus(ctx, req); err != nil {
		return "", errors.New("não foi possível ativar a conta")
	}

	return "ativo", nil
}
