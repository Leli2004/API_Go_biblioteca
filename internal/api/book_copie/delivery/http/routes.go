package http

import (
	book_copie "github.com/Leli2004/API_Go_biblioteca/internal/api/book_copie"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Echo, h book_copie.Handler) {
	e.GET("book_copie/list", h.List())
	e.GET("book_copie/:id", h.Get())
	e.POST("book_copie", h.Create())
	e.PUT("book_copie/:id", h.Update())
	e.DELETE("book_copie/:id", h.Delete())
}
