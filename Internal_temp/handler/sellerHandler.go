package handler

import (
	"awesomeProject/Internal_temp/model"
	"awesomeProject/Internal_temp/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SellerHandler struct {
	Service *service.SellerService
}

func NewSellerHandler(s *service.SellerService) *SellerHandler {
	return &SellerHandler{Service: s}
}

type VerifySellerRequest struct {
	ActivationCode int64 `json:"activation_code"`
	CadastroId     int64 `json:"cadastro_id"`
}

func (h *SellerHandler) VerifySeller(c echo.Context) error {
	var req VerifySellerRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "requisição inválida",
		})
	}

	seller := model.Seller{
		ActivationCode: req.ActivationCode,
		CadastroId:     req.CadastroId,
	}

	status, err := h.Service.VerifySeller(c.Request().Context(), seller)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": status,
	})
}
