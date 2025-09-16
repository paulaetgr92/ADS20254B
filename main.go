package main

import (
	"awesomeProject/Internal_temp/handler"
	"awesomeProject/Internal_temp/repository"
	"awesomeProject/Internal_temp/service"
	"awesomeProject/config"
	"awesomeProject/db/dataSrc"
	dbsqlc "awesomeProject/db/sqlc"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar .env: ", err)
	}

	e := echo.New()

	conn, err := dataSrc.Connect()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco: ", err)
	}
	queries := dbsqlc.New(conn)

	baseRepo := Repository.NewBaseRepository(queries, conn)
	cadastroRepo := Repository.NewCadastroNewRepository(baseRepo)
	tokenHistRepo := Repository.NewUserTokensHistRepository(*baseRepo)
	loginRepo := Repository.NewLoginRepository(baseRepo)
	sellerRepo := Repository.NewSellerRepository(*baseRepo)
	produtoRepo := Repository.NewProdutosRepository(baseRepo)

	twilioService := service.NewTwilioService()

	cadastroSvc := service.NewCadastroService(cadastroRepo, sellerRepo, twilioService)
	tokenHistSvc := service.NewUserTokensHistService(tokenHistRepo)
	loginSvc := service.NewLoginoService(loginRepo)
	produtoSvc := service.NewProdutoService(produtoRepo)

	cadastroHandler := handler.NewCadastroHandler(cadastroSvc)
	userTokensHistHandler := handler.NewUserTokensHistHandler(tokenHistSvc)
	loginHandler := handler.NewLoginHandler(loginSvc)
	sellerHandler := handler.NewSellerHandler(cadastroSvc)
	produtoHandler := handler.NewProdutoHandler(produtoSvc)

	config.SetupRoutes(
		e,
		cadastroHandler,
		userTokensHistHandler,
		loginHandler,
		loginHandler,
		sellerHandler,
		produtoHandler,
	)

	log.Println("Servidor rodando na porta 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
