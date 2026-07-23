package http

import (
	"net/http"
	"strconv"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/Leli2004/API_Go_biblioteca/internal/middleware"
	"github.com/labstack/echo"
)

type Handler struct {
	loanUC loan.UseCase
}

func NewHandler(loanUC loan.UseCase) *Handler {
	return &Handler{loanUC: loanUC}
}

func (h *Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := helpers.InitCtx(c)
		var payload entity.Loan
		if err := c.Bind(&payload); err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		claims, err := middleware.GetAuthClaims(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		ctx, err, created := h.loanUC.Create(ctx, payload, claims)
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}
		return c.JSON(helpers.REPONSE_HTTP_OK, created)
	}
}

func (h *Handler) Return() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := helpers.InitCtx(c)
		var payload struct {
			LoanId     int     `json:"loan_id"`
			ReturnedAt *string `json:"returned_at"`
		}
		if err := c.Bind(&payload); err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}
		if payload.LoanId <= 0 {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, echo.NewHTTPError(400, "loan_id is required"))
		}

		claims, err := middleware.GetAuthClaims(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		ctx, err, updated := h.loanUC.Return(ctx, payload.LoanId, payload.ReturnedAt, claims)
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}
		return c.JSON(helpers.REPONSE_HTTP_OK, updated)
	}
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

		input := entity.LoanFilters{Offset: offset, Limit: limit}

		ctx, err, resp := h.loanUC.List(ctx, input)
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
		ctx, err, resp := h.loanUC.Get(ctx, id)
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

		ctx, err, resp := h.loanUC.Delete(ctx, id, claims)
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		return c.JSON(helpers.REPONSE_HTTP_OK, resp)
	}
}
