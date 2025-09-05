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
	"time"
)

type CadastroService struct {
	Repo   Repository.CadastroRepositoryInterface
	R      Repository.SellerRepositoryInterface
	Twilio *TwilioService
}

func NewCadastroService(cadastroRepo Repository.CadastroRepositoryInterface, sellerRepo Repository.SellerRepositoryInterface, twilioService *TwilioService) *CadastroService {
	return &CadastroService{
		Repo:   cadastroRepo,
		R:      sellerRepo,
		Twilio: twilioService,
	}
}
func (s *CadastroService) CreateCadastro(ctx context.Context, data model.CadastroRequest) (db.Cadastro, error) {
	arg := db.CreateCadastroParams{
		Name: data.Name,
		Cnpj: sql.NullString{
			String: data.CNPJ,
			Valid:  true,
		},
		Email:    data.Email,
		Celular:  data.Celular,
		Password: data.Password,
	}

	cadastro, err := s.Repo.CreateCadastroRepository(ctx, arg)
	if err != nil {
		return cadastro, err
	}

	code := GenerateActivationCode()

	err = s.R.UpdateSellerStatus(ctx, db.UpdateCadastroStatusParams{
		ActivationCode: sql.NullString{
			String: data.ActivationCode,
			Valid:  true,
		},
		Status: data.Status,
	})
	if err != nil {
		return cadastro, fmt.Errorf("erro ao salvar código de ativação: %w", err)
	}

	err = s.Twilio.SendActivationCodeTemplate(data.Celular, code)
	if err != nil {
		return cadastro, fmt.Errorf("erro ao enviar código de ativação: %w", err)
	}

	return cadastro, nil
}

func GenerateActivationCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%04d", rand.Intn(10000))
}

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
			String: data.Code,
			Valid:  true,
		},
		Status: "inativo",
	}
	if err := s.R.UpdateSellerStatus(ctx, params); err != nil {
		return errors.New("erro ao ativar conta")
	}

	return nil
}
