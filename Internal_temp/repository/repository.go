package Repository

import (
	db "awesomeProject/db/sqlc"
	"context"
)

type CadastroRepositoryInterface interface {
	CreateCadastroRepository(ctx context.Context, arg db.CreateCadastroParams) (db.Cadastro, error)
	UpdateActivationCode(ctx context.Context, arg db.UpdateActivationCodeParams) error
}

type TokenHistRepositoryInterface interface {
	CreateTokenHist(ctx context.Context, arg db.CreateTokenHistParams) error
	GetUserTokensHist(ctx context.Context, arg db.GetUserTokensHistParams) (db.GetUserTokensHistRow, error)
}

type CreateLoginRepositoryInterface interface {
	CreateLogin(ctx context.Context, arg db.CreateLoginParams) (db.Login, error)
	GetLogin(ctx context.Context, arg string) (db.Login, error)
}

type SellerRepositoryInterface interface {
	UpdateSellerStatus(ctx context.Context, data db.UpdateCadastroStatusParams) error
	GetSellerByCNPJ(ctx context.Context, code string) (db.GetSellerByCNPJRow, error)
}
