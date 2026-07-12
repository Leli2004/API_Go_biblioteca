package http

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/book"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Echo, h book.Handler) {
	e.GET("book/list", h.List())
	e.GET("book/:id", h.Get())
	e.POST("book", h.Create())
	e.PUT("book/:id", h.Update())
	e.DELETE("book/:id", h.Delete())
}
