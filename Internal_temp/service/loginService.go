package service

import (
	"awesomeProject/Internal_temp/model"
	Repository "awesomeProject/Internal_temp/repository"
	db "awesomeProject/db/sqlc"
	"awesomeProject/middleware"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	Repo Repository.CreateLoginRepositoryInterface
}

// Construtor corrigido
func NewLoginService(r Repository.CreateLoginRepositoryInterface) *LoginService {
	return &LoginService{Repo: r}
}

// Cria usuário com senha hash
func (s *LoginService) CreateLoginUser(ctx context.Context, data model.LoginRequest) (db.Cadastro, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.Cadastro{}, errors.New("erro ao gerar hash da senha")
	}

	arg := db.CreateLoginParams{
		Email:    data.Email,
		Password: string(hashedPassword),
	}

	return s.Repo.CreateLogin(ctx, arg)
}

// Login: valida senha e retorna JWT
func (s *LoginService) LoginUser(ctx context.Context, data model.LoginRequest) (string, error) {
	// Busca usuário pelo email
	user, err := s.Repo.GetLogin(ctx, data.Email)
	if err != nil {
		return "", errors.New("usuário ou senha inválidos")
	}

	// Compara senha com hash armazenado
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return "", errors.New("usuário ou senha inválidos")
	}

	// Cria claims do JWT
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	// Gera token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JWTKey)
	if err != nil {
		return "", errors.New("erro ao gerar token")
	}

	return tokenString, nil
}
