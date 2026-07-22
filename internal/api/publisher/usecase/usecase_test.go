package usecase

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	mm "github.com/Leli2004/API_Go_biblioteca/internal/api/publisher/mocks"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type useCaseSetup struct {
	uc      *PublisherUC
	repo    *mm.Repository
	sqlMock sqlmock.Sqlmock
	db      *sqlx.DB
	ctx     context.Context
}

func setup(t *testing.T) useCaseSetup {
	s, m, e := sqlmock.New()
	assert.NoError(t, e)
	db := sqlx.NewDb(s, "sqlmock")
	r := mm.NewRepository(t)
	t.Cleanup(func() { _ = db.Close() })
	return useCaseSetup{NewUseCase(db, r), r, m, db, context.Background()}
}

func Test_Get_UseCase(t *testing.T) {

	t.Run("Happy Path - Retorna registro", func(t *testing.T) {
		s := setup(t)
		expected := entity.Publisher{}
		s.sqlMock.ExpectBegin()
		s.repo.On("Get", mock.Anything, mock.Anything, 1).Return(s.ctx, nil, expected).Once()
		s.sqlMock.ExpectCommit()
		c, e, r := s.uc.Get(s.ctx, 1)
		assert.NoError(t, e)
		assert.Equal(t, s.ctx, c)
		assert.Equal(t, expected, r)
	})

	t.Run("Error Path - Repository retorna erro", func(t *testing.T) {
		s := setup(t)
		x := errors.New("repository error")
		s.sqlMock.ExpectBegin()
		s.repo.On("Get", mock.Anything, mock.Anything, 1).Return(s.ctx, x, entity.Publisher{}).Once()
		s.sqlMock.ExpectRollback()
		_, e, _ := s.uc.Get(s.ctx, 1)
		assert.ErrorIs(t, e, x)
	})
}

func Test_List_UseCase(t *testing.T) {
	s := setup(t)
	in := entity.PublisherFilters{}
	expectedIn := in
	expectedIn.SetDefault()
	expected := entity.PublisherList{}
	s.sqlMock.ExpectBegin()
	s.repo.On("List", mock.Anything, mock.Anything, expectedIn).Return(s.ctx, nil, expected).Once()
	s.sqlMock.ExpectCommit()
	_, e, r := s.uc.List(s.ctx, in)
	assert.NoError(t, e)
	assert.Equal(t, expected, r)
}

func Test_Create_UseCase_Validation(t *testing.T) {
	s := setup(t)
	s.sqlMock.ExpectBegin()
	_, e, r := s.uc.Create(s.ctx, entity.Publisher{})
	assert.Error(t, e)
	assert.Equal(t, entity.Publisher{}, r)
	s.repo.AssertNotCalled(t, "Create")
	s.sqlMock.ExpectRollback()
}

func Test_Update_UseCase_Validation(t *testing.T) {
	s := setup(t)
	s.sqlMock.ExpectBegin()
	_, e, r := s.uc.Update(s.ctx, 1, entity.Publisher{})
	assert.Error(t, e)
	assert.Equal(t, entity.Publisher{}, r)
	s.repo.AssertNotCalled(t, "Update")
	s.sqlMock.ExpectRollback()
}

func Test_Delete_UseCase(t *testing.T) {
	s := setup(t)
	s.sqlMock.ExpectBegin()
	s.repo.On("Delete", mock.Anything, mock.Anything, 1).Return(s.ctx, nil).Once()
	s.sqlMock.ExpectCommit()
	c, e := s.uc.Delete(s.ctx, 1)
	assert.NoError(t, e)
	assert.Equal(t, s.ctx, c)
}
