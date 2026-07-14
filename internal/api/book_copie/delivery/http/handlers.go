package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/book_copie/usecase"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/labstack/echo"
)

type Handler struct {
	bookCopieUC usecase.BookCopieUC
}

func NewHandler(bookCopieUC usecase.BookCopieUC) *Handler {
	return &Handler{bookCopieUC: bookCopieUC}
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

		input := entity.BookCopyFilters{
			Offset: offset,
			Limit:  limit,
		}

		ctx, err, resp := h.bookCopieUC.List(ctx, input)
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

		ctx, err, resp := h.bookCopieUC.Get(ctx, id)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result") {
				return c.JSON(http.StatusNotFound, map[string]string{"error": "book copie not found"})
			}
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		return c.JSON(helpers.REPONSE_HTTP_OK, resp)
	}
}

func (h *Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := helpers.InitCtx(c)
		var input entity.BookCopy
		if err := c.Bind(&input); err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		ctx, err, resp := h.bookCopieUC.Create(ctx, input)
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

		var input entity.BookCopy
		if err := c.Bind(&input); err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		ctx, err, resp := h.bookCopieUC.Update(ctx, id, input)
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

		ctx, err = h.bookCopieUC.Delete(ctx, id)

		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		return c.JSON(helpers.REPONSE_HTTP_OK, map[string]string{"status": "deleted"})
	}
}
