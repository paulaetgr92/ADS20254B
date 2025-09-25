package handler

import (
	"awesomeProject/Internal_temp/model"
	"awesomeProject/Internal_temp/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CadastroHandler struct {
	service *service.CadastroService
}

func NewCadastroHandler(svc *service.CadastroService) *CadastroHandler {
	return &CadastroHandler{
		service: svc,
	}
}

func (h *CadastroHandler) CreateCadastro(c echo.Context) error {
	var req model.CadastroRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	if req.Name == "" || req.Celular == "" || req.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "nome, celular e email são obrigatórios",
		})
	}

	cadastro, err := h.service.CreateCadastro(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":  "Cadastro realizado com sucesso! Verifique seu WhatsApp para ativar sua conta.",
		"cadastro": cadastro,
	})
}
