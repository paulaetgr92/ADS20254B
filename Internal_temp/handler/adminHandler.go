package handler

import (
	"awesomeProject/Internal_temp/model"
	"awesomeProject/Internal_temp/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	Service service.AdminServiceInterface
}

func NewAdminHandler(s service.AdminServiceInterface) *AdminHandler {
	return &AdminHandler{Service: s}
}

func (h *AdminHandler) CreateAdminHandler(c echo.Context) error {
	var req model.AdminRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err := h.Service.CreateAdminService(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"success": "success",
	})
}

func (h *AdminHandler) CreateLoginAdminHandler(c echo.Context) error {
	var req model.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err := h.Service.LoginAdminService(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"success": "success",
	})
}

func (h *AdminHandler) ListAllProdutosAdminHandler(c echo.Context) error {
	produtos, err := h.Service.ListAllProdutosService(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao listar produtos para o admin"})
	}
	return c.JSON(http.StatusOK, produtos)
}

