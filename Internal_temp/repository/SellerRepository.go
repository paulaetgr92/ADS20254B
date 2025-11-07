package Repository

import (
	db "awesomeProject/db/sqlc"
	"context"

	"fmt"
)

type SellerRepository struct {
	BaseRepository
}

func NewSellerRepository(repository BaseRepository) *SellerRepository {
	return &SellerRepository{
		BaseRepository: repository,
	}
}

func (r *SellerRepository) CreateSeller(ctx context.Context, arg db.CreateSellerParams) (db.CreateSellerRow, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return db.CreateSellerRow{}, err
	}
	return r.Queries.CreateSeller(ctx, arg)
}

func (r *SellerRepository) UpdateSellerStatus(ctx context.Context, params db.UpdateCadastroStatusParams) error {
	err := r.Queries.UpdateCadastroStatus(ctx, params)
	if err != nil {
		return fmt.Errorf("erro ao atualizar status do seller: %v", err)
	}
	return nil
}
