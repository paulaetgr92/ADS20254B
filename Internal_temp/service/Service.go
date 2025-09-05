package service

import (
	"awesomeProject/Internal_temp/model"
	db "awesomeProject/db/sqlc"
	"context"
)

type CadastroServiceInterface interface {
	CreateCadastro(ctx context.Context, data model.CadastroRequest) (db.Cadastro, error)
	VerifySeller(ctx context.Context, data model.TwillioModelRequest) error
}

type TokenServiceInterface interface {
	GetUserTokensHist(ctx context.Context, payload model.PayloadDTO) error
}

type LoginServiceinterface interface {
	LoginUser(ctx context.Context, data model.LoginRequest) (string, error)
	CreateLoginUser(ctx context.Context, data model.LoginRequest) (db.Login, error)
}
