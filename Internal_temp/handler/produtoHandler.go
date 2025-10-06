package handler

import (
	"awesomeProject/Internal_temp/model"
	"awesomeProject/Internal_temp/service"
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

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

func (h *ProdutoHandler) GetProductByIdHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
	}

	produto, err := h.Service.GetProdutoByIdService(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Produto não encontrado",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao buscar produto",
		})
	}

	return c.JSON(http.StatusOK, produto)
}

func (h *ProdutoHandler) ListProdutosHandler(c echo.Context) error {
	produtos, err := h.Service.ListProdutoService(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao listar produtos"})
	}
	return c.JSON(http.StatusOK, produtos)
}

func (h *ProdutoHandler) UpdateProdutoByIdHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	var req model.ProdutosRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
	}

	produtoAtualizado, err := h.Service.UpdateProdutoByIdService(c.Request().Context(), id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, produtoAtualizado)
}

func (h *ProdutoHandler) InativarProdutoHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	err = h.Service.DeleteProdutoByIdService(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Produto inativado com sucesso"})
}
