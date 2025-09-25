package service

import (
	"awesomeProject/Internal_temp/model"
	Repository "awesomeProject/Internal_temp/repository"
	db "awesomeProject/db/sqlc"
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type CadastroService struct {
	Repo        Repository.CadastroRepositoryInterface
	R           Repository.SellerRepositoryInterface
	Twilio      *TwilioService
	Repos       *Repository.ActivationNewRepository
	CountryCode string
}

func NewCadastroService(cadastroRepo Repository.CadastroRepositoryInterface, repo *Repository.SellerRepository, twilioService *TwilioService, activationRepo *Repository.ActivationNewRepository) *CadastroService {
	countryCode := os.Getenv("DEFAULT_COUNTRY_CODE")

	return &CadastroService{
		Repo:        cadastroRepo,
		Twilio:      twilioService,
		Repos:       activationRepo,
		CountryCode: countryCode,
	}
}

func GenerateActivationCode() string {
	num, err := rand.Int(rand.Reader, big.NewInt(10000))
	if err != nil {
		return "0000"
	}
	return fmt.Sprintf("%04d", num.Int64())
}

func (s *CadastroService) CreateCadastro(ctx context.Context, data model.CadastroRequest) (db.Cadastro, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.Cadastro{}, fmt.Errorf("erro ao gerar hash da senha: %w", err)
	}

	arg := db.CreateCadastroParams{
		Name: data.Name,
		Cpf: sql.NullString{
			String: data.CPF,
			Valid:  true,
		},
		Cnpj: sql.NullString{
			String: data.CNPJ,
			Valid:  true,
		},
		Email:    data.Email,
		Celular:  data.Celular,
		Password: string(hashedPassword),
		Status:   "pendente",
		ActivationCode: sql.NullString{
			String: "",
			Valid:  true,
		},
	}

	cadastro, err := s.Repo.CreateCadastroRepository(ctx, arg)
	if err != nil {
		return cadastro, err
	}

	code := GenerateActivationCode()
	expiresAt := time.Now().Add(5 * time.Minute)

	params := db.SaveActivationCodeParams{
		CadastroID: cadastro.ID,
		Code:       code,
		ExpiresAt:  expiresAt,
	}
	if err := s.Repos.SaveActivationCode(ctx, params); err != nil {
		return cadastro, fmt.Errorf("erro ao salvar código de ativação: %w", err)
	}

	to := cadastro.Celular
	code = GenerateActivationCode()

	if err := s.Twilio.SendActivationCode(to, code); err != nil {
		return cadastro, fmt.Errorf("erro ao enviar código de ativação: %w", err)
	}

	return cadastro, nil
}
