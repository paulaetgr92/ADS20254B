package model

import "time"

type AdminRequest struct {
	Id        int64     `json:"id"`
	Cnpj      string    `json:"cnpj"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type ModelRequest struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}
