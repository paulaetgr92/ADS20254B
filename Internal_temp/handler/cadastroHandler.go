package handler

import (
	"awesomeProject/Internal_temp/model"
	"awesomeProject/Internal_temp/service"
	_ "awesomeProject/Internal_temp/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CadastroHandler struct {
	service service.CadastroServiceInterface
}

func NewCadastroHandler(s service.CadastroServiceInterface) *CadastroHandler {
	return &CadastroHandler{
		service: s,
	}
}

func (h *CadastroHandler) CreateCadastro(c echo.Context) error {
	var req model.CadastroRequest

	// Faz o bind do JSON para a struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "requisição inválida",
		})
	}

	// Chama a service
	if err := h.service.CreateCadastro(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// Retorna sucesso
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Cadastro criado com sucesso",
	})
}
