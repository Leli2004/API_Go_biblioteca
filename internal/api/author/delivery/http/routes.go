package http

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/author"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Echo, h author.Handler) {
	e.GET("author/list", h.List())
}
