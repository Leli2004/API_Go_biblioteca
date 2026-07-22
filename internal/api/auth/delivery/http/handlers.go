package http

import (
	"errors"
	"net/http"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/auth"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	appmiddleware "github.com/Leli2004/API_Go_biblioteca/internal/middleware"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/labstack/echo"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) auth.Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input entity.LoginRequest

		if err := c.Bind(&input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
		}

		_, err, result := h.useCase.Login(c.Request().Context(), input)
		if err != nil {
			if errors.Is(err, security.ErrInvalidCredentials) {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid username or password")
			}

			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	}
}

func (h *Handler) Me() echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := appmiddleware.GetAuthClaims(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		_, err, result := h.useCase.Me(c.Request().Context(), claims.UserId)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	}
}
