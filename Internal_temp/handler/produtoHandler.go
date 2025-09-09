package handler

import (
	"awesomeProject/Internal_temp/model"
	"awesomeProject/Internal_temp/service"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProdutoHandler struct {
	Service *service.ProdutoService
}

func NewProdutoHandler(s *service.ProdutoService) *ProdutoHandler {
	return &ProdutoHandler{
		Service: s,
	}
}
func (h *ProdutoHandler) CreateProductHandler(c echo.Context) error {
	var req model.ProdutosRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos: " + err.Error(),
		})
	}

	id, err := h.Service.CreateProduct(context.Background(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao criar produto: " + err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "produto criado com sucesso",
		"id":      id,
	})
}
