package http

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/api/reservation"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/labstack/echo"
)

type Handler struct{ reservationUC reservation.UseCase }

func NewHandler(uc reservation.UseCase) *Handler { return &Handler{reservationUC: uc} }

func (h *Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := helpers.InitCtx(c)
		var payload entity.Reservation
		if err := c.Bind(&payload); err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}
		_, err, created := h.reservationUC.Create(ctx, payload)
		if err != nil {
			return helpers.ResponseErrorHTTP(c, helpers.REPONSE_HTTP_BAD_REQUEST, err)
		}
		return c.JSON(helpers.REPONSE_HTTP_OK, created)
	}
}
