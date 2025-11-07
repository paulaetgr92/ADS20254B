package handler

import (
	"context"
	"net/http"

	"awesomeProject/Internal_temp/model"
	"awesomeProject/Internal_temp/service"

	"github.com/labstack/echo/v4"
)

// Struct de resposta do seller
type SellerResponse struct {
	CadastroID     int    `json:"cadastro_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	ActivationCode string `json:"activation_code"`
}

// Handler do Seller
type SellerHandler struct {
	Service service.SellerServiceInterface
}

// Construtor
func NewSellerHandler(service service.SellerServiceInterface) *SellerHandler {
	return &SellerHandler{
		Service: service,
	}
}

// POST /sellers
func (h *SellerHandler) CreateSeller(c echo.Context) error {
	ctx := context.Background()
	var req model.SellerRequest

	// Faz o bind do JSON enviado no corpo da requisição
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Erro ao processar a requisição: " + err.Error(),
		})
	}

	// Chama o serviço para criar o seller
	sellerData, err := h.Service.CreateSeller(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao criar vendedor: " + err.Error(),
		})
	}

	sellerResp := SellerResponse{
		CadastroID:     sellerData.CadastroID,
		Name:           sellerData.Name,
		ActivationCode: sellerData.ActivationCode,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Vendedor criado com sucesso. Código de ativação enviado via SMS.",
		"seller":  sellerResp,
	})
}
