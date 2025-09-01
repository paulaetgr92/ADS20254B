package config

import (
	handler2 "awesomeProject/Internal_temp/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, cadastroHandler *handler2.CadastroHandler, handler *handler2.UserTokensHistHandler, loginHandler *handler2.LoginHandler, getLoginHandler *handler2.LoginHandler) {
	api := e.Group("/api/v1")

	cadastros := api.Group("/cadastros")
	{
		cadastros.POST("", cadastroHandler.CreateCadastro)
	}

	login := api.Group("/login")
	{
		login.POST("", loginHandler.Login)
	}
}
