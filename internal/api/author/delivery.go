package author

import "github.com/labstack/echo"

type Handler interface {
	List() echo.HandlerFunc
}
