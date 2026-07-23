package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	authorMock "github.com/Leli2004/API_Go_biblioteca/internal/api/author/mocks"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/middleware"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type handlerSetup struct {
	handler *Handler
	useCase *authorMock.UseCase
	echo    *echo.Echo
	ctx     context.Context
}

func setup(t *testing.T) handlerSetup {
	useCase := authorMock.NewUseCase(t)
	e := echo.New()
	ctx := context.Background()

	return handlerSetup{
		handler: NewHandler(useCase),
		useCase: useCase,
		echo:    e,
		ctx:     ctx,
	}
}

func Test_Get_Handler(t *testing.T) {

	t.Run("Happy Path - Retorna autor", func(t *testing.T) {
		setup := setup(t)

		expected := entity.Author{
			Id:   1,
			Name: "Machado de Assis",
		}

		setup.useCase.
			EXPECT().
			Get(mock.Anything, expected.Id).
			Return(setup.ctx, nil, expected).
			Once()

		request := httptest.NewRequest(
			http.MethodGet,
			"/author/1",
			nil,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		echoCtx.SetPath("/author/:id")
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		err := setup.handler.Get()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Code)

		var result entity.Author
		err = json.Unmarshal(response.Body.Bytes(), &result)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Error Path - Id inválido", func(t *testing.T) {
		setup := setup(t)

		request := httptest.NewRequest(
			http.MethodGet,
			"/author/invalid",
			nil,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		echoCtx.SetPath("/author/:id")
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("invalid")

		err := setup.handler.Get()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		setup.useCase.AssertNotCalled(t, "Get")
	})

	t.Run("Error Path - Use case retorna erro", func(t *testing.T) {
		setup := setup(t)

		expectedError := errors.New("erro ao buscar autor")

		setup.useCase.
			EXPECT().
			Get(mock.Anything, 1).
			Return(setup.ctx, expectedError, entity.Author{}).
			Once()

		request := httptest.NewRequest(
			http.MethodGet,
			"/author/1",
			nil,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		echoCtx.SetPath("/author/:id")
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		err := setup.handler.Get()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func Test_List_Handler(t *testing.T) {

	t.Run("Happy Path - Retorna lista de autores", func(t *testing.T) {
		setup := setup(t)

		input := entity.AuthorFilters{
			Offset: 0,
			Limit:  10,
		}

		expected := entity.AuthorList{
			Offset: 0,
			Limit:  2,
			Data: []*entity.Author{
				{
					Id:   1,
					Name: "Machado de Assis",
				},
				{
					Id:   2,
					Name: "Clarice Lispector",
				},
			},
		}

		setup.useCase.
			EXPECT().
			List(mock.Anything, input).
			Return(setup.ctx, nil, expected).
			Once()

		request := httptest.NewRequest(
			http.MethodGet,
			"/author?offset=0&limit=10",
			nil,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		err := setup.handler.List()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Code)

		var result entity.AuthorList
		err = json.Unmarshal(response.Body.Bytes(), &result)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Happy Path - Parâmetros não informados", func(t *testing.T) {
		setup := setup(t)

		input := entity.AuthorFilters{}

		expected := entity.AuthorList{
			Data: []*entity.Author{},
		}

		setup.useCase.
			EXPECT().
			List(mock.Anything, input).
			Return(setup.ctx, nil, expected).
			Once()

		request := httptest.NewRequest(
			http.MethodGet,
			"/author",
			nil,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		err := setup.handler.List()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Error Path - Use case retorna erro", func(t *testing.T) {
		setup := setup(t)

		input := entity.AuthorFilters{
			Offset: 0,
			Limit:  10,
		}

		expectedError := errors.New("erro ao listar autores")

		setup.useCase.
			EXPECT().
			List(mock.Anything, input).
			Return(setup.ctx, expectedError, entity.AuthorList{}).
			Once()

		request := httptest.NewRequest(
			http.MethodGet,
			"/author?offset=0&limit=10",
			nil,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		err := setup.handler.List()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func Test_Create_Handler(t *testing.T) {

	t.Run("Happy Path - Cria autor", func(t *testing.T) {
		setup := setup(t)

		input := entity.Author{
			Name: "Conceição Evaristo",
		}

		expected := entity.Author{
			Id:   1,
			Name: input.Name,
		}

		claims := &entity.AuthClaims{
			Role: entity.RoleAdmin,
		}

		setup.useCase.
			EXPECT().
			Create(mock.Anything, input, claims).
			Return(setup.ctx, nil, expected).
			Once()

		body, err := json.Marshal(input)
		assert.NoError(t, err)

		request := httptest.NewRequest(
			http.MethodPost,
			"/author",
			bytes.NewReader(body),
		)

		request.Header.Set(
			echo.HeaderContentType,
			echo.MIMEApplicationJSON,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		echoCtx.Set(
			middleware.AuthClaimsContextKey,
			claims,
		)

		err = setup.handler.Create()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Code)

		var result entity.Author
		err = json.Unmarshal(response.Body.Bytes(), &result)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Error Path - JSON inválido", func(t *testing.T) {
		setup := setup(t)

		request := httptest.NewRequest(
			http.MethodPost,
			"/author",
			bytes.NewBufferString(`{"name":`),
		)

		request.Header.Set(
			echo.HeaderContentType,
			echo.MIMEApplicationJSON,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		echoCtx.Set(
			middleware.AuthClaimsContextKey,
			&entity.AuthClaims{
				Role: entity.RoleAdmin,
			},
		)

		err := setup.handler.Create()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		setup.useCase.AssertNotCalled(t, "Create")
	})

	t.Run("Error Path - Use case retorna erro", func(t *testing.T) {
		setup := setup(t)

		input := entity.Author{
			Name: "Conceição Evaristo",
		}

		expectedError := errors.New(
			"erro ao criar autor",
		)

		claims := &entity.AuthClaims{
			Role: entity.RoleAdmin,
		}

		setup.useCase.
			EXPECT().
			Create(mock.Anything, input, claims).
			Return(setup.ctx, expectedError, entity.Author{}).
			Once()

		body, err := json.Marshal(input)
		assert.NoError(t, err)

		request := httptest.NewRequest(
			http.MethodPost,
			"/author",
			bytes.NewReader(body),
		)

		request.Header.Set(
			echo.HeaderContentType,
			echo.MIMEApplicationJSON,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		echoCtx.Set(
			middleware.AuthClaimsContextKey,
			claims,
		)

		err = setup.handler.Create()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func Test_Update_Handler(t *testing.T) {

	t.Run("Happy Path - Atualiza autor", func(t *testing.T) {
		setup := setup(t)

		id := 1

		input := entity.Author{
			Name: "Machado de Assis Atualizado",
		}

		expected := entity.Author{
			Id:   id,
			Name: input.Name,
		}

		claims := &entity.AuthClaims{
			Role: entity.RoleAdmin,
		}

		setup.useCase.
			EXPECT().
			Update(mock.Anything, id, input, claims).
			Return(setup.ctx, nil, expected).
			Once()

		body, err := json.Marshal(input)
		assert.NoError(t, err)

		request := httptest.NewRequest(
			http.MethodPut,
			"/author/1",
			bytes.NewReader(body),
		)

		request.Header.Set(
			echo.HeaderContentType,
			echo.MIMEApplicationJSON,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		echoCtx.SetPath("/author/:id")
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		// Simula as claims inseridas pelo JWT middleware.
		echoCtx.Set(
			middleware.AuthClaimsContextKey,
			claims,
		)

		err = setup.handler.Update()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Code)

		var result entity.Author
		err = json.Unmarshal(response.Body.Bytes(), &result)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Error Path - Id inválido", func(t *testing.T) {
		setup := setup(t)

		input := entity.Author{
			Name: "Machado de Assis",
		}

		body, err := json.Marshal(input)
		assert.NoError(t, err)

		request := httptest.NewRequest(
			http.MethodPut,
			"/author/invalid",
			bytes.NewReader(body),
		)

		request.Header.Set(
			echo.HeaderContentType,
			echo.MIMEApplicationJSON,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		echoCtx.SetPath("/author/:id")
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("invalid")

		// Simula autenticação para o teste chegar à validação do ID.
		echoCtx.Set(
			middleware.AuthClaimsContextKey,
			&entity.AuthClaims{
				Role: entity.RoleAdmin,
			},
		)

		err = setup.handler.Update()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		setup.useCase.AssertNotCalled(t, "Update")
	})

	t.Run("Error Path - JSON inválido", func(t *testing.T) {
		setup := setup(t)

		request := httptest.NewRequest(
			http.MethodPut,
			"/author/1",
			bytes.NewBufferString(`{"name":`),
		)

		request.Header.Set(
			echo.HeaderContentType,
			echo.MIMEApplicationJSON,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		echoCtx.SetPath("/author/:id")
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		// Simula autenticação para o teste chegar à validação do JSON.
		echoCtx.Set(
			middleware.AuthClaimsContextKey,
			&entity.AuthClaims{
				Role: entity.RoleAdmin,
			},
		)

		err := setup.handler.Update()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		setup.useCase.AssertNotCalled(t, "Update")
	})

	t.Run("Error Path - Use case retorna erro", func(t *testing.T) {
		setup := setup(t)

		id := 1

		input := entity.Author{
			Name: "Machado de Assis Atualizado",
		}

		expectedError := errors.New(
			"erro ao atualizar autor",
		)

		claims := &entity.AuthClaims{
			Role: entity.RoleAdmin,
		}

		setup.useCase.
			EXPECT().
			Update(mock.Anything, id, input, claims).
			Return(setup.ctx, expectedError, entity.Author{}).
			Once()

		body, err := json.Marshal(input)
		assert.NoError(t, err)

		request := httptest.NewRequest(
			http.MethodPut,
			"/author/1",
			bytes.NewReader(body),
		)

		request.Header.Set(
			echo.HeaderContentType,
			echo.MIMEApplicationJSON,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		echoCtx.SetPath("/author/:id")
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		// Simula autenticação para o teste chegar ao use case.
		echoCtx.Set(
			middleware.AuthClaimsContextKey,
			claims,
		)

		err = setup.handler.Update()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func Test_Delete_Handler(t *testing.T) {

	t.Run("Happy Path - Exclui autor", func(t *testing.T) {
		setup := setup(t)

		id := 1

		claims := &entity.AuthClaims{
			Role: entity.RoleAdmin,
		}

		setup.useCase.
			EXPECT().
			Delete(mock.Anything, id, claims).
			Return(setup.ctx, nil).
			Once()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/author/1",
			nil,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		echoCtx.SetPath("/author/:id")
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		echoCtx.Set(
			middleware.AuthClaimsContextKey,
			claims,
		)

		err := setup.handler.Delete()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Error Path - Id inválido", func(t *testing.T) {
		setup := setup(t)

		request := httptest.NewRequest(
			http.MethodDelete,
			"/author/invalid",
			nil,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		echoCtx.SetPath("/author/:id")
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("invalid")

		echoCtx.Set(
			middleware.AuthClaimsContextKey,
			&entity.AuthClaims{
				Role: entity.RoleAdmin,
			},
		)

		err := setup.handler.Delete()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		setup.useCase.AssertNotCalled(t, "Delete")
	})

	t.Run("Error Path - Use case retorna erro", func(t *testing.T) {
		setup := setup(t)

		id := 1

		expectedError := errors.New(
			"erro ao excluir autor",
		)

		claims := &entity.AuthClaims{
			Role: entity.RoleAdmin,
		}

		setup.useCase.
			EXPECT().
			Delete(mock.Anything, id, claims).
			Return(setup.ctx, expectedError).
			Once()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/author/1",
			nil,
		)

		response := httptest.NewRecorder()
		echoCtx := setup.echo.NewContext(request, response)

		echoCtx.SetPath("/author/:id")
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		echoCtx.Set(
			middleware.AuthClaimsContextKey,
			claims,
		)

		err := setup.handler.Delete()(echoCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}
