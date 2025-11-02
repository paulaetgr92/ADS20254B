package handler

import (
	"awesomeProject/Internal_temp/model"
	"awesomeProject/Internal_temp/service"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

// SellerHandler gerencia as rotas de Seller
type SellerHandler struct {
	Service *service.SellerService
}

// Novo handler
func NewSellerHandler(svc *service.SellerService) *SellerHandler {
	return &SellerHandler{
		Service: svc,
	}
}

// CreateSeller cria um novo seller, salva o código de ativação e envia SMS
func (h *SellerHandler) CreateSeller(c echo.Context) error {
	ctx := context.Background()

	var req model.Seller
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	seller, err := h.Service.CreateSeller(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// Retorna informações do seller (sem expor senha)
	resp := map[string]interface{}{
		"id":               seller.CadastroId,
		"name":             seller.Name,
		"email":            seller.Email,
		"phone":            seller.Phone,
		"status":           seller.Status,
		"activation_codes": seller.ActivationCodes, // opcional mostrar
	}

	return c.JSON(http.StatusCreated, resp)
}

// VerifySeller valida o código de ativação e ativa a conta
func (h *SellerHandler) VerifySeller(c echo.Context) error {
	ctx := context.Background()

	type VerifyRequest struct {
		CadastroID int64  `json:"cadastro_id"`
		Code       string `json:"code"`
	}

	var req VerifyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	status, err := h.Service.VerifySeller(ctx, req.CadastroID, req.Code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": status,
	})
}
