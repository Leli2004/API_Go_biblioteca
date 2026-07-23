package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	authMocks "github.com/Leli2004/API_Go_biblioteca/internal/api/auth/mocks"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	appMiddleware "github.com/Leli2004/API_Go_biblioteca/internal/middleware"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginHandler(t *testing.T) {
	
	t.Run("Happy Path - retorna token", func(t *testing.T) {
		e := echo.New()
		uc := authMocks.NewUseCase(t)
		h := NewHandler(uc)

		input := entity.LoginRequest{
			Username: "admin",
			Password: "senha123",
		}
		expected := entity.LoginResponse{
			Token:     "jwt-token",
			TokenType: security.TokenTypeBearer,
			ExpiresIn: 86400,
			User: entity.AuthUser{
				Id:       1,
				Name:     "Administrador",
				Username: "admin",
				Email:    "admin@biblioteca.com",
				Role:     entity.RoleAdmin,
			},
		}

		body, err := json.Marshal(input)
		assert.NoError(t, err)

		uc.On("Login", mock.Anything, input).
			Return(context.Background(), nil, expected).
			Once()

		req := httptest.NewRequest(
			http.MethodPost,
			"/auth/login",
			bytes.NewReader(body),
		)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		err = h.Login()(e.NewContext(req, rec))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "jwt-token")
	})

	t.Run("Error Path - JSON invalido", func(t *testing.T) {
		e := echo.New()
		uc := authMocks.NewUseCase(t)
		h := NewHandler(uc)

		req := httptest.NewRequest(
			http.MethodPost,
			"/auth/login",
			bytes.NewBufferString(`{"username":`),
		)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		err := h.Login()(e.NewContext(req, rec))

		assert.Error(t, err)
		uc.AssertNotCalled(t, "Login")
	})

	t.Run("Error Path - credenciais invalidas", func(t *testing.T) {
		e := echo.New()
		uc := authMocks.NewUseCase(t)
		h := NewHandler(uc)

		input := entity.LoginRequest{
			Username: "admin",
			Password: "senha-errada",
		}
		body, err := json.Marshal(input)
		assert.NoError(t, err)

		uc.On("Login", mock.Anything, input).
			Return(context.Background(), security.ErrInvalidCredentials, entity.LoginResponse{}).
			Once()

		req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		err = h.Login()(e.NewContext(req, rec))

		httpError, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusUnauthorized, httpError.Code)
	})
}

func TestMeHandler(t *testing.T) {

	t.Run("Happy Path - retorna usuario autenticado", func(t *testing.T) {
		e := echo.New()
		uc := authMocks.NewUseCase(t)
		h := NewHandler(uc)

		expected := entity.AuthUser{
			Id:       1,
			Name:     "Administrador",
			Username: "admin",
			Email:    "admin@biblioteca.com",
			Role:     entity.RoleAdmin,
		}

		uc.On("Me", mock.Anything, 1).
			Return(context.Background(), nil, expected).
			Once()

		req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set(appMiddleware.AuthClaimsContextKey, &entity.AuthClaims{
			UserId:   1,
			Username: "admin",
			Role:     entity.RoleAdmin,
		})

		err := h.Me()(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "admin")
	})

	t.Run("Error Path - claims ausentes", func(t *testing.T) {
		e := echo.New()
		uc := authMocks.NewUseCase(t)
		h := NewHandler(uc)

		req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
		rec := httptest.NewRecorder()

		err := h.Me()(e.NewContext(req, rec))

		httpError, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusUnauthorized, httpError.Code)
		uc.AssertNotCalled(t, "Me")
	})

	t.Run("Error Path - usecase retorna erro", func(t *testing.T) {
		e := echo.New()
		uc := authMocks.NewUseCase(t)
		h := NewHandler(uc)
		expectedError := errors.New("user not found")

		uc.On("Me", mock.Anything, 1).
			Return(context.Background(), expectedError, entity.AuthUser{}).
			Once()

		req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set(appMiddleware.AuthClaimsContextKey, &entity.AuthClaims{
			UserId:   1,
			Username: "admin",
			Role:     entity.RoleAdmin,
		})

		err := h.Me()(c)

		assert.Error(t, err)
	})
}
