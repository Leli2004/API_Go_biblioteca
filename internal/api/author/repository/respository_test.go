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

	t.Cleanup(func() {
		_ = db.Close()
	})

	return repositorySetup{
		db:      db,
		sqlMock: sqlMock,
		ctx:     context.Background(),
	}
}

func Test_Get_Repository(t *testing.T) {

	t.Run("Happy Path - Retorna autor", func(t *testing.T) {
		setup := setupRepository(t)

		repository := NewGetRepo()

		now := time.Now()

		expected := entity.Author{
			Id:   1,
			Name: "Machado de Assis",
		}

		rows := sqlmock.NewRows(
			[]string{
				"id",
				"name",
				"created_at",
				"updated_at",
			},
		).AddRow(
			expected.Id,
			expected.Name,
			now,
			now,
		)

		setup.sqlMock.ExpectBegin()

		tx, err := setup.db.Beginx()
		assert.NoError(t, err)

		setup.sqlMock.
			ExpectQuery(regexp.QuoteMeta(getSql)).
			WithArgs(expected.Id).
			WillReturnRows(rows)

		returnedCtx, err, result := repository.Execute(
			setup.ctx,
			tx,
			expected.Id,
		)

		assert.NoError(t, err)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, expected.Id, result.Id)
		assert.Equal(t, expected.Name, result.Name)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		setup := setupRepository(t)

		repository := NewGetRepo()
		expectedError := errors.New("erro ao consultar autor")

		setup.sqlMock.ExpectBegin()

		tx, err := setup.db.Beginx()
		assert.NoError(t, err)

		setup.sqlMock.
			ExpectQuery(regexp.QuoteMeta(getSql)).
			WithArgs(1).
			WillReturnError(expectedError)

		returnedCtx, err, result := repository.Execute(
			setup.ctx,
			tx,
			1,
		)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Author{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Autor não encontrado", func(t *testing.T) {
		setup := setupRepository(t)

		repository := NewGetRepo()

		setup.sqlMock.ExpectBegin()

		tx, err := setup.db.Beginx()
		assert.NoError(t, err)

		setup.sqlMock.
			ExpectQuery(regexp.QuoteMeta(getSql)).
			WithArgs(999).
			WillReturnRows(
				sqlmock.NewRows(
					[]string{
						"id",
						"name",
						"created_at",
						"updated_at",
					},
				),
			)

		returnedCtx, err, result := repository.Execute(
			setup.ctx,
			tx,
			999,
		)

		assert.Error(t, err)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Author{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}

func Test_Create_Repository(t *testing.T) {

	t.Run("Happy Path - Cria autor", func(t *testing.T) {
		setup := setupRepository(t)

		repository := NewCreateRepo()

		input := entity.Author{
			Name: "Conceição Evaristo",
		}

		expected := entity.Author{
			Id:   1,
			Name: input.Name,
		}

		now := time.Now()

		rows := sqlmock.NewRows(
			[]string{
				"id",
				"name",
				"created_at",
				"updated_at",
			},
		).AddRow(
			expected.Id,
			expected.Name,
			now,
			now,
		)

		setup.sqlMock.ExpectBegin()

		tx, err := setup.db.Beginx()
		assert.NoError(t, err)

		setup.sqlMock.
			ExpectQuery(regexp.QuoteMeta(createSql)).
			WithArgs(input.Name).
			WillReturnRows(rows)

		returnedCtx, err, result := repository.Execute(
			setup.ctx,
			tx,
			input,
		)

		assert.NoError(t, err)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, expected.Id, result.Id)
		assert.Equal(t, expected.Name, result.Name)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		setup := setupRepository(t)

		repository := NewCreateRepo()

		input := entity.Author{
			Name: "Conceição Evaristo",
		}

		expectedError := errors.New("erro ao inserir autor")

		setup.sqlMock.ExpectBegin()

		tx, err := setup.db.Beginx()
		assert.NoError(t, err)

		setup.sqlMock.
			ExpectQuery(regexp.QuoteMeta(createSql)).
			WithArgs(input.Name).
			WillReturnError(expectedError)

		returnedCtx, err, result := repository.Execute(
			setup.ctx,
			tx,
			input,
		)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Author{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}

func Test_Update_Repository(t *testing.T) {

	t.Run("Happy Path - Atualiza autor", func(t *testing.T) {
		setup := setupRepository(t)

		repository := NewUpdateRepo()

		id := 1

		input := entity.Author{
			Name: "Machado de Assis Atualizado",
		}

		expected := entity.Author{
			Id:   id,
			Name: input.Name,
		}

		now := time.Now()

		rows := sqlmock.NewRows(
			[]string{
				"id",
				"name",
				"created_at",
				"updated_at",
			},
		).AddRow(
			expected.Id,
			expected.Name,
			now,
			now,
		)

		setup.sqlMock.ExpectBegin()

		tx, err := setup.db.Beginx()
		assert.NoError(t, err)

		setup.sqlMock.
			ExpectQuery(regexp.QuoteMeta(updateSql)).
			WithArgs(input.Name, id).
			WillReturnRows(rows)

		returnedCtx, err, result := repository.Execute(
			setup.ctx,
			tx,
			id,
			input,
		)

		assert.NoError(t, err)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, expected.Id, result.Id)
		assert.Equal(t, expected.Name, result.Name)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		setup := setupRepository(t)

		repository := NewUpdateRepo()

		id := 1

		input := entity.Author{
			Name: "Machado de Assis Atualizado",
		}

		expectedError := errors.New("erro ao atualizar autor")

		setup.sqlMock.ExpectBegin()

		tx, err := setup.db.Beginx()
		assert.NoError(t, err)

		setup.sqlMock.
			ExpectQuery(regexp.QuoteMeta(updateSql)).
			WithArgs(input.Name, id).
			WillReturnError(expectedError)

		returnedCtx, err, result := repository.Execute(
			setup.ctx,
			tx,
			id,
			input,
		)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Author{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}

func Test_Delete_Repository(t *testing.T) {

	t.Run("Happy Path - Exclui autor", func(t *testing.T) {
		setup := setupRepository(t)

		repository := NewDeleteRepo()

		id := 1
		now := time.Now()

		rows := sqlmock.NewRows(
			[]string{
				"id",
				"name",
				"created_at",
				"updated_at",
			},
		).AddRow(
			id,
			"Machado de Assis",
			now,
			now,
		)

		setup.sqlMock.ExpectBegin()

		tx, err := setup.db.Beginx()
		assert.NoError(t, err)

		setup.sqlMock.
			ExpectQuery(regexp.QuoteMeta(deleteSql)).
			WithArgs(id).
			WillReturnRows(rows)

		returnedCtx, err, result := repository.Execute(
			setup.ctx,
			tx,
			id,
		)

		assert.NoError(t, err)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, id, result.Id)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		setup := setupRepository(t)

		repository := NewDeleteRepo()

		id := 1
		expectedError := errors.New("erro ao excluir autor")

		setup.sqlMock.ExpectBegin()

		tx, err := setup.db.Beginx()
		assert.NoError(t, err)

		setup.sqlMock.
			ExpectQuery(regexp.QuoteMeta(deleteSql)).
			WithArgs(id).
			WillReturnError(expectedError)

		returnedCtx, err, result := repository.Execute(
			setup.ctx,
			tx,
			id,
		)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Author{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}

func Test_List_Repository(t *testing.T) {
	
	t.Run("Happy Path - Retorna lista de autores", func(t *testing.T) {
		setup := setupRepository(t)

		repository := NewListRepo()

		input := entity.AuthorFilters{
			Offset: 0,
			Limit:  10,
		}

		now := time.Now()

		rows := sqlmock.NewRows(
			[]string{
				"id",
				"name",
				"created_at",
				"updated_at",
			},
		).
			AddRow(
				2,
				"Clarice Lispector",
				now,
				now,
			).
			AddRow(
				1,
				"Machado de Assis",
				now,
				now,
			)

		setup.sqlMock.ExpectBegin()

		tx, err := setup.db.Beginx()
		assert.NoError(t, err)

		setup.sqlMock.
			ExpectQuery(regexp.QuoteMeta(listSql)).
			WithArgs(input.Offset, input.Limit).
			WillReturnRows(rows)

		returnedCtx, err, result := repository.Execute(
			setup.ctx,
			tx,
			input,
		)

		assert.NoError(t, err)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, input.Offset, result.Offset)
		assert.Equal(t, 2, result.Limit)
		assert.Len(t, result.Data, 2)
		assert.Equal(t, 2, result.Data[0].Id)
		assert.Equal(t, "Clarice Lispector", result.Data[0].Name)
		assert.Equal(t, 1, result.Data[1].Id)
		assert.Equal(t, "Machado de Assis", result.Data[1].Name)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Happy Path - Retorna lista vazia", func(t *testing.T) {
		setup := setupRepository(t)

		repository := NewListRepo()

		input := entity.AuthorFilters{
			Offset: 0,
			Limit:  10,
		}

		rows := sqlmock.NewRows(
			[]string{
				"id",
				"name",
				"created_at",
				"updated_at",
			},
		)

		setup.sqlMock.ExpectBegin()

		tx, err := setup.db.Beginx()
		assert.NoError(t, err)

		setup.sqlMock.
			ExpectQuery(regexp.QuoteMeta(listSql)).
			WithArgs(input.Offset, input.Limit).
			WillReturnRows(rows)

		returnedCtx, err, result := repository.Execute(
			setup.ctx,
			tx,
			input,
		)

		assert.NoError(t, err)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, input.Offset, result.Offset)
		assert.Equal(t, 0, result.Limit)
		assert.Empty(t, result.Data)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		setup := setupRepository(t)

		repository := NewListRepo()

		input := entity.AuthorFilters{
			Offset: 0,
			Limit:  10,
		}

		expectedError := errors.New("erro ao listar autores")

		setup.sqlMock.ExpectBegin()

		tx, err := setup.db.Beginx()
		assert.NoError(t, err)

		setup.sqlMock.
			ExpectQuery(regexp.QuoteMeta(listSql)).
			WithArgs(input.Offset, input.Limit).
			WillReturnError(expectedError)

		returnedCtx, err, result := repository.Execute(
			setup.ctx,
			tx,
			input,
		)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.AuthorList{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}
