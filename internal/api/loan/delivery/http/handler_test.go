package http

import (
	"context"
	mm "github.com/Leli2004/API_Go_biblioteca/internal/api/loan/mocks"
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
	u.On("List", mock.Anything, entity.LoanFilters{Offset: 1, Limit: 2}).Return(context.Background(), nil, entity.LoanList{}).Once()
	req := httptest.NewRequest(http.MethodGet, "/?offset=1&limit=2", nil)
	rec := httptest.NewRecorder()
	assert.NoError(t, h.List()(e.NewContext(req, rec)))
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
	assert.NoError(t, h.Get()(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	u.AssertNotCalled(t, "Get")
}
