package loan

import "github.com/labstack/echo"

type Handler interface {
	Create() echo.HandlerFunc
	Return() echo.HandlerFunc
	List() echo.HandlerFunc
	Get() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
