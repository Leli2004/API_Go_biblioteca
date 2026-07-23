package http

import (
	"context"
	mm "github.com/Leli2004/API_Go_biblioteca/internal/api/genre/mocks"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_List_Handler(t *testing.T) {
	e := echo.New()
	u := mm.NewUseCase(t)
	h := NewHandler(u)
	expected := entity.GenreList{}
	u.On("List", mock.Anything, entity.GenreFilters{Offset: 2, Limit: 5}).Return(context.Background(), nil, expected).Once()
	req := httptest.NewRequest(http.MethodGet, "/?offset=2&limit=5", nil)
	rec := httptest.NewRecorder()
	err := h.List()(e.NewContext(req, rec))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func Test_Get_Handler_InvalidID(t *testing.T) {
	e := echo.New()
	u := mm.NewUseCase(t)
	h := NewHandler(u)
	req := httptest.NewRequest(http.MethodGet, "/x", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("x")
	err := h.Get()(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	u.AssertNotCalled(t, "Get")
}

func Test_Delete_Handler_InvalidID(t *testing.T) {
	e := echo.New()
	u := mm.NewUseCase(t)
	h := NewHandler(u)
	req := httptest.NewRequest(http.MethodDelete, "/x", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("x")
	err := h.Delete()(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	u.AssertNotCalled(t, "Delete")
}
