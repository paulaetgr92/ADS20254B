package Repository

import (
	"awesomeProject/Internal_temp/model"
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

type ProdutoRepositoryInterface interface {
	DeleteProdutoByIdRepository(ctx context.Context, Id int64) error
	AtualizarProduto(ctx context.Context, arg db.AtualizarProdutoByIDParams) (db.Produto, error)
	GetProdutoByIdRepository(ctx context.Context, arg int64) (db.GetProdutoByIdRow, error)
	CreateProdutoRepository(ctx context.Context, arg db.CreateProductParams) (int64, error)
	GetProdutoRepository(ctx context.Context) ([]db.Produto, error)
}

type ActivationCodeRepositoryInterface interface {
	SaveActivationCode(ctx context.Context, data db.SaveActivationCodeParams) error
	GetActivationCode(ctx context.Context, arg db.GetCadastroByActivationCodeParams) (db.GetCadastroByActivationCodeRow, error)
}

type CreateSaleInterface interface {
	CreateSale(ctx context.Context, arg db.CreateSaleParams) (db.Sale, error)
}

type SellersRepositoryInterface interface {
	CreateSeller(ctx context.Context, seller model.Seller) (db.CreateSellerRow, error)
	UpdateSellerStatus(ctx context.Context, params db.UpdateCadastroStatusParams) error
}
