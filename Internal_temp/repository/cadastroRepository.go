package Repository

import (
	db "awesomeProject/db/sqlc"
	"context"
)

type CadastroNewRepository struct {
	*BaseRepository
}

func NewCadastroNewRepository(base *BaseRepository) *CadastroNewRepository {
	return &CadastroNewRepository{
		BaseRepository: base,
	}
}

func (r *CadastroNewRepository) CreateCadastroRepository(ctx context.Context, arg db.CreateCadastroParams) error {
	if err := r.GetConnection(ctx); err != nil {
		return err
	}

	return r.Queries.CreateCadastro(ctx, arg)
}
