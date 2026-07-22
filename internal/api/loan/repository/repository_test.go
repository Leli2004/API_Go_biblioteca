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
	return repositorySetup{db, sqlMock, context.Background()}
}

func loanRows(l entity.Loan, now string) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "user_id", "book_copy_id", "loan_date", "due_date", "returned_at", "status", "created_at", "updated_at"}).AddRow(l.Id, l.UserId, l.BookCopyId, l.LoanDate, l.DueDate, l.ReturnedAt, l.Status, now, now)
}

func Test_Get_Repository(t *testing.T) {

	t.Run("Happy Path - Retorna empréstimo", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		due := "2026-08-01T00:00:00Z"
		loanDate := "2026-07-21T00:00:00Z"
		expected := entity.Loan{Id: 1, UserId: 2, BookCopyId: 3, LoanDate: &loanDate, DueDate: &due, Status: "active"}
		now := time.Now().UTC().Format(time.RFC3339)
		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		q := `SELECT id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at FROM biblioteca.loans WHERE id = $1`
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1).WillReturnRows(loanRows(expected, now))
		ctx, err, result := r.Get(s.ctx, tx, 1)
		assert.NoError(t, err)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, expected.Id, result.Id)
		assert.Equal(t, expected.Status, result.Status)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		expectedError := errors.New("erro")
		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		q := `SELECT id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at FROM biblioteca.loans WHERE id = $1`
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1).WillReturnError(expectedError)
		ctx, err, result := r.Get(s.ctx, tx, 1)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, entity.Loan{}, result)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})
}

func Test_Create_Repository(t *testing.T) {
	s := setupRepository(t)
	r := NewRepository()
	loanDate := "2026-07-21T00:00:00Z"
	due := "2026-08-01T00:00:00Z"
	input := entity.Loan{UserId: 2, BookCopyId: 3, LoanDate: &loanDate, DueDate: &due, Status: "active"}
	expected := input
	expected.Id = 1
	now := time.Now().UTC().Format(time.RFC3339)
	s.sqlMock.ExpectBegin()
	tx, err := s.db.Beginx()
	assert.NoError(t, err)
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(createAndUpdateSql)).WithArgs(input.UserId, input.BookCopyId, input.LoanDate, input.DueDate, input.Status).WillReturnRows(loanRows(expected, now))
	ctx, err, result := r.Create(s.ctx, tx, input)
	assert.NoError(t, err)
	assert.Equal(t, s.ctx, ctx)
	assert.Equal(t, 1, result.Id)
	assert.NoError(t, s.sqlMock.ExpectationsWereMet())
}

func Test_Update_Repository(t *testing.T) {
	s := setupRepository(t)
	r := NewRepository()
	returned := "2026-07-25T00:00:00Z"
	input := entity.Loan{ReturnedAt: &returned, Status: "returned"}
	expected := entity.Loan{Id: 1, UserId: 2, BookCopyId: 3, ReturnedAt: &returned, Status: "returned"}
	now := time.Now().UTC().Format(time.RFC3339)
	s.sqlMock.ExpectBegin()
	tx, err := s.db.Beginx()
	assert.NoError(t, err)
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(updateAndMaybeRestoreSql)).WithArgs(input.ReturnedAt, input.Status, 1).WillReturnRows(loanRows(expected, now))
	ctx, err, result := r.Update(s.ctx, tx, 1, input)
	assert.NoError(t, err)
	assert.Equal(t, s.ctx, ctx)
	assert.Equal(t, "returned", result.Status)
	assert.NoError(t, s.sqlMock.ExpectationsWereMet())
}

func Test_GetActiveByUserAndBookCopy_Repository(t *testing.T) {
	s := setupRepository(t)
	r := NewRepository()
	expected := entity.Loan{Id: 1, UserId: 2, BookCopyId: 3, Status: "active"}
	now := time.Now().UTC().Format(time.RFC3339)
	q := `SELECT id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at FROM biblioteca.loans WHERE user_id = $1 AND book_copy_id = $2 AND status = 'active' LIMIT 1`
	s.sqlMock.ExpectBegin()
	tx, err := s.db.Beginx()
	assert.NoError(t, err)
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(2, 3).WillReturnRows(loanRows(expected, now))
	ctx, err, result := r.GetActiveByUserAndBookCopy(s.ctx, tx, 2, 3)
	assert.NoError(t, err)
	assert.Equal(t, s.ctx, ctx)
	assert.Equal(t, expected.Id, result.Id)
	assert.NoError(t, s.sqlMock.ExpectationsWereMet())
}

func Test_List_Repository(t *testing.T) {
	s := setupRepository(t)
	r := NewRepository()
	input := entity.LoanFilters{Offset: 0, Limit: 10}
	expected := entity.Loan{Id: 1, UserId: 2, BookCopyId: 3, Status: "active"}
	now := time.Now().UTC().Format(time.RFC3339)
	s.sqlMock.ExpectBegin()
	tx, err := s.db.Beginx()
	assert.NoError(t, err)
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(listSql)).WithArgs(input.Limit, input.Offset).WillReturnRows(loanRows(expected, now))
	ctx, err, result := r.List(s.ctx, tx, input)
	assert.NoError(t, err)
	assert.Equal(t, s.ctx, ctx)
	assert.Equal(t, input.Limit, result.Limit)
	assert.Len(t, result.Data, 1)
	assert.NoError(t, s.sqlMock.ExpectationsWereMet())
}

func Test_Delete_Repository(t *testing.T) {
	s := setupRepository(t)
	r := NewRepository()
	expected := entity.Loan{Id: 1, UserId: 2, BookCopyId: 3, Status: "active"}
	now := time.Now().UTC().Format(time.RFC3339)
	s.sqlMock.ExpectBegin()
	tx, err := s.db.Beginx()
	assert.NoError(t, err)
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(deleteSql)).WithArgs(1).WillReturnRows(loanRows(expected, now))
	ctx, err, result := r.Delete(s.ctx, tx, 1)
	assert.NoError(t, err)
	assert.Equal(t, s.ctx, ctx)
	assert.Equal(t, 1, result.Id)
	assert.NoError(t, s.sqlMock.ExpectationsWereMet())
}
