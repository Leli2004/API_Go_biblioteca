package http

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/genre"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Echo, h genre.Handler) {
	e.GET("genre/list", h.List())
}
