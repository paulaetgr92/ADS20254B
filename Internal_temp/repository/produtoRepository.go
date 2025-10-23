package Repository

import (
	db "awesomeProject/db/sqlc"
	"context"
)

type ProdutosNewRepository struct {
	*BaseRepository
}

func NewProdutosRepository(base *BaseRepository) *ProdutosNewRepository {
	return &ProdutosNewRepository{BaseRepository: base}
}

func (r *ProdutosNewRepository) CreateProdutoRepository(ctx context.Context, arg db.CreateProductParams) (int64, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return 0, err
	}
	return r.Queries.CreateProduct(ctx, arg)
}

func (r *ProdutosNewRepository) GetProdutoByIdRepository(ctx context.Context, arg int64) (db.GetProdutoByIdRow, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return db.GetProdutoByIdRow{}, err
	}
	return r.Queries.GetProdutoById(ctx, arg)
}

func (r *ProdutosNewRepository) GetProdutoByDisponibilidade(ctx context.Context, arg db.GetProdutoByDisponibilidadeParams) ([]db.GetProdutoByDisponibilidadeRow, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return []db.GetProdutoByDisponibilidadeRow{}, err
	}
	return r.Queries.GetProdutoByDisponibilidade(ctx, arg)
}

func (r *ProdutosNewRepository) AtualizarProduto(ctx context.Context, arg db.AtualizarProdutoByIDParams) (db.Produto, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return db.Produto{}, err
	}
	return r.Queries.AtualizarProdutoByID(ctx, arg)
}

func (r *ProdutosNewRepository) DeleteProdutoByIdRepository(ctx context.Context, Id int64) error {
	err := r.GetConnection(ctx)
	if err != nil {
		return err
	}
	return r.Queries.InativarProdutoByID(ctx, Id)
}

func (r *ProdutosNewRepository) GetProdutoRepository(ctx context.Context) ([]db.Produto, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	return r.Queries.ListProducts(ctx)
}
