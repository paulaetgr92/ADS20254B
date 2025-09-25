package model

import "database/sql"

type LoginRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	Status         string `json:"status"`
	ActivationCode string `json:"activationCode"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type GetLoginRow struct {
	Email          string
	Password       string
	ActivationCode sql.NullString
	Status         string
}
