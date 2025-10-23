package config

import (
	handler2 "awesomeProject/Internal_temp/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(
	e *echo.Echo,
	cadastroHandler *handler2.CadastroHandler,
	handler *handler2.UserTokensHistHandler,
	loginHandler *handler2.LoginHandler,
	sellerHandler *handler2.SellerHandler,
	produtoHandler *handler2.ProdutoHandler,
	salesHandler *handler2.SaleHandler,
	adminHandler *handler2.AdminHandler,
) {
	api := e.Group("/api/v1")

	cadastros := api.Group("/cadastros")
	{
		cadastros.POST("", cadastroHandler.CreateCadastro)
	}

	login := api.Group("/login")
	{
		login.POST("", loginHandler.Login)
	}

	produtos := api.Group("/produtos")
	{
		produtos.POST("", produtoHandler.CreateProductHandler)
		produtos.GET("", produtoHandler.ListProdutosHandler)
		produtos.GET("/:id", produtoHandler.GetProductByIdHandler)
		produtos.PUT("/:id", produtoHandler.UpdateProdutoByIdHandler)
		produtos.PUT("/:id/inativar", produtoHandler.InativarProdutoHandler)
	}

	admin := api.Group("/admin")
	{
		admin.POST("/create", adminHandler.CreateAdminHandler)
		admin.POST("/login", adminHandler.CreateLoginAdminHandler)
	}

	vendas := api.Group("/sales")
	{
		vendas.POST("", salesHandler.CreateSale)

	}

	api.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
}
