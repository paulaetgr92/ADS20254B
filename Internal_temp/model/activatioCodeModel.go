package model

type ActivationCode struct {
	CadastroID     int    `db:"cadastro_id"`
	Code           string `db:"code"`
	ActivationCode string `db:"activation_code"`
}

type ActivationCodeDTO struct {
	ActivationCode string `db:"activation_code"`
	CadastroID     int    `db:"cadastro_id"`
	Code           string `db:"code"`
}
