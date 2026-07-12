package http

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/user"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Echo, h user.Handler) {
	e.GET("user/list", h.List())
	e.GET("user/:id", h.Get())
	e.POST("user", h.Create())
	e.PUT("user/:id", h.Update())
	e.DELETE("user/:id", h.Delete())
}
