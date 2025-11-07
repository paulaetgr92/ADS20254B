package service

import (
	"awesomeProject/Internal_temp/model"
	Repository "awesomeProject/Internal_temp/repository"
	db "awesomeProject/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

type ActivationService struct {
	repo *Repository.ActivationNewRepository
}

func NewActivationService(repo *Repository.ActivationNewRepository) *ActivationService {
	return &ActivationService{repo: repo}
}

func (s *ActivationService) SendActivationCode(ctx context.Context, cadastroID int64) (string, error) {
	if cadastroID == 0 {
		return "", fmt.Errorf("ID do cadastro inválido")
	}

	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	expiresAt := time.Now().Add(5 * time.Minute)

	arg := db.SaveActivationCodeParams{
		CadastroID:     cadastroID,
		ActivationCode: code,
		Code:           code,
		ExpiresAt:      expiresAt,
		Status: sql.NullString{
			String: "pendente",
			Valid:  true,
		},
	}

	_, err := s.repo.SaveActivationCode(ctx, arg)
	if err != nil {
		return "", fmt.Errorf("não foi possível salvar o código de ativação: %w", err)
	}

	return code, nil
}

func (s *ActivationService) GetActivationCode(ctx context.Context, activationCode string, cadastroID int64) (db.GetCadastroByActivationCodeRow, error) {
	if cadastroID == 0 {
		return db.GetCadastroByActivationCodeRow{}, fmt.Errorf("ID do cadastro inválido")
	}
	if activationCode == "" {
		return db.GetCadastroByActivationCodeRow{}, fmt.Errorf("código de ativação inválido")
	}

	codeRow, err := s.repo.GetActivationCode(ctx, db.GetCadastroByActivationCodeParams{
		ActivationCode: sql.NullString{
			String: activationCode,
			Valid:  true,
		},
		ID: cadastroID,
	})
	if err != nil {
		return db.GetCadastroByActivationCodeRow{}, fmt.Errorf("erro ao buscar código de ativação: %v", err)
	}

	return codeRow, nil
}

func (s *ActivationService) VerifyActivationCode(ctx context.Context, data model.ActivationCode) (db.VerifyActivationCodeRow, error) {
	arg := db.VerifyActivationCodeParams{
		CadastroID: int64(data.CadastroID),
		Code:       data.Code,
	}
	return s.repo.VerifyActivationCode(ctx, arg)
}

func (s *ActivationService) SaveActivationCode(ctx context.Context, req model.ActivationCode) (model.ActivationCode, error) {
	params := db.SaveActivationCodeParams{
		CadastroID: int64(req.CadastroID),
		Code:       req.Code,
		Status: sql.NullString{
			String: "pending",
			Valid:  true,
		},
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}

	response, err := s.repo.SaveActivationCode(ctx, params)
	if err != nil {
		return model.ActivationCode{
			CadastroID: int(response.CadastroID),
			Code:       response.Code,
		}, err
	}
	return req, nil

}
