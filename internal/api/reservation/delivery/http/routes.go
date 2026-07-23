package http

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/reservation"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Group, h reservation.Handler) {
	e.POST("/create", h.Create())
}
