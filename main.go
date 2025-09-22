package main

import (
	"awesomeProject/Internal_temp/handler"
	Repository "awesomeProject/Internal_temp/repository"
	"awesomeProject/Internal_temp/service"
	"awesomeProject/config"
	"awesomeProject/db/dataSrc"
	dbsqlc "awesomeProject/db/sqlc"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar .env: ", err)
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

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

	// Cria Activation Repository
	activationRepo := Repository.NewActivationNewRepository(baseRepo)

	twilioService := service.NewTwilioService()

	// Serviços
	cadastroSvc := service.NewCadastroService(cadastroRepo, sellerRepo, twilioService, activationRepo)
	tokenHistSvc := service.NewUserTokensHistService(tokenHistRepo)
	loginSvc := service.NewLoginService(loginRepo) // nome corrigido
	produtoSvc := service.NewProdutoService(produtoRepo)

	// Handlers
	cadastroHandler := handler.NewCadastroHandler(cadastroSvc)
	userTokensHistHandler := handler.NewUserTokensHistHandler(tokenHistSvc)
	loginHandler := handler.NewLoginHandler(loginSvc)
	sellerHandler := handler.NewSellerHandler(cadastroSvc)
	produtoHandler := handler.NewProdutoHandler(produtoSvc)

	// Rotas
	config.SetupRoutes(
		e,
		cadastroHandler,
		userTokensHistHandler,
		loginHandler,
		sellerHandler,
		produtoHandler,
	)

	log.Println("Servidor rodando na porta 8080")
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
