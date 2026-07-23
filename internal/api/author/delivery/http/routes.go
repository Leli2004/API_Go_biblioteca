package http

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/author"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Group, h author.Handler) {
	e.GET("/list", h.List())
	e.GET("/:id", h.Get())
	e.POST("", h.Create())
	e.PUT("/:id", h.Update())
	e.DELETE("/:id", h.Delete())
}
