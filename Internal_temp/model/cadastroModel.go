package model

type CadastroRequest struct {
	Name string `json:"name"`

	CPF        string     `json:"cpf"`
	CNPJ       string     `json:"cnpj"`
	Celular    string     `json:"celular"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	PayloadDTO PayloadDTO `json:"payloadDTO"`
}
