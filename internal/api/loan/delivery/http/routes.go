package http

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Group, h loan.Handler) {
	e.POST("/create", h.Create())
	e.POST("/return", h.Return())
	e.GET("/list", h.List())
	e.GET("/:id", h.Get())
	e.DELETE("/:id", h.Delete())
}
