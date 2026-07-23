package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	reservationMock "github.com/Leli2004/API_Go_biblioteca/internal/api/reservation/mocks"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/middleware"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type handlerSetup struct {
	handler *Handler
	useCase *reservationMock.UseCase
	echo    *echo.Echo
	ctx     context.Context
}

func setupHandler(t *testing.T) handlerSetup {
	useCase := reservationMock.NewUseCase(t)
	e := echo.New()
	return handlerSetup{NewHandler(useCase), useCase, e, context.Background()}
}

func Test_Create_Handler(t *testing.T) {

	t.Run("Happy Path - Cria reserva", func(t *testing.T) {
		s := setupHandler(t)

		input := entity.Reservation{
			UserId: 1,
			BookId: 2,
		}

		position := 1

		expected := entity.Reservation{
			Id:       10,
			UserId:   1,
			BookId:   2,
			Position: &position,
			Status:   entity.ReservationStatusWaiting,
		}

		claims := &entity.AuthClaims{
			Role: entity.RoleAdmin,
		}

		body, err := json.Marshal(input)
		assert.NoError(t, err)

		s.useCase.
			On("Create", mock.Anything, input, claims).
			Return(s.ctx, nil, expected).
			Once()

		req := httptest.NewRequest(
			http.MethodPost,
			"/reservation/create",
			bytes.NewReader(body),
		)

		req.Header.Set(
			echo.HeaderContentType,
			echo.MIMEApplicationJSON,
		)

		response := httptest.NewRecorder()
		echoCtx := s.echo.NewContext(req, response)

		echoCtx.Set(
			middleware.AuthClaimsContextKey,
			claims,
		)

		err = s.handler.Create()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Code)

		var result entity.Reservation

		assert.NoError(
			t,
			json.Unmarshal(
				response.Body.Bytes(),
				&result,
			),
		)

		assert.Equal(t, expected, result)
	})

	t.Run("Error Path - JSON inválido", func(t *testing.T) {
		s := setupHandler(t)

		req := httptest.NewRequest(
			http.MethodPost,
			"/reservation/create",
			bytes.NewBufferString(`{"user_id":`),
		)

		req.Header.Set(
			echo.HeaderContentType,
			echo.MIMEApplicationJSON,
		)

		response := httptest.NewRecorder()
		echoCtx := s.echo.NewContext(req, response)

		echoCtx.Set(
			middleware.AuthClaimsContextKey,
			&entity.AuthClaims{
				Role: entity.RoleAdmin,
			},
		)

		err := s.handler.Create()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)

		s.useCase.AssertNotCalled(
			t,
			"Create",
		)
	})

	t.Run("Error Path - Use case retorna erro", func(t *testing.T) {
		s := setupHandler(t)

		input := entity.Reservation{
			UserId: 1,
			BookId: 2,
		}

		claims := &entity.AuthClaims{
			Role: entity.RoleAdmin,
		}

		body, err := json.Marshal(input)
		assert.NoError(t, err)

		expectedError := errors.New(
			"reservation error",
		)

		s.useCase.
			On("Create", mock.Anything, input, claims).
			Return(
				s.ctx,
				expectedError,
				entity.Reservation{},
			).
			Once()

		req := httptest.NewRequest(
			http.MethodPost,
			"/reservation/create",
			bytes.NewReader(body),
		)

		req.Header.Set(
			echo.HeaderContentType,
			echo.MIMEApplicationJSON,
		)

		response := httptest.NewRecorder()
		echoCtx := s.echo.NewContext(req, response)

		echoCtx.Set(
			middleware.AuthClaimsContextKey,
			claims,
		)

		err = s.handler.Create()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}
