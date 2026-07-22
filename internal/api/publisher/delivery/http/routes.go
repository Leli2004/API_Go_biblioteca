package http

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/publisher"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Echo, h publisher.Handler) {
	e.GET("publisher/list", h.List())
	e.GET("publisher/:id", h.Get())
	e.POST("publisher", h.Create())
	e.PUT("publisher/:id", h.Update())
	e.DELETE("publisher/:id", h.Delete())
}
