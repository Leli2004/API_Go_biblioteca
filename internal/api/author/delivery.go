package author

import "github.com/labstack/echo"

type Handler interface {
	List() echo.HandlerFunc
	Get() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
