package service

import (
	"awesomeProject/Internal_temp/model"
	Repository "awesomeProject/Internal_temp/repository"
	db "awesomeProject/db/sqlc"
	"context"
)

type AdminServiceInterface interface {
	CreateAdminService(ctx context.Context, data model.AdminRequest) error
	LoginAdminService(ctx context.Context, data model.LoginRequest) error
	ListAllProdutosService(ctx context.Context) ([]model.ProdutosResponse, error)
}

type AdminService struct {
	Repo Repository.AdminNewRepository
	ProdutoService ProdutoServiceInterface
}

var _ AdminServiceInterface = (*AdminService)(nil)

func NewAdminService(r Repository.AdminNewRepository, ps ProdutoServiceInterface) *AdminService {
	return &AdminService{Repo: r, ProdutoService: ps}
}

func (r *AdminService) CreateAdminService(ctx context.Context, data model.AdminRequest) error {
	arg := db.CreateAdminParams{
		Cnpj:     data.Cnpj,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	_, err := r.Repo.CreateAdmin(ctx, arg)
	return err
}

func (r *AdminService) LoginAdminService(ctx context.Context, data model.LoginRequest) error {
	arg := db.CreateLoginAdminParams{
		Email:    data.Email,
		Password: data.Password,
	}
	_, err := r.Repo.LoginNewAdmin(ctx, arg)
	return err
}

func (r *AdminService) ListAllProdutosService(ctx context.Context) ([]model.ProdutosResponse, error) {
	return r.ProdutoService.ListAllProdutosService(ctx)
}

