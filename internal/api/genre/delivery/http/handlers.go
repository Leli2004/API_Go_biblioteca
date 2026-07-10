package http

import (
	"strconv"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/genre/usecase"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/labstack/echo"
)

type Handler struct {
	genreUC usecase.GenreUC
}

func NewHandler(genreUC usecase.GenreUC) *Handler {
	return &Handler{genreUC: genreUC}
}

func (h *Handler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		var offset int
		offsetStr := c.QueryParam("offset")
		offset, _ = strconv.Atoi(offsetStr)

		var limit int
		limitStr := c.QueryParam("limit")
		limit, _ = strconv.Atoi(limitStr)

		input := entity.GenreFilters{
			Offset: offset,
			Limit:  limit,
		}

		err, resp := h.genreUC.List(input)
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		return c.JSON(helpers.REPONSE_HTTP_OK, resp)
	}
}
