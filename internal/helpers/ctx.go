package helpers

import (
	"context"

	"github.com/labstack/echo"
)

func InitCtx(c echo.Context) context.Context {
	return c.Request().Context()
}
