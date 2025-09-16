package service

import (
	"awesomeProject/Internal_temp/model"
	Repository "awesomeProject/Internal_temp/repository"
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

func NewLoginService(r Repository.CreateLoginRepositoryInterface) *LoginService {
	return &LoginService{Repo: r}
}

// Apenas autenticação + JWT
func (s *LoginService) LoginUser(ctx context.Context, data model.LoginRequest) (string, error) {
	user, err := s.Repo.GetLogin(ctx, data.Email)
	if err != nil {
		return "", errors.New("usuário ou senha inválidos")
	}

	// 🔐 Comparar senha digitada com hash salvo no banco
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return "", errors.New("usuário ou senha inválidos")
	}

	// Claims do JWT
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	// Gera o token assinado
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
