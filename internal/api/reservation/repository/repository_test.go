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

func reservationRows(r entity.Reservation, now string) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "user_id", "book_id", "position", "status", "expires_at", "created_at", "updated_at"}).
		AddRow(r.Id, r.UserId, r.BookId, r.Position, r.Status, r.ExpiresAt, now, now)
}

func Test_GetActiveByUserAndBook_Repository(t *testing.T) {

	t.Run("Happy Path - Retorna reserva ativa", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		position := 1
		expected := entity.Reservation{Id: 1, UserId: 2, BookId: 3, Position: &position, Status: entity.ReservationStatusWaiting}
		now := time.Now().UTC().Format(time.RFC3339)
		query := `SELECT id,user_id,book_id,position,status,expires_at,created_at,updated_at FROM biblioteca.reservations WHERE user_id=$1 AND book_id=$2 AND status IN ('waiting','available') ORDER BY id DESC LIMIT 1`

		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(2, 3).WillReturnRows(reservationRows(expected, now))

		ctx, err, result := r.GetActiveByUserAndBook(s.ctx, tx, 2, 3)

		assert.NoError(t, err)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, expected.Id, result.Id)
		assert.Equal(t, expected.Status, result.Status)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		expectedError := errors.New("database error")
		query := `SELECT id,user_id,book_id,position,status,expires_at,created_at,updated_at FROM biblioteca.reservations WHERE user_id=$1 AND book_id=$2 AND status IN ('waiting','available') ORDER BY id DESC LIMIT 1`

		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(2, 3).WillReturnError(expectedError)

		ctx, err, result := r.GetActiveByUserAndBook(s.ctx, tx, 2, 3)

		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, entity.Reservation{}, result)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})
}

func Test_GetNextPosition_Repository(t *testing.T) {

	t.Run("Happy Path - Bloqueia livro e calcula próxima posição", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		lockQuery := `SELECT pg_advisory_xact_lock($1)`
		positionQuery := `SELECT COALESCE(MAX(position),0)+1 FROM biblioteca.reservations WHERE book_id=$1 AND status='waiting'`

		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		s.sqlMock.ExpectExec(regexp.QuoteMeta(lockQuery)).WithArgs(3).WillReturnResult(sqlmock.NewResult(0, 1))
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(positionQuery)).WithArgs(3).
			WillReturnRows(sqlmock.NewRows([]string{"position"}).AddRow(4))

		ctx, err, result := r.GetNextPosition(s.ctx, tx, 3)

		assert.NoError(t, err)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, 4, result)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Falha no lock", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		expectedError := errors.New("lock error")
		lockQuery := `SELECT pg_advisory_xact_lock($1)`

		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		s.sqlMock.ExpectExec(regexp.QuoteMeta(lockQuery)).WithArgs(3).WillReturnError(expectedError)

		ctx, err, result := r.GetNextPosition(s.ctx, tx, 3)

		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, s.ctx, ctx)
		assert.Zero(t, result)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Falha ao calcular posição", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		expectedError := errors.New("position error")
		lockQuery := `SELECT pg_advisory_xact_lock($1)`
		positionQuery := `SELECT COALESCE(MAX(position),0)+1 FROM biblioteca.reservations WHERE book_id=$1 AND status='waiting'`

		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		s.sqlMock.ExpectExec(regexp.QuoteMeta(lockQuery)).WithArgs(3).WillReturnResult(sqlmock.NewResult(0, 1))
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(positionQuery)).WithArgs(3).WillReturnError(expectedError)

		ctx, err, result := r.GetNextPosition(s.ctx, tx, 3)

		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, s.ctx, ctx)
		assert.Zero(t, result)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})
}

func Test_Create_Repository(t *testing.T) {

	t.Run("Happy Path - Cria reserva", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		position := 2
		expiresAt := "2026-08-01T12:00:00Z"
		input := entity.Reservation{UserId: 1, BookId: 3, Position: &position, Status: entity.ReservationStatusWaiting, ExpiresAt: &expiresAt}
		expected := input
		expected.Id = 5
		now := time.Now().UTC().Format(time.RFC3339)
		query := `INSERT INTO biblioteca.reservations(user_id,book_id,position,status,expires_at) VALUES($1,$2,$3,$4,$5::timestamptz) RETURNING id,user_id,book_id,position,status,expires_at,created_at,updated_at`

		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(input.UserId, input.BookId, input.Position, input.Status, input.ExpiresAt).
			WillReturnRows(reservationRows(expected, now))

		ctx, err, result := r.Create(s.ctx, tx, input)

		assert.NoError(t, err)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, expected.Id, result.Id)
		assert.Equal(t, expected.Status, result.Status)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Banco retorna erro", func(t *testing.T) {
		s := setupRepository(t)
		r := NewRepository()
		expectedError := errors.New("insert error")
		input := entity.Reservation{UserId: 1, BookId: 3, Status: entity.ReservationStatusWaiting}
		query := `INSERT INTO biblioteca.reservations(user_id,book_id,position,status,expires_at) VALUES($1,$2,$3,$4,$5::timestamptz) RETURNING id,user_id,book_id,position,status,expires_at,created_at,updated_at`

		s.sqlMock.ExpectBegin()
		tx, err := s.db.Beginx()
		assert.NoError(t, err)
		s.sqlMock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(input.UserId, input.BookId, input.Position, input.Status, input.ExpiresAt).
			WillReturnError(expectedError)

		ctx, err, result := r.Create(s.ctx, tx, input)

		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, entity.Reservation{}, result)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})
}