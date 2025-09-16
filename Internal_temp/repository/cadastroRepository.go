package repository

import (
	db "awesomeProject/db/sqlc"
	"context"
)

type CadastroRepository struct {
	*BaseRepository
}

func NewCadastroRepository(base *BaseRepository) *CadastroRepository {
	return &CadastroRepository{
		BaseRepository: base,
	}
}

// Criação de cadastro (já recebe senha criptografada do service)
func (r *CadastroRepository) CreateCadastro(ctx context.Context, arg db.CreateCadastroParams) (db.Cadastro, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return db.Cadastro{}, err
	}

	// 👉 aqui não mexemos na senha, apenas salvamos o que o service passou
	return r.Queries.CreateCadastro(ctx, arg)
}

// Atualização do código de ativação
func (r *CadastroRepository) UpdateActivationCode(ctx context.Context, arg db.UpdateActivationCodeParams) error {
	err := r.GetConnection(ctx)
	if err != nil {
		return err
	}
	return r.Queries.UpdateActivationCode(ctx, arg)
}
