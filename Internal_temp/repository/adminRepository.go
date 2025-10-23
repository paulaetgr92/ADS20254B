package Repository

import (
	db "awesomeProject/db/sqlc"
	"context"
)

type AdminNewRepository struct {
	*BaseRepository
}

func NewAdminRepository(base *BaseRepository) *AdminNewRepository {
	return &AdminNewRepository{BaseRepository: base}
}

func (r *AdminNewRepository) CreateAdmin(ctx context.Context, arg db.CreateAdminParams) (db.Admin, error) {
	if err := r.GetConnection(ctx); err != nil {
		return db.Admin{}, err
	}
	return r.Queries.CreateAdmin(ctx, arg)
}

func (r *AdminNewRepository) LoginNewAdmin(ctx context.Context, arg db.CreateLoginAdminParams) (db.CreateLoginAdminRow, error) {
	if err := r.GetConnection(ctx); err != nil {
		return db.CreateLoginAdminRow{}, err
	}
	return r.Queries.CreateLoginAdmin(ctx, arg)
}
