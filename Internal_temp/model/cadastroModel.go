package model

type CadastroRequest struct {
	Name           string     `json:"name"`
	ActivationCode string     `json:"activationCode"`
	CNPJ           string     `json:"cnpj"`
	Celular        string     `json:"celular"`
	Email          string     `json:"email"`
	Password       string     `json:"password"`
	Status         string     `json:"status"`
	PayloadDTO     PayloadDTO `json:"payloadDTO"` // Payload completo
}
