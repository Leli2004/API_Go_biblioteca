package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/genre"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/Leli2004/API_Go_biblioteca/internal/middleware"
	"github.com/labstack/echo"
)

type Handler struct {
	genreUC genre.UseCase
}

func NewHandler(genreUC genre.UseCase) *Handler {
	return &Handler{genreUC: genreUC}
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

		input := entity.GenreFilters{
			Offset: offset,
			Limit:  limit,
		}

		ctx, err, resp := h.genreUC.List(ctx, input)
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

		ctx, err, resp := h.genreUC.Get(ctx, id)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result") {
				return c.JSON(http.StatusNotFound, map[string]string{"error": "genre not found"})
			}
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		return c.JSON(helpers.REPONSE_HTTP_OK, resp)
	}
}

func (h *Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := helpers.InitCtx(c)
		var input entity.Genre
		if err := c.Bind(&input); err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		claims, err := middleware.GetAuthClaims(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		ctx, err, resp := h.genreUC.Create(ctx, input, claims)
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

		var input entity.Genre
		if err := c.Bind(&input); err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		claims, err := middleware.GetAuthClaims(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		ctx, err, resp := h.genreUC.Update(ctx, id, input, claims)
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

		claims, err := middleware.GetAuthClaims(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		ctx, err = h.genreUC.Delete(ctx, id, claims)

		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		return c.JSON(helpers.REPONSE_HTTP_OK, map[string]string{"status": "deleted"})
	}
}
