package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/book/usecase"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/labstack/echo"
)

type Handler struct {
	bookUC usecase.BookUC
}

func NewHandler(bookUC usecase.BookUC) *Handler {
	return &Handler{bookUC: bookUC}
}

func (h *Handler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := helpers.InitCtx(c)
		var offset int
		offsetStr := c.QueryParam("offset")
		offset, _ = strconv.Atoi(offsetStr)

		var limit int
		limitStr := c.QueryParam("limit")
		limit, _ = strconv.Atoi(limitStr)

		input := entity.BookFilters{
			Offset: offset,
			Limit:  limit,
		}

		ctx, err, resp := h.bookUC.List(ctx, input)
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		return c.JSON(helpers.REPONSE_HTTP_OK, resp)
	}
}

func (h *Handler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := helpers.InitCtx(c)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		ctx, err, resp := h.bookUC.Get(ctx, id)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result") {
				return c.JSON(http.StatusNotFound, map[string]string{"error": "book not found"})
			}
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		return c.JSON(helpers.REPONSE_HTTP_OK, resp)
	}
}

func (h *Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := helpers.InitCtx(c)
		var input entity.Book
		if err := c.Bind(&input); err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		ctx, err, resp := h.bookUC.Create(ctx, input)
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		return c.JSON(helpers.REPONSE_HTTP_OK, resp)
	}
}

func (h *Handler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := helpers.InitCtx(c)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		var input entity.Book
		if err := c.Bind(&input); err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		ctx, err, resp := h.bookUC.Update(ctx, id, input)
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		return c.JSON(helpers.REPONSE_HTTP_OK, resp)
	}
}

func (h *Handler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := helpers.InitCtx(c)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		ctx, err = h.bookUC.Delete(ctx, id)

		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		return c.JSON(helpers.REPONSE_HTTP_OK, map[string]string{"status": "deleted"})
	}
}
