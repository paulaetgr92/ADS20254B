package service

import (
	"awesomeProject/Internal_temp/model"
	"awesomeProject/Internal_temp/repository"
	db "awesomeProject/db/sqlc"
	db2 "awesomeProject/db/sqlc"
	"context"
	"database/sql"
)

type CadastroService struct {
	Repo Repository.CadastroRepositoryInterface
}

func NewCadastroService(r Repository.CadastroRepositoryInterface) *CadastroService {
	return &CadastroService{Repo: r}
}

func (a *CadastroService) CreateCadastro(ctx context.Context, data model.CadastroRequest) error {

	arg := db.CreateCadastroParams{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Acessid: sql.NullInt64{
			Int64: data.PayloadDTO.AccessID,
			Valid: true,
		},
	}

	return a.Repo.CreateCadastroRepository(ctx, db2.CreateCadastroParams(arg))
}
