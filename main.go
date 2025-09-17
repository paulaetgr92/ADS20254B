package main

import (
	"database/sql"
	"log"
	"time"

	"awesomeProject/Internal_temp/handler"
	Repository "awesomeProject/Internal_temp/repository"
	"awesomeProject/Internal_temp/service"
	"awesomeProject/config"
	"awesomeProject/db/dataSrc"
	dbsqlc "awesomeProject/db/sqlc"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: nenhum arquivo .env encontrado, usando variáveis do ambiente")
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Retry na conexão com o Postgres
	var conn *sql.DB
	var err error
	for i := 0; i < 10; i++ {
		conn, err = dataSrc.Connect()
		if err == nil {
			break
		}
		log.Println("Banco ainda não pronto, tentando novamente em 3s...")
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal("Erro ao conectar ao banco após várias tentativas:", err)
	}

	queries := dbsqlc.New(conn)

	// Repositórios
	baseRepo := Repository.NewBaseRepository(queries, conn)
	cadastroRepo := Repository.NewCadastroNewRepository(baseRepo)
	tokenHistRepo := Repository.NewUserTokensHistRepository(*baseRepo)
	loginRepo := Repository.NewLoginRepository(baseRepo)
	sellerRepo := Repository.NewSellerRepository(*baseRepo)
	produtoRepo := Repository.NewProdutosRepository(baseRepo)

	// Serviços
	twilioService := service.NewTwilioService()
	cadastroSvc := service.NewCadastroService(cadastroRepo, sellerRepo, twilioService)
	tokenHistSvc := service.NewUserTokensHistService(tokenHistRepo)
	loginSvc := service.NewLoginoService(loginRepo)
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
