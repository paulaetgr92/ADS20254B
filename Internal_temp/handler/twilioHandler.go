package handler

import (
	"awesomeProject/Internal_temp/model"
	"net/http"

	"awesomeProject/Internal_temp/service"
	"github.com/labstack/echo/v4"
)

type SellerHandler struct {
	Service *service.CadastroService
}

func NewSellerHandler(s *service.CadastroService) *SellerHandler {
	return &SellerHandler{Service: s}
}
func (h *SellerHandler) VerifyCode(c echo.Context) error {
	var req model.TwillioModelRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	if err := h.Service.VerifySeller(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": "Conta ativada com sucesso!",
	})
}
