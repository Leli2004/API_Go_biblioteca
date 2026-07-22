package http

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/genre"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Echo, h genre.Handler) {
	e.GET("genre/list", h.List())
	e.GET("genre/:id", h.Get())
	e.POST("genre", h.Create())
	e.PUT("genre/:id", h.Update())
	e.DELETE("genre/:id", h.Delete())
}
