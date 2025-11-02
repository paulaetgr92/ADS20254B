package Repository

import (
	db "awesomeProject/db/sqlc"
	"context"
)

type ActivationNewRepository struct {
	*BaseRepository
}

func NewActivationNewRepository(base *BaseRepository) *ActivationNewRepository {
	return &ActivationNewRepository{
		BaseRepository: base,
	}
}

func (r *ActivationNewRepository) SaveActivationCode(ctx context.Context, arg db.SaveActivationCodeParams) error {
	if err := r.GetConnection(ctx); err != nil {
		return err
	}
	return r.Queries.SaveActivationCode(ctx, arg)
}

func (r *ActivationNewRepository) GetActivationCode(ctx context.Context, arg db.GetCadastroByActivationCodeParams) (db.GetCadastroByActivationCodeRow, error) {
	if err := r.GetConnection(ctx); err != nil {
		return db.GetCadastroByActivationCodeRow{}, err
	}
	return r.Queries.GetCadastroByActivationCode(ctx, arg)
}
