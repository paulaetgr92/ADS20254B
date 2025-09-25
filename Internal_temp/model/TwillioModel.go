package model

type TwillioModelRequest struct {
	Payload        PayloadDTO `json:"payload"`
	ActivationCode string     `json:"code"`
	Status         string     `json:"status"`
	Celular        string     `json:"Celular"`
}
