package service

import (
	"awesomeProject/Internal_temp/model"
	"awesomeProject/Internal_temp/repository"
	db "awesomeProject/db/sqlc"
	"context"
	"strconv"
)

type SaleService struct {
	repo Repository.SalesNewRepository
}

func NewSaleService(repo Repository.SalesNewRepository) *SaleService {
	return &SaleService{
		repo: repo,
	}
}

func (s *SaleService) CreateSale(ctx context.Context, request model.SaleRequest) (db.Sale, error) {
	arg := db.CreateSaleParams{
		ProdutoID:  request.ProdutoID,
		Quantidade: request.Quantidade,
		TempoValor: strconv.FormatInt(request.TempoValor, 10),
	}

	sale, err := s.repo.CreateSale(ctx, arg)
	if err != nil {
		return db.Sale{}, err
	}

	return sale, nil
}
