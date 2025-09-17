package repository

import (
	db "awesomeProject/db/sqlc"
	"context"
)

type LoginNewRepository struct {
	*BaseRepository
}

func NewLoginRepository(base *BaseRepository) *LoginNewRepository {
	return &LoginNewRepository{BaseRepository: base}
}

func (r *LoginNewRepository) CreateLogin(ctx context.Context, arg db.CreateLoginParams) (db.Login, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return db.Login{}, err
	}
	return r.Queries.CreateLogin(ctx, arg)
}

func (r *LoginNewRepository) GetLogin(ctx context.Context, arg string) (db.Login, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return db.Login{}, err
	}
	return r.Queries.GetLogin(ctx, arg)
}
