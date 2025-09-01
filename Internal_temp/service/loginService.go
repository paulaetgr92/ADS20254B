package service

import (
	"awesomeProject/Internal_temp/model"
	Repository "awesomeProject/Internal_temp/repository"
	"awesomeProject/middleware"

	db "awesomeProject/db/sqlc"
	"context"

	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginService struct {
	Repo Repository.CreateLoginRepositoryInterface
}

func NewLoginoService(r Repository.CreateLoginRepositoryInterface) *LoginService {
	return &LoginService{Repo: r}
}

func (s *LoginService) CreateLoginUser(ctx context.Context, data model.LoginRequest) (db.Login, error) {
	arg := db.CreateLoginParams{
		Email:    data.Email,
		Password: data.Password,
	}
	return s.Repo.CreateLogin(ctx, arg)
}

func (s *LoginService) LoginUser(ctx context.Context, data model.LoginRequest) (string, error) {

	user, err := s.Repo.GetLogin(ctx, data.Email)
	if err != nil {
		return "", errors.New("usuário ou senha inválidos")
	}

	if user.Password != data.Password {
		return "", errors.New("usuário ou senha inválidos")
	}

	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
