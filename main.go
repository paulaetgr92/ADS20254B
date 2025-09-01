package main

import (
	"log"

	"awesomeProject/Internal_temp/handler"
	"awesomeProject/Internal_temp/repository"
	"awesomeProject/Internal_temp/service"
	"awesomeProject/config"
	"awesomeProject/db/dataSrc"
	dbsqlc "awesomeProject/db/sqlc"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	conn, err := dataSrc.Connect()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco: ", err)
	}

	queries := dbsqlc.New(conn)

	baseRepo := Repository.NewBaseRepository(queries, conn)

	cadastroRepo := Repository.NewCadastroNewRepository(baseRepo)
	tokenHistRepo := Repository.NewUserTokensHistRepository(*baseRepo)
	LoginRepo := Repository.NewLoginRepository(baseRepo)
	Login := Repository.NewLoginRepository(baseRepo)

	cadastroSvc := service.NewCadastroService(cadastroRepo)
	tokenHistSvc := service.NewUserTokensHistService(tokenHistRepo)
	LoginSvc := service.NewLoginoService(LoginRepo)
	GetLoginSVC := service.NewLoginoService(Login)

	cadastroHandler := handler.NewCadastroHandler(cadastroSvc)
	userTokensHistHandler := handler.NewUserTokensHistHandler(tokenHistSvc)
	loginHandler := handler.NewLoginHandler(LoginSvc)
	GetLoginHandler := handler.NewLoginHandler(GetLoginSVC)

	config.SetupRoutes(e, cadastroHandler, userTokensHistHandler, loginHandler, GetLoginHandler)

	log.Println("Servidor rodando na porta 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
