package service

import (
	"awesomeProject/Internal_temp/model"
	Repository "awesomeProject/Internal_temp/repository"
	db "awesomeProject/db/sqlc"
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
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

// Construtor
func NewCadastroService(
	cadastroRepo Repository.CadastroRepositoryInterface,
	sellerRepo Repository.SellerRepositoryInterface,
	twilioService *TwilioService,
	repos *Repository.ActivationNewRepository,
) *CadastroService {
	countryCode := os.Getenv("DEFAULT_COUNTRY_CODE") // pega da env, ex: "+55"

	return &CadastroService{
		Repo:        cadastroRepo,
		R:           sellerRepo,
		Twilio:      twilioService,
		Repos:       repos,
		CountryCode: countryCode,
	}
}

// Gera código aleatório de 4 dígitos
func GenerateActivationCode() string {
	num, err := rand.Int(rand.Reader, big.NewInt(10000))
	if err != nil {
		return "0000"
	}
	return fmt.Sprintf("%04d", num.Int64())
}

// Cria cadastro + código de ativação
func (s *CadastroService) CreateCadastro(ctx context.Context, data model.CadastroRequest) (db.Cadastro, error) {
	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.Cadastro{}, fmt.Errorf("erro ao gerar hash da senha: %w", err)
	}

	// Cria cadastro no banco
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
	}

	cadastro, err := s.Repo.CreateCadastroRepository(ctx, arg)
	if err != nil {
		return cadastro, err
	}

	// Gera código de ativação
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

// Verifica código de ativação e ativa usuário
func (s *CadastroService) VerifySeller(ctx context.Context, data model.TwillioModelRequest) error {
	seller, err := s.R.GetSellerByCNPJ(ctx, data.Celular)
	if err != nil {
		return errors.New("seller não encontrado")
	}

	if !seller.ActivationCode.Valid || data.Code != seller.ActivationCode.String {
		return errors.New("código inválido")
	}

	params := db.UpdateCadastroStatusParams{
		ActivationCode: sql.NullString{
			String: "",
			Valid:  false,
		},
		Status: "ativo",
	}
	if err := s.R.UpdateSellerStatus(ctx, params); err != nil {
		return errors.New("erro ao ativar conta")
	}

	return nil
}
