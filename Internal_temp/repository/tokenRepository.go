package Repository

import (
	db "awesomeProject/db/sqlc"
	"context"
)

type UserTokensHistRepository struct {
	BaseRepository
}

func NewUserTokensHistRepository(baseRepo BaseRepository) *UserTokensHistRepository {
	return &UserTokensHistRepository{
		BaseRepository: baseRepo,
	}
}

func (b *UserTokensHistRepository) GetUserTokensHist(ctx context.Context, arg db.GetUserTokensHistParams) (db.GetUserTokensHistRow, error) {
	err := b.GetConnection(ctx)
	if err != nil {
		return db.GetUserTokensHistRow{}, err
	}
	return b.Queries.GetUserTokensHist(ctx, arg)
}

func (b *UserTokensHistRepository) CreateTokenHist(ctx context.Context, arg db.CreateTokenHistParams) error {
	err := b.GetConnection(ctx)
	if err != nil {
		return err
	}
	return b.Queries.CreateTokenHist(ctx, arg)
}
