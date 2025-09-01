package service

import (
	"awesomeProject/Internal_temp/model"
	Repository "awesomeProject/Internal_temp/repository"
	db "awesomeProject/db/sqlc"
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
)

type UserTokensHistService struct {
	UserTokensHist *Repository.UserTokensHistRepository
}

func NewUserTokensHistService(repo *Repository.UserTokensHistRepository) *UserTokensHistService {
	return &UserTokensHistService{
		UserTokensHist: repo,
	}
}

func (b *UserTokensHistService) GetUserTokensHist(ctx context.Context, payload model.PayloadDTO) error {
	convertUserid := payload.UserID

	tokenHist, err := b.UserTokensHist.GetUserTokensHist(ctx, db.GetUserTokensHistParams{
		UserID:   convertUserid,
		AccessID: payload.AccessID,
	})
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("convertUserid: %v", convertUserid)
		log.Printf("payload.AccessID: %v", payload.AccessID)
		log.Printf("payload.TenantID: %v", payload.TenantID)
		return errors.New("token not found")
	} else if err != nil {
		return err
	}

	if !tokenHist.Active {
		return errors.New("token is inactive")
	}

	if tokenHist.ExpiresAt.Before(time.Now()) {
		return errors.New("token expired")
	}

	return nil
}

func (b *UserTokensHistService) CreateUserTokensHist(ctx context.Context, payload model.PayloadDTO, tokenString string, expiresAt time.Time) error {

	err := b.UserTokensHist.CreateTokenHist(ctx, db.CreateTokenHistParams{
		UserID:    payload.UserID,
		Token:     tokenString,
		ExpiresAt: expiresAt,
		AccessID:  payload.AccessID,
	})
	if err != nil {
		return err
	}

	return nil
}
