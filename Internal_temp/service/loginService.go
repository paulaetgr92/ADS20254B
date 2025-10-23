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

func NewLoginService(r Repository.CreateLoginRepositoryInterface) *LoginService {
	return &LoginService{Repo: r}
}

func (s *LoginService) CreateLoginUser(ctx context.Context, data model.LoginRequest) (db.CreateLoginRow, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.CreateLoginRow{}, errors.New("erro ao gerar hash da senha")
	}

	arg := db.CreateLoginParams{
		Email:    data.Email,
		Password: string(hashedPassword),
		Status:   "pendente",
	}

	return s.Repo.CreateLogin(ctx, arg)
}

func (s *LoginService) LoginUser(ctx context.Context, data model.LoginRequest) (string, error) {

	user, err := s.Repo.GetLogin(ctx, data.Email)
	if err != nil {
		return "", errors.New("usuário ou senha inválidos")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return "", errors.New("usuário ou senha inválidos")
	}



	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JWTKey)
	if err != nil {
		return "", errors.New("erro ao gerar token")
	}

	return tokenString, nil
}
