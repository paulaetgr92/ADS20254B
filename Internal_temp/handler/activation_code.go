package handler

import (
	"awesomeProject/Internal_temp/model"
	"awesomeProject/Internal_temp/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ActivationHandler struct {
	Service *service.ActivationService
}

func NewActivationHandler(s *service.ActivationService) *ActivationHandler {
	return &ActivationHandler{Service: s}
}

func (h *ActivationHandler) VerifyActivationCode(c echo.Context) error {
	var req model.ActivationCode

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Requisição inválida"})
	}

	row, err := h.Service.VerifyActivationCode(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Código inválido ou expirado"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Código verificado com sucesso",
		"cadastro_id":   row.CadastroID,
		"activation_id": row.ID,
	})
}

func (h *ActivationHandler) SaveActivationCode(c echo.Context) error {
	var req model.ActivationCode
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Requisição inválida"})
	}

	savedCode, err := h.Service.SaveActivationCode(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, savedCode)
}

type GetActivationCodeRequest struct {
	CadastroID     int64  `json:"cadastro_id"`
	ActivationCode string `json:"activation_code"`
}

func (h *ActivationHandler) GetActivationCode(c echo.Context) error {
	var req GetActivationCodeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Requisição inválida",
		})
	}

	ctx := c.Request().Context()

	codeRow, err := h.Service.GetActivationCode(ctx, req.ActivationCode, req.CadastroID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, codeRow)
}
