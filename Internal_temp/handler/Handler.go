package handler

import "github.com/labstack/echo/v4"

type CadastroHandlerInterface interface {
	CreateCadastro(c echo.Context) error
}

type TokenHandlerInterface interface {
	GetUserTokensHist(next echo.HandlerFunc) echo.HandlerFunc
	CreateUserToken(c echo.Context) error
}

type CreateLoginHandlerInterface interface {
	Login(c echo.Context) error
}

type twillioHandlerInterface interface {
	VerifyCode(c echo.Context) error
}
type ProdutoHandlerInterface interface {
	CreateProductHandler(c echo.Context) error
	InativarProdutoHandler(c echo.Context) error
	UpdateProdutoByIdHandler(c echo.Context) error
	ListProdutosHandler(c echo.Context) error
	GetProductByIdHandler(c echo.Context) error
}
