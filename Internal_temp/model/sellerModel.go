package model

import "time"

type Seller struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Cpf             string `json:"cpf"`
	Cnpj            string `json:"cnpj"`
	ActivationCodes string `json:"activation_code"`
	CadastroId      int64  `json:"cadastro_id"`
	Phone           string `json:"phone"`
	Status          string `json:"status"`
}

type SaveActivationCodeParams struct {
	CadastroID      int64     // ID do cadastro
	ActivationCodes string    // código de ativação (string)
	Code            string    // código de ativação real (string, para coluna NOT NULL)
	ExpiresAt       time.Time // data/hora de expiração do código
	Status          string    // status: "pendente", "ativo", etc.
}
