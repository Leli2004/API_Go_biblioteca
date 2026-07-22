package http

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/auth"
	"github.com/labstack/echo"
)

func MapPublicRoutes(e *echo.Echo, handler auth.Handler) {
	e.POST("/auth/login", handler.Login())
}

func MapProtectedRoutes(group *echo.Group, handler auth.Handler) {
	group.GET("/auth/me", handler.Me())
}
