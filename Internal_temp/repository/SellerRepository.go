package Repository

import (
	db "awesomeProject/db/sqlc"
	"context"
)

type SellerRepository struct {
	BaseRepository
}

func NewSellerRepository(baseRepo BaseRepository) *SellerRepository {
	return &SellerRepository{
		BaseRepository: baseRepo,
	}
}

func (r *SellerRepository) UpdateSellerStatus(ctx context.Context, data db.UpdateCadastroStatusParams) error {
	if err := r.GetConnection(ctx); err != nil {
		return err
	}
	return r.Queries.UpdateCadastroStatus(ctx, data)
}

func (r *SellerRepository) UpdateActivationCode(ctx context.Context, arg db.UpdateCadastroStatusParams) (error, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return err, nil
	}
	return r.Queries.UpdateCadastroStatus(ctx, arg), nil
}
