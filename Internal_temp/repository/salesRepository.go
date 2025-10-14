package Repository

import (
	db "awesomeProject/db/sqlc"
	"context"
)

type SalesNewRepository struct {
	*BaseRepository
}

func NewSalesRepository(base *BaseRepository) *SalesNewRepository {
	return &SalesNewRepository{BaseRepository: base}
}

func (r *SalesNewRepository) CreateSale(ctx context.Context, arg db.CreateSaleParams) (db.Sale, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return db.Sale{}, err
	}
	return r.Queries.CreateSale(ctx, arg)
}
