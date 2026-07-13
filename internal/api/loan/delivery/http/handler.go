package http

import (
	"strconv"

	loanu "github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/labstack/echo"
)

type Handler struct {
	loanUC loanu.UseCase
}

func NewHandler(loanUC loanu.UseCase) *Handler {
	return &Handler{loanUC: loanUC}
}

func (h *Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload entity.Loan
		if err := c.Bind(&payload); err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}

		err, created := h.loanUC.Create(payload)
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}
		return c.JSON(helpers.REPONSE_HTTP_OK, created)
	}
}

func (h *Handler) Return() echo.HandlerFunc {
	return func(c echo.Context) error {
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

		err, updated := h.loanUC.Return(payload.LoanId, payload.ReturnedAt)
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}
		return c.JSON(helpers.REPONSE_HTTP_OK, updated)
	}
}

func (h *Handler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		var offset int
		offsetStr := c.QueryParam("offset")
		offset, _ = strconv.Atoi(offsetStr)

		var limit int
		limitStr := c.QueryParam("limit")
		limit, _ = strconv.Atoi(limitStr)

		input := entity.LoanFilters{Offset: offset, Limit: limit}

		err, resp := h.loanUC.List(input)
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}
		return c.JSON(helpers.REPONSE_HTTP_OK, resp)
	}
}

func (h *Handler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}
		err, resp := h.loanUC.Get(id)
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}
		return c.JSON(helpers.REPONSE_HTTP_OK, resp)
	}
}

func (h *Handler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}
		err, resp := h.loanUC.Delete(id)
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}
		return c.JSON(helpers.REPONSE_HTTP_OK, resp)
	}
}
