package model

type CadastroRequest struct {
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	PayloadDTO PayloadDTO `json:"payloadDTO"` // Payload completo
}
