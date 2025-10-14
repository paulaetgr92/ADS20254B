package handler

import (
	"awesomeProject/Internal_temp/service"
	"net/http"

	"awesomeProject/Internal_temp/model"

	"github.com/labstack/echo/v4"
)

type SaleHandler struct {
	Service service.SaleServiceInterface
}

func NewSaleHandler(s *service.SaleService) *SaleHandler {
	return &SaleHandler{Service: s}
}

func (h *SaleHandler) CreateSale(c echo.Context) error {
	var req model.SaleRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Erro ao ler corpo da requisição",
		})
	}

	sale, err := h.Service.CreateSale(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao criar venda",
		})
	}
	return c.JSON(http.StatusCreated, sale)
}
