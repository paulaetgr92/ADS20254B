package service

import (
	"awesomeProject/Internal_temp/model"
	Repository "awesomeProject/Internal_temp/repository"
	db "awesomeProject/db/sqlc"
	"context"
	"database/sql"
)

type ProdutoService struct {
	repo Repository.ProdutoRepositoryInterface
}

func NewProdutoService(produtoRepo Repository.ProdutoRepositoryInterface) *ProdutoService {
	return &ProdutoService{
		repo: produtoRepo,
	}
}

func (s *ProdutoService) CreateProduct(ctx context.Context, data model.ProdutosRequest) (int64, error) {
	arg := db.CreateProductParams{
		Categoria: data.Categoria,
		Tamanho:   data.Tamanho,
		Cores: sql.NullString{
			String: data.Cores,
			Valid:  data.Cores != "",
		},
		TempoValor: sql.NullFloat64{
			Float64: float64(data.TempoValor),
			Valid:   data.TempoValor > 0,
		},
		Status: sql.NullString{
			String: data.Status,
			Valid:  data.Status != "",
		},
		Localizacao: sql.NullString{
			String: data.Localizacao,
			Valid:  data.Localizacao != "",
		},
	}

	id, err := s.repo.CreateProdutoRepository(ctx, arg)
	if err != nil {
		return 0, err
	}

	return id, nil
}
