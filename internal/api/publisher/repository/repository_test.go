package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type repositorySetup struct {
	db      *sqlx.DB
	sqlMock sqlmock.Sqlmock
	ctx     context.Context
}

func setupRepository(t *testing.T) repositorySetup {
	sqlDB, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	db := sqlx.NewDb(sqlDB, "sqlmock")
	t.Cleanup(func() { _ = db.Close() })
	return repositorySetup{db: db, sqlMock: sqlMock, ctx: context.Background()}
}

func strPtr(v string) *string { return &v }
func boolPtr(v bool) *bool    { return &v }

func Test_Get_Repository(t *testing.T) {

	t.Run("Happy Path - Retorna publisher", func(t *testing.T) {
		setup := setupRepository(t)
		repository := NewRepository()
		now := time.Now().UTC().Format(time.RFC3339)
		expected := entity.Publisher{Name: "Companhia das Letras", Website: strPtr("https://example.com")}
		expected.Id = 1
		rows := sqlmock.NewRows([]string{"id", "name", "website", "created_at", "updated_at"}).AddRow(expected.Id,
			expected.Name,
			expected.Website,
			now,
			now)
		setup.sqlMock.ExpectBegin()
		tx, err := setup.db.Beginx()
		assert.NoError(t, err)
		setup.sqlMock.ExpectQuery(regexp.QuoteMeta(getSql)).WithArgs(expected.Id).WillReturnRows(rows)
		returnedCtx, err, result := repository.Get(setup.ctx, tx, expected.Id)
		assert.NoError(t, err)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, expected.Id, result.Id)
		assert.Equal(t, expected.Name, result.Name)
		assert.Equal(t, expected.Website, result.Website)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		setup := setupRepository(t)
		repository := NewRepository()
		expectedError := errors.New("erro ao consultar")
		setup.sqlMock.ExpectBegin()
		tx, err := setup.db.Beginx()
		assert.NoError(t, err)
		setup.sqlMock.ExpectQuery(regexp.QuoteMeta(getSql)).WithArgs(1).WillReturnError(expectedError)
		returnedCtx, err, result := repository.Get(setup.ctx, tx, 1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Publisher{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}

func Test_Create_Repository(t *testing.T) {

	t.Run("Happy Path - Cria publisher", func(t *testing.T) {
		setup := setupRepository(t)
		repository := NewRepository()
		input := entity.Publisher{Name: "Companhia das Letras", Website: strPtr("https://example.com")}
		expected := input
		expected.Id = 1
		now := time.Now().UTC().Format(time.RFC3339)
		rows := sqlmock.NewRows([]string{"id", "name", "website", "created_at", "updated_at"}).AddRow(expected.Id,
			expected.Name,
			expected.Website,
			now,
			now)
		setup.sqlMock.ExpectBegin()
		tx, err := setup.db.Beginx()
		assert.NoError(t, err)
		setup.sqlMock.ExpectQuery(regexp.QuoteMeta(createSql)).WithArgs(input.Name, input.Website).WillReturnRows(rows)
		returnedCtx, err, result := repository.Create(setup.ctx, tx, input)
		assert.NoError(t, err)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, expected.Id, result.Id)
		assert.Equal(t, expected.Name, result.Name)
		assert.Equal(t, expected.Website, result.Website)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		setup := setupRepository(t)
		repository := NewRepository()
		input := entity.Publisher{Name: "Companhia das Letras", Website: strPtr("https://example.com")}
		expectedError := errors.New("erro ao inserir")
		setup.sqlMock.ExpectBegin()
		tx, err := setup.db.Beginx()
		assert.NoError(t, err)
		setup.sqlMock.ExpectQuery(regexp.QuoteMeta(createSql)).WithArgs(input.Name, input.Website).WillReturnError(expectedError)
		returnedCtx, err, result := repository.Create(setup.ctx, tx, input)
		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Publisher{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}

func Test_Update_Repository(t *testing.T) {

	t.Run("Happy Path - Atualiza publisher", func(t *testing.T) {
		setup := setupRepository(t)
		repository := NewRepository()
		id := 1
		input := entity.Publisher{Name: "Companhia das Letras", Website: strPtr("https://example.com")}
		expected := input
		expected.Id = id
		now := time.Now().UTC().Format(time.RFC3339)
		rows := sqlmock.NewRows([]string{"id", "name", "website", "created_at", "updated_at"}).AddRow(expected.Id,
			expected.Name,
			expected.Website,
			now,
			now)
		setup.sqlMock.ExpectBegin()
		tx, err := setup.db.Beginx()
		assert.NoError(t, err)
		setup.sqlMock.ExpectQuery(regexp.QuoteMeta(updateSql)).WithArgs(input.Name, input.Website, id).WillReturnRows(rows)
		returnedCtx, err, result := repository.Update(setup.ctx, tx, id, input)
		assert.NoError(t, err)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, expected.Id, result.Id)
		assert.Equal(t, expected.Name, result.Name)
		assert.Equal(t, expected.Website, result.Website)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		setup := setupRepository(t)
		repository := NewRepository()
		id := 1
		input := entity.Publisher{Name: "Companhia das Letras", Website: strPtr("https://example.com")}
		expectedError := errors.New("erro ao atualizar")
		setup.sqlMock.ExpectBegin()
		tx, err := setup.db.Beginx()
		assert.NoError(t, err)
		setup.sqlMock.ExpectQuery(regexp.QuoteMeta(updateSql)).WithArgs(input.Name, input.Website, id).WillReturnError(expectedError)
		returnedCtx, err, result := repository.Update(setup.ctx, tx, id, input)
		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Publisher{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}

func Test_Delete_Repository(t *testing.T) {

	t.Run("Happy Path - Exclui publisher", func(t *testing.T) {
		setup := setupRepository(t)
		repository := NewRepository()
		id := 1
		expected := entity.Publisher{Name: "Companhia das Letras", Website: strPtr("https://example.com")}
		expected.Id = id
		now := time.Now().UTC().Format(time.RFC3339)
		rows := sqlmock.NewRows([]string{"id", "name", "website", "created_at", "updated_at"}).AddRow(expected.Id,
			expected.Name,
			expected.Website,
			now,
			now)
		setup.sqlMock.ExpectBegin()
		tx, err := setup.db.Beginx()
		assert.NoError(t, err)
		setup.sqlMock.ExpectQuery(regexp.QuoteMeta(deleteSql)).WithArgs(id).WillReturnRows(rows)
		returnedCtx, err := repository.Delete(setup.ctx, tx, id)
		assert.NoError(t, err)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		setup := setupRepository(t)
		repository := NewRepository()
		id := 1
		expectedError := errors.New("erro ao excluir")
		setup.sqlMock.ExpectBegin()
		tx, err := setup.db.Beginx()
		assert.NoError(t, err)
		setup.sqlMock.ExpectQuery(regexp.QuoteMeta(deleteSql)).WithArgs(id).WillReturnError(expectedError)
		returnedCtx, err := repository.Delete(setup.ctx, tx, id)
		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}

func Test_List_Repository(t *testing.T) {

	t.Run("Happy Path - Retorna lista", func(t *testing.T) {
		setup := setupRepository(t)
		repository := NewRepository()
		input := entity.PublisherFilters{Offset: 0, Limit: 10}
		now := time.Now().UTC().Format(time.RFC3339)
		expected := entity.Publisher{Name: "Companhia das Letras", Website: strPtr("https://example.com")}
		expected.Id = 1
		rows := sqlmock.NewRows([]string{"id", "name", "website", "created_at", "updated_at"}).AddRow(expected.Id,
			expected.Name,
			expected.Website,
			now,
			now)
		setup.sqlMock.ExpectBegin()
		tx, err := setup.db.Beginx()
		assert.NoError(t, err)
		setup.sqlMock.ExpectQuery(regexp.QuoteMeta(listSql)).WithArgs(input.Offset, input.Limit).WillReturnRows(rows)
		returnedCtx, err, result := repository.List(setup.ctx, tx, input)
		assert.NoError(t, err)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, input.Offset, result.Offset)
		assert.Equal(t, 1, result.Limit)
		assert.Len(t, result.Data, 1)
		assert.Equal(t, expected.Id, result.Data[0].Id)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Happy Path - Retorna lista vazia", func(t *testing.T) {
		setup := setupRepository(t)
		repository := NewRepository()
		input := entity.PublisherFilters{Offset: 0, Limit: 10}
		rows := sqlmock.NewRows([]string{"id", "name", "website", "created_at", "updated_at"})
		setup.sqlMock.ExpectBegin()
		tx, err := setup.db.Beginx()
		assert.NoError(t, err)
		setup.sqlMock.ExpectQuery(regexp.QuoteMeta(listSql)).WithArgs(input.Offset, input.Limit).WillReturnRows(rows)
		returnedCtx, err, result := repository.List(setup.ctx, tx, input)
		assert.NoError(t, err)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, 0, result.Limit)
		assert.Empty(t, result.Data)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		setup := setupRepository(t)
		repository := NewRepository()
		input := entity.PublisherFilters{Offset: 0, Limit: 10}
		expectedError := errors.New("erro ao listar")
		setup.sqlMock.ExpectBegin()
		tx, err := setup.db.Beginx()
		assert.NoError(t, err)
		setup.sqlMock.ExpectQuery(regexp.QuoteMeta(listSql)).WithArgs(input.Offset, input.Limit).WillReturnError(expectedError)
		returnedCtx, err, result := repository.List(setup.ctx, tx, input)
		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.PublisherList{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}
