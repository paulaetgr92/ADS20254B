package repository

import (
	db "awesomeProject/db/sqlc"
	"context"
)

type SellerRepository struct {
	*BaseRepository
}

func NewSellerRepository(baseRepo *BaseRepository) *SellerRepository {
	return &SellerRepository{
		BaseRepository: baseRepo,
	}
}

func (b *SellerRepository) GetSellerByCNPJ(ctx context.Context, code string) (db.GetSellerByCNPJRow, error) {
	err := b.GetConnection(ctx)
	if err != nil {
		return db.GetSellerByCNPJRow{}, err
	}
	return b.Queries.GetSellerByCNPJ(ctx, code)
}

func (r *SellerRepository) UpdateSellerStatus(ctx context.Context, data db.UpdateCadastroStatusParams) error {
	if err := r.GetConnection(ctx); err != nil {
		return err
	}
	return r.Queries.UpdateCadastroStatus(ctx, data)
}
