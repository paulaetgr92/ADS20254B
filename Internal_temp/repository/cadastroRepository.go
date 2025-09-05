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

func (r *CadastroNewRepository) CreateCadastroRepository(ctx context.Context, arg db.CreateCadastroParams) (db.Cadastro, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return db.Cadastro{}, err
	}
	return r.Queries.CreateCadastro(ctx, arg)
}

func (r *CadastroNewRepository) UpdateActivationCode(ctx context.Context, arg db.UpdateActivationCodeParams) error {
	err := r.GetConnection(ctx)
	if err != nil {
		return err
	}
	return r.Queries.UpdateActivationCode(ctx, arg)
}
