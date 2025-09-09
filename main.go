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

	twilioService := service.NewTwilioService()

	queries := dbsqlc.New(conn)

	baseRepo := Repository.NewBaseRepository(queries, conn)

	cadastroRepo := Repository.NewCadastroNewRepository(baseRepo)
	tokenHistRepo := Repository.NewUserTokensHistRepository(*baseRepo)
	LoginRepo := Repository.NewLoginRepository(baseRepo)
	Login := Repository.NewLoginRepository(baseRepo)
	SellerRepo := Repository.NewSellerRepository(*baseRepo)
	ProdutoRepo := Repository.NewProdutosRepository(baseRepo)

	cadastroSvc := service.NewCadastroService(cadastroRepo, SellerRepo, twilioService)
	tokenHistSvc := service.NewUserTokensHistService(tokenHistRepo)
	LoginSvc := service.NewLoginoService(LoginRepo)
	GetLoginSVC := service.NewLoginoService(Login)
	CreateProdutoSVC := service.NewProdutoService(ProdutoRepo)

	cadastroHandler := handler.NewCadastroHandler(cadastroSvc)
	userTokensHistHandler := handler.NewUserTokensHistHandler(tokenHistSvc)
	loginHandler := handler.NewLoginHandler(LoginSvc)
	GetLoginHandler := handler.NewLoginHandler(GetLoginSVC)
	SellerHandler := handler.NewSellerHandler(cadastroSvc)
	ProdutoHandler := handler.NewProdutoHandler(CreateProdutoSVC)

	config.SetupRoutes(e, cadastroHandler, userTokensHistHandler, loginHandler, GetLoginHandler, SellerHandler, ProdutoHandler)

	log.Println("Servidor rodando na porta 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
