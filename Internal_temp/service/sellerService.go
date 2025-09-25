package service

import (
	"awesomeProject/Internal_temp/model"
	Repository "awesomeProject/Internal_temp/repository"
	db "awesomeProject/db/sqlc"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type SellerService struct {
	Repo Repository.SellerRepositoryInterface
	AC   Repository.ActivationCodeRepositoryInterface
}

func NewSellerService(
	sellerRepo Repository.SellerRepositoryInterface,
	activationRepo Repository.ActivationCodeRepositoryInterface,
) *SellerService {
	return &SellerService{
		Repo: sellerRepo,
		AC:   activationRepo,
	}
}

func (s *SellerService) VerifySeller(ctx context.Context, data model.Seller) (string, error) {

	arg := db.GetCadastroByActivationCodeParams{
		ActivationCode: sql.NullString{
			String: strconv.FormatInt(data.ActivationCode, 10),
			Valid:  true,
		},
		ID: data.CadastroId,
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

	err = s.Repo.UpdateSellerStatus(ctx, req)
	if err != nil {
		return "", errors.New("não foi possível ativar a conta")
	}

	return "ativo", nil
}
