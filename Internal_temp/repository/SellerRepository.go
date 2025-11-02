package Repository

import (
	"awesomeProject/Internal_temp/model"
	db "awesomeProject/db/sqlc"
	"context"
	"database/sql"
	"fmt"
)

type SellerRepository struct {
	Queries *db.Queries
}

func NewSellerRepository(queries *db.Queries) *SellerRepository {
	return &SellerRepository{
		Queries: queries,
	}
}

func (r *SellerRepository) CreateSeller(ctx context.Context, seller model.Seller) (db.CreateSellerRow, error) {
	return r.Queries.CreateSeller(ctx, db.CreateSellerParams{
		ActivationCode: sql.NullString{
			String: seller.ActivationCodes,
		},
		Name:     seller.Name,
		Email:    seller.Email,
		Password: seller.Password,
		Cpf: sql.NullString{
			String: seller.Cpf,
			Valid:  true,
		},
		Cnpj: sql.NullString{
			String: seller.Cnpj,
			Valid:  true,
		},
		Celular: seller.Phone,
		Status:  seller.Status,
	})
}

func (r *SellerRepository) UpdateSellerStatus(ctx context.Context, params db.UpdateCadastroStatusParams) error {
	err := r.Queries.UpdateCadastroStatus(ctx, params)
	if err != nil {
		return fmt.Errorf("erro ao atualizar status do seller: %v", err)
	}
	return nil
}
