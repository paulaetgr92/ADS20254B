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

	// Conexão com o banco
	conn, err := dataSrc.Connect()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco: ", err)
	}

	// Queries e base repository
	queries := dbsqlc.New(conn)
	baseRepo := repository.NewBaseRepository(queries, conn)

	// Serviços auxiliares
	twilioService := service.NewTwilioService()

	// Repositórios
	cadastroRepo := repository.NewCadastroRepository(baseRepo)
	sellerRepo := repository.NewSellerRepository(baseRepo)
	tokenHistRepo := repository.NewUserTokensHistRepository(baseRepo)
	loginRepo := repository.NewLoginRepository(baseRepo)
	produtoRepo := repository.NewProdutosRepository(baseRepo)

	// Serviços
	cadastroSvc := service.NewCadastroService(cadastroRepo, sellerRepo, twilioService)
	tokenHistSvc := service.NewUserTokensHistService(tokenHistRepo)
	loginSvc := service.NewLoginService(loginRepo)
	produtoSvc := service.NewProdutoService(produtoRepo)

	// Handlers
	cadastroHandler := handler.NewCadastroHandler(cadastroSvc)
	userTokensHistHandler := handler.NewUserTokensHistHandler(tokenHistSvc)
	loginHandler := handler.NewLoginHandler(loginSvc)
	getLoginHandler := handler.NewLoginHandler(loginSvc) // Pode ser o mesmo service
	sellerHandler := handler.NewSellerHandler(sellerRepo)
	produtoHandler := handler.NewProdutoHandler(produtoSvc)

	// Configuração das rotas
	config.SetupRoutes(e,
		cadastroHandler,
		userTokensHistHandler,
		loginHandler,
		getLoginHandler,
		sellerHandler,
		produtoHandler,
	)

	log.Println("Servidor rodando na porta 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
