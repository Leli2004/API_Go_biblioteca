package helpers

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/labstack/echo"
)

func ResponseErrorHTTP(c echo.Context, code int, err error) error {
	if err == nil {
		return nil
	}

	return c.JSON(REPONSE_HTTP_BAD_REQUEST, entity.RespError{
		Code: code,
		Err:  err.Error(),
	})
}
