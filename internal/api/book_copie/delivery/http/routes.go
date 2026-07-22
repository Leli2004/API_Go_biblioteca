package http

import (
	book_copie "github.com/Leli2004/API_Go_biblioteca/internal/api/book_copie"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Group, h book_copie.Handler) {
	e.GET("/list", h.List())
	e.GET("/:id", h.Get())
	e.POST("", h.Create())
	e.PUT("/:id", h.Update())
	e.DELETE("/:id", h.Delete())
}
