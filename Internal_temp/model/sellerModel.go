package model

import (
	"time"
)

type SaveActivationCodeParams struct {
	CadastroID      int64     // ID do cadastro
	ActivationCodes string    // código de ativação (string)
	Code            string    // código de ativação real (string, para coluna NOT NULL)
	ExpiresAt       time.Time // data/hora de expiração do código
	Status          string    // status: "pendente", "ativo", etc.
}

type SellerRequest struct {
	ID             int64  `json:"id"`
	Password       string `json:"password"`
	CPF            string `json:"cpf"`
	CNPJ           string `json:"cnpj"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Celular        string `json:"celular"`
	Status         string `json:"status"`
	ActivationCode string `json:"activation_code"`
}

type SellerResponse struct {
	CadastroID     int    `json:"cadastro_id"`
	ActivationCode string `json:"activation_code"`
	Name           string `json:"name"`
}
