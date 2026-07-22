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

func intPtr(v int) *int       { return &v }
func strPtr(v string) *string { return &v }
func bookRows(b entity.Book, now string) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "publisher_id", "title", "publication_year", "description", "created_at", "updated_at"}).AddRow(b.Id, b.PublisherId, b.Title, b.PublicationYear, b.Description, now, now)
}

func Test_Get_Repository(t *testing.T) {
	t.Run("Happy Path - Retorna livro com sublistas", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		now := time.Now().UTC().Format(time.RFC3339)
		expected := entity.Book{Id: 1, PublisherId: intPtr(2), Title: "Dom Casmurro", PublicationYear: intPtr(1899), Description: strPtr("Romance")}
		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(getSql)).WithArgs(1).WillReturnRows(bookRows(expected, now))
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(authorsSql)).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).AddRow(1, "Machado de Assis", now, now))
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(genresSql)).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).AddRow(1, "Romance", nil, now, now))
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(copiesSql)).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "book_id", "barcode", "status", "created_at", "updated_at"}).AddRow(1, 1, "ABC", "available", now, now))
		ctx, err, result := r.Get(s.ctx, tx, 1)
		assert.NoError(t, err)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, expected.Id, result.Id)
		assert.Len(t, result.Authors, 1)
		assert.Len(t, result.Genres, 1)
		assert.Len(t, result.Copies, 1)
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
		assert.Equal(t, entity.Book{}, result)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})
}

func Test_Create_Repository(t *testing.T) {
	s := setupRepository(t)
	r := NewRepository()
	input := entity.Book{PublisherId: intPtr(2), Title: "Livro", PublicationYear: intPtr(2026), Description: strPtr("Descrição"), AuthorIds: []int{1, 2}, GenreIds: []int{3}}
	expected := input
	expected.Id = 1
	now := time.Now().UTC().Format(time.RFC3339)
	s.sqlMock.ExpectBegin()
	tx, err := s.db.Beginx()
	assert.NoError(t, err)
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(createSql)).WithArgs(input.PublisherId, input.Title, input.PublicationYear, input.Description).WillReturnRows(bookRows(expected, now))
	s.sqlMock.ExpectExec(regexp.QuoteMeta(insertAuthorSql)).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	s.sqlMock.ExpectExec(regexp.QuoteMeta(insertAuthorSql)).WithArgs(1, 2).WillReturnResult(sqlmock.NewResult(1, 1))
	s.sqlMock.ExpectExec(regexp.QuoteMeta(insertGenreSql)).WithArgs(1, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	ctx, err, result := r.Create(s.ctx, tx, input)
	assert.NoError(t, err)
	assert.Equal(t, s.ctx, ctx)
	assert.Equal(t, 1, result.Id)
	assert.NoError(t, s.sqlMock.ExpectationsWereMet())
}

func Test_Update_Repository(t *testing.T) {
	s := setupRepository(t)
	r := NewRepository()
	input := entity.Book{PublisherId: intPtr(2), Title: "Atualizado", PublicationYear: intPtr(2026), Description: strPtr("Descrição"), AuthorIds: []int{1}, GenreIds: []int{2}}
	expected := input
	expected.Id = 1
	now := time.Now().UTC().Format(time.RFC3339)
	s.sqlMock.ExpectBegin()
	tx, err := s.db.Beginx()
	assert.NoError(t, err)
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(updateSql)).WithArgs(input.PublisherId, input.Title, input.PublicationYear, input.Description, 1).WillReturnRows(bookRows(expected, now))
	s.sqlMock.ExpectExec(regexp.QuoteMeta(deleteAuthorsSql)).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
	s.sqlMock.ExpectExec(regexp.QuoteMeta(deleteGenresSql)).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
	s.sqlMock.ExpectExec(regexp.QuoteMeta(insertAuthorSql)).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	s.sqlMock.ExpectExec(regexp.QuoteMeta(insertGenreSql)).WithArgs(1, 2).WillReturnResult(sqlmock.NewResult(1, 1))
	ctx, err, result := r.Update(s.ctx, tx, 1, input)
	assert.NoError(t, err)
	assert.Equal(t, s.ctx, ctx)
	assert.Equal(t, input.Title, result.Title)
	assert.NoError(t, s.sqlMock.ExpectationsWereMet())
}

func Test_Delete_Repository(t *testing.T) {
	s := setupRepository(t)
	r := NewRepository()
	expected := entity.Book{Id: 1, Title: "Livro"}
	now := time.Now().UTC().Format(time.RFC3339)
	s.sqlMock.ExpectBegin()
	tx, err := s.db.Beginx()
	assert.NoError(t, err)
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(deleteSql)).WithArgs(1).WillReturnRows(bookRows(expected, now))
	ctx, err := r.Delete(s.ctx, tx, 1)
	assert.NoError(t, err)
	assert.Equal(t, s.ctx, ctx)
	assert.NoError(t, s.sqlMock.ExpectationsWereMet())
}

func Test_List_Repository(t *testing.T) {
	s := setupRepository(t)
	r := NewRepository()
	input := entity.BookFilters{Offset: 0, Limit: 10}
	expected := entity.Book{Id: 1, Title: "Livro"}
	now := time.Now().UTC().Format(time.RFC3339)
	s.sqlMock.ExpectBegin()
	tx, err := s.db.Beginx()
	assert.NoError(t, err)
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(listSql)).WithArgs(input.Offset, input.Limit).WillReturnRows(bookRows(expected, now))
	ctx, err, result := r.List(s.ctx, tx, input)
	assert.NoError(t, err)
	assert.Equal(t, s.ctx, ctx)
	assert.Equal(t, 1, result.Limit)
	assert.Len(t, result.Data, 1)
	assert.NoError(t, s.sqlMock.ExpectationsWereMet())
}
