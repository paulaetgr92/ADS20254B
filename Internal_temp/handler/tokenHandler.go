package handler

import (
	"awesomeProject/Internal_temp/service"
	helper "awesomeProject/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserTokensHistHandler struct {
	UserTokensHist service.TokenServiceInterface
}

func NewUserTokensHistHandler(userTokensHist service.TokenServiceInterface) *UserTokensHistHandler {
	return &UserTokensHistHandler{
		UserTokensHist: userTokensHist,
	}
}

func (h *UserTokensHistHandler) GetUserTokensHist(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := helper.GetPayloadToken(c)

		err := h.UserTokensHist.GetUserTokensHist(c.Request().Context(), payload)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		}

		return next(c)
	}
}

func (h *UserTokensHistHandler) CreateUserToken(c echo.Context) error {

	payload := helper.GetPayloadToken(c)

	tokenString, expiresAt, err := helper.GenerateJWT(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token":      tokenString,
		"expires_at": expiresAt,
	})
}
