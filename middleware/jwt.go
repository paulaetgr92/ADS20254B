package middleware

import (
	"awesomeProject/Internal_temp/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTKey = []byte("sua_chave_super_secreta")

func GenerateJWT(payload model.PayloadDTO) (string, time.Time, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id":   payload.UserID,
		"access_id": payload.AccessID,
		"tenant_id": payload.TenantID,
		"exp":       expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwt.SigningMethodHS256)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}
