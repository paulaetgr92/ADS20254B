package config

import (
	handler2 "awesomeProject/Internal_temp/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, cadastroHandler *handler2.CadastroHandler, handler *handler2.UserTokensHistHandler, loginHandler *handler2.LoginHandler, sellerHandler *handler2.SellerHandler, produtoHandler *handler2.ProdutoHandler) {
	api := e.Group("/api/v1")

	cadastros := api.Group("/cadastros")
	{
		cadastros.POST("", cadastroHandler.CreateCadastro)
	}

	login := api.Group("/login")
	{
		login.POST("", loginHandler.Login)
	}

	sellers := api.Group("/sellers")
	{
		sellers.POST("/verify", sellerHandler.VerifySeller)
	}

	produtos := api.Group("/produtos")
	{
		produtos.POST("", produtoHandler.CreateProductHandler)
	}

	api.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
}
