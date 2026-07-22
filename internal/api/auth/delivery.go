package auth

import "github.com/labstack/echo"

type Handler interface {
	Login() echo.HandlerFunc
	Me() echo.HandlerFunc
}
