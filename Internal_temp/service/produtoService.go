package service

import (
	"awesomeProject/Internal_temp/model"
	Repository "awesomeProject/Internal_temp/repository"
	db "awesomeProject/db/sqlc"
	"context"
	"database/sql"
	"fmt"
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

func (s *ProdutoService) GetProdutoByIdService(ctx context.Context, id int64) (model.GetProdutosByIdResponse, error) {
	produto, err := s.repo.GetProdutoByIdRepository(ctx, id)
	if err != nil {
		return model.GetProdutosByIdResponse{}, err
	}

	response := model.GetProdutosByIdResponse{
		Id:          produto.IDRoupa,
		Categoria:   produto.Categoria,
		Tamanho:     produto.Tamanho,
		Cores:       produto.Cores.String,
		TempoValor:  int64(produto.TempoValor.Float64),
		Localizacao: produto.Localizacao.String,
		Ativo:       produto.Status.Valid,
	}

	return response, nil
}

func (s *ProdutoService) ListProdutoService(ctx context.Context) ([]model.ProdutosResponse, error) {
	produtos, err := s.repo.GetProdutoRepository(ctx)
	if err != nil {
		fmt.Println("Erro ao chamar repository:", err)
		return nil, err
	}

	var result []model.ProdutosResponse
	for _, p := range produtos {

		result = append(result, model.ProdutosResponse{
			Categoria:   p.Categoria,
			Tamanho:     p.Tamanho,
			Cores:       p.Cores.String,
			TempoValor:  int64(p.TempoValor.Float64),
			Localizacao: p.Localizacao.String,
		})
	}

	fmt.Println("Produtos mapeados:", result)
	return result, nil
}

func (s *ProdutoService) UpdateProdutoByIdService(ctx context.Context, id int64, data model.ProdutosRequest) (model.ProdutosResponse, error) {
	arg := db.AtualizarProdutoByIDParams{
		IDRoupa:   id,
		Categoria: data.Categoria,
		Tamanho:   data.Tamanho,
		Cores: sql.NullString{
			String: data.Cores,
			Valid:  true,
		},
		TempoValor: sql.NullFloat64{
			Float64: float64(data.TempoValor),
			Valid:   true,
		},
		Status: sql.NullString{
			String: data.Status,
			Valid:  true,
		},
		Localizacao: sql.NullString{
			String: data.Localizacao,
			Valid:  true,
		},
	}

	produto, err := s.repo.AtualizarProduto(ctx, arg)
	if err != nil {
		return model.ProdutosResponse{}, err
	}

	return model.ProdutosResponse{
		Categoria:   produto.Categoria,
		Tamanho:     produto.Tamanho,
		Cores:       produto.Cores.String,
		TempoValor:  int64(produto.TempoValor.Float64),
		Localizacao: produto.Localizacao.String,
	}, nil
}

func (s *ProdutoService) DeleteProdutoByIdService(ctx context.Context, id int64) error {

	return s.repo.DeleteProdutoByIdRepository(ctx, id)
}
