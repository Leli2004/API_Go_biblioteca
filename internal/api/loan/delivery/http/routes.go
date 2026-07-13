package http

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Echo, h loan.Handler) {
	e.POST("loan/create", h.Create())
	e.POST("loan/return", h.Return())
	e.GET("loan/list", h.List())
	e.GET("loan/:id", h.Get())
	e.DELETE("loan/:id", h.Delete())
}
