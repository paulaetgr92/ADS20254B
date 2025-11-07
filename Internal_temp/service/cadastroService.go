package service

import (
	"awesomeProject/Internal_temp/model"
	Repository "awesomeProject/Internal_temp/repository"
	db "awesomeProject/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type CadastroService struct {
	Repo        Repository.CadastroRepositoryInterface
	R           Repository.SellersRepositoryInterface
	Twilio      *TwilioService
	Repos       *Repository.ActivationNewRepository
	CountryCode string
}

func NewCadastroService(
	cadastroRepo Repository.CadastroRepositoryInterface,
	activationRepo *Repository.ActivationNewRepository,
) *CadastroService {
	countryCode := os.Getenv("DEFAULT_COUNTRY_CODE")

	twilioService := NewTwilioService()
	return &CadastroService{
		Repo:        cadastroRepo,
		Twilio:      twilioService,
		Repos:       activationRepo,
		CountryCode: countryCode,
	}
}

func (s *CadastroService) CreateCadastro(ctx context.Context, data model.CadastroRequest) (db.Cadastro, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.Cadastro{}, fmt.Errorf("erro ao gerar hash da senha: %w", err)
	}

	// Cria o cadastro no banco
	arg := db.CreateCadastroParams{
		Name: data.Name,
		Cpf: sql.NullString{
			String: data.CPF,
			Valid:  data.CPF != "",
		},
		Cnpj: sql.NullString{
			String: data.CNPJ,
			Valid:  data.CNPJ != "",
		},
		Email:    data.Email,
		Celular:  data.Celular,
		Password: string(hashedPassword),
		Status:   "pendente",
	}

	cadastro, err := s.Repo.CreateCadastroRepository(ctx, arg)
	if err != nil {
		return db.Cadastro{}, fmt.Errorf("erro ao criar cadastro: %w", err)
	}

	codeSID, err := s.Twilio.SendActivationCode(cadastro.Celular)
	if err != nil {
		return cadastro, fmt.Errorf("erro ao enviar SMS com código de ativação: %w", err)
	}

	params := db.SaveActivationCodeParams{
		CadastroID:     cadastro.ID,
		ActivationCode: codeSID, // vem do Twilio
	}

	if _, err := s.Repos.SaveActivationCode(ctx, params); err != nil {
		return cadastro, fmt.Errorf("erro ao salvar código de ativação: %w", err)
	}

	return cadastro, nil
}
