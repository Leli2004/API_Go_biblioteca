package repository

import (
	"context"
	"database/sql"
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
	return repositorySetup{db, sqlMock, context.Background()}
}

func Test_Get_Repository(t *testing.T) {
	
	t.Run("Happy Path - Retorna exemplar", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		now := time.Now().UTC().Format(time.RFC3339)
		expected := entity.BookCopy{Id: 1, BookId: 2, Barcode: "ABC", Status: "available"}
		rows := sqlmock.NewRows([]string{"id", "book_id", "barcode", "status", "created_at", "updated_at"}).AddRow(expected.Id, expected.BookId, expected.Barcode, expected.Status, now, now)
		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(getSql)).WithArgs(1).WillReturnRows(rows)
		ctx, err, result := r.Get(s.ctx, tx, 1)
		assert.NoError(t, err)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, expected.Id, result.Id)
		assert.Equal(t, expected.Barcode, result.Barcode)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		expectedError := errors.New("erro")
		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(getSql)).WithArgs(1).WillReturnError(expectedError)
		ctx, err, result := r.Get(s.ctx, tx, 1)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, entity.BookCopy{}, result)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})
}

func Test_Create_Repository(t *testing.T) {

	t.Run("Happy Path - Cria exemplar", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		input := entity.BookCopy{BookId: 2, Barcode: "ABC", Status: "available"}
		now := time.Now().UTC().Format(time.RFC3339)
		rows := sqlmock.NewRows([]string{"id", "book_id", "barcode", "status", "created_at", "updated_at"}).AddRow(1, input.BookId, input.Barcode, input.Status, now, now)
		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(checkBarcodeSql)).WithArgs(input.Barcode).WillReturnError(sql.ErrNoRows)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(createSql)).WithArgs(input.BookId, input.Barcode, input.Status).WillReturnRows(rows)
		ctx, err, result := r.Create(s.ctx, tx, input)
		assert.NoError(t, err)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, 1, result.Id)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})
	t.Run("Error Path - Barcode já existe", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		input := entity.BookCopy{BookId: 2, Barcode: "ABC", Status: "available"}
		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(checkBarcodeSql)).WithArgs(input.Barcode).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(9))
		ctx, err, result := r.Create(s.ctx, tx, input)
		assert.Error(t, err)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, entity.BookCopy{}, result)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})
	t.Run("Error Path - Banco retorna erro ao criar", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		input := entity.BookCopy{BookId: 2, Barcode: "ABC", Status: "available"}
		expectedError := errors.New("erro")
		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(checkBarcodeSql)).WithArgs(input.Barcode).WillReturnError(sql.ErrNoRows)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(createSql)).WithArgs(input.BookId, input.Barcode, input.Status).WillReturnError(expectedError)
		ctx, err, result := r.Create(s.ctx, tx, input)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, entity.BookCopy{}, result)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})
}

func Test_Update_Repository(t *testing.T) {
	s := setupRepository(t)
	r := NewRepository()
	input := entity.BookCopy{BookId: 2, Barcode: "ABC2", Status: "maintenance"}
	now := time.Now().UTC().Format(time.RFC3339)
	rows := sqlmock.NewRows([]string{"id", "book_id", "barcode", "status", "created_at", "updated_at"}).AddRow(1, input.BookId, input.Barcode, input.Status, now, now)
	s.sqlMock.ExpectBegin()
	tx, err := s.db.Beginx()
	assert.NoError(t, err)
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(updateSql)).WithArgs(input.BookId, input.Barcode, input.Status, 1).WillReturnRows(rows)
	ctx, err, result := r.Update(s.ctx, tx, 1, input)
	assert.NoError(t, err)
	assert.Equal(t, s.ctx, ctx)
	assert.Equal(t, input.Status, result.Status)
	assert.NoError(t, s.sqlMock.ExpectationsWereMet())
}

func Test_Delete_Repository(t *testing.T) {
	s := setupRepository(t)
	r := NewRepository()
	s.sqlMock.ExpectBegin()
	tx, err := s.db.Beginx()
	assert.NoError(t, err)
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(deleteSql)).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	ctx, err := r.Delete(s.ctx, tx, 1)
	assert.NoError(t, err)
	assert.Equal(t, s.ctx, ctx)
	assert.NoError(t, s.sqlMock.ExpectationsWereMet())
}

func Test_List_Repository(t *testing.T) {
	s := setupRepository(t)
	r := NewRepository()
	input := entity.BookCopyFilters{Offset: 0, Limit: 10}
	now := time.Now().UTC().Format(time.RFC3339)
	rows := sqlmock.NewRows([]string{"id", "book_id", "barcode", "status", "created_at", "updated_at"}).AddRow(1, 2, "ABC", "available", now, now)
	s.sqlMock.ExpectBegin()
	tx, err := s.db.Beginx()
	assert.NoError(t, err)
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(listSql)).WithArgs(input.Offset, input.Limit).WillReturnRows(rows)
	ctx, err, result := r.List(s.ctx, tx, input)
	assert.NoError(t, err)
	assert.Equal(t, s.ctx, ctx)
	assert.Len(t, result.Data, 1)
	assert.Equal(t, 1, result.Limit)
	assert.NoError(t, s.sqlMock.ExpectationsWereMet())
}