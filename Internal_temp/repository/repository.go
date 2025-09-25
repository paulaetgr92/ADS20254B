package Repository

import (
	db "awesomeProject/db/sqlc"
	"context"
)

type CadastroRepositoryInterface interface {
	CreateCadastroRepository(ctx context.Context, arg db.CreateCadastroParams) (db.Cadastro, error)
}

type TokenHistRepositoryInterface interface {
	CreateTokenHist(ctx context.Context, arg db.CreateTokenHistParams) error
	GetUserTokensHist(ctx context.Context, arg db.GetUserTokensHistParams) (db.GetUserTokensHistRow, error)
}

type CreateLoginRepositoryInterface interface {
	CreateLogin(ctx context.Context, arg db.CreateLoginParams) (db.CreateLoginRow, error)
	GetLogin(ctx context.Context, arg string) (db.GetLoginRow, error)
}

type SellerRepositoryInterface interface {
	UpdateSellerStatus(ctx context.Context, data db.UpdateCadastroStatusParams) error
	UpdateActivationCode(ctx context.Context, arg db.UpdateCadastroStatusParams) (error, error)
}

type ProdutoRepositoryInterface interface {
	DeleteProdutoByIdRepository(ctx context.Context, arg db.DeleteProdutoByIDParams) (db.Produto, error)
	AtualizarProduto(ctx context.Context, arg db.AtualizarProdutoByIDParams) (db.Produto, error)
	GetProdutoByDisponibilidade(ctx context.Context, arg db.GetProdutoByDisponibilidadeParams) ([]db.GetProdutoByDisponibilidadeRow, error)
	GetProdutoByIdRepository(ctx context.Context, arg int64) (db.GetProdutoByIdRow, error)
	CreateProdutoRepository(ctx context.Context, arg db.CreateProductParams) (int64, error)
}

type ActivationCodeRepositoryInterface interface {
	SaveActivationCode(ctx context.Context, data db.SaveActivationCodeParams) error
	GetActivationCode(ctx context.Context, arg db.GetCadastroByActivationCodeParams) (db.GetCadastroByActivationCodeRow, error)
}
