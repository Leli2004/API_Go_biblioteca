package usecase

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	reservationMock "github.com/Leli2004/API_Go_biblioteca/internal/api/reservation/mocks"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type useCaseSetup struct {
	useCase *ReservationUC
	repo    *reservationMock.Repository
	sqlMock sqlmock.Sqlmock
	ctx     context.Context
}

func setupUseCase(t *testing.T) useCaseSetup {
	sqlDB, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	
	db := sqlx.NewDb(sqlDB, "sqlmock")
	repo := reservationMock.NewRepository(t)

	t.Cleanup(func() { _ = db.Close() })
	return useCaseSetup{NewUseCase(db, repo), repo, sqlMock, context.Background()}
}

func Test_Create_UseCase(t *testing.T) {

	t.Run("Happy Path - Cria reserva waiting com posição automática", func(t *testing.T) {
		s := setupUseCase(t)
		input := entity.Reservation{UserId: 1, BookId: 2}
		expected := entity.Reservation{Id: 10, UserId: 1, BookId: 2, Status: entity.ReservationStatusWaiting}
		position := 3
		expected.Position = &position
		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		s.sqlMock.ExpectBegin()
		s.repo.On("GetActiveByUserAndBook", mock.Anything, mock.Anything, 1, 2).
			Return(s.ctx, sql.ErrNoRows, entity.Reservation{}).Once()
		s.repo.On("GetNextPosition", mock.Anything, mock.Anything, 2).
			Return(s.ctx, nil, position).Once()
		s.repo.On("Create", mock.Anything, mock.Anything, mock.MatchedBy(func(v entity.Reservation) bool {
			return v.UserId == 1 && v.BookId == 2 && v.Status == entity.ReservationStatusWaiting && v.Position != nil && *v.Position == position
		})).Return(s.ctx, nil, expected).Once()
		s.sqlMock.ExpectCommit()

		ctx, err, result := s.useCase.Create(s.ctx, input, claims)

		assert.NoError(t, err)
		assert.Equal(t, s.ctx, ctx)
		assert.Equal(t, expected, result)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})

	t.Run("Happy Path - Cria reserva com status manual sem posição", func(t *testing.T) {
		s := setupUseCase(t)
		position := 99
		input := entity.Reservation{UserId: 1, BookId: 2, Status: entity.ReservationStatusAvailable, Position: &position}
		expected := entity.Reservation{Id: 11, UserId: 1, BookId: 2, Status: entity.ReservationStatusAvailable}
		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		s.sqlMock.ExpectBegin()
		s.repo.On("GetActiveByUserAndBook", mock.Anything, mock.Anything, 1, 2).
			Return(s.ctx, sql.ErrNoRows, entity.Reservation{}).Once()
		s.repo.On("Create", mock.Anything, mock.Anything, mock.MatchedBy(func(v entity.Reservation) bool {
			return v.Status == entity.ReservationStatusAvailable && v.Position == nil
		})).Return(s.ctx, nil, expected).Once()
		s.sqlMock.ExpectCommit()

		_, err, result := s.useCase.Create(s.ctx, input, claims)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		s.repo.AssertNotCalled(t, "GetNextPosition", mock.Anything, mock.Anything, mock.Anything)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Dados inválidos executam rollback", func(t *testing.T) {
		s := setupUseCase(t)
		s.sqlMock.ExpectBegin()
		s.sqlMock.ExpectRollback()
		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		_, err, result := s.useCase.Create(s.ctx, entity.Reservation{}, claims)

		assert.Error(t, err)
		assert.Equal(t, entity.Reservation{}, result)
		s.repo.AssertNotCalled(t, "GetActiveByUserAndBook")
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Reserva ativa duplicada", func(t *testing.T) {
		s := setupUseCase(t)
		input := entity.Reservation{UserId: 1, BookId: 2}
		active := entity.Reservation{Id: 7, UserId: 1, BookId: 2, Status: entity.ReservationStatusWaiting}

		s.sqlMock.ExpectBegin()
		s.repo.On("GetActiveByUserAndBook", mock.Anything, mock.Anything, 1, 2).
			Return(s.ctx, nil, active).Once()
		s.sqlMock.ExpectRollback()
		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		_, err, result := s.useCase.Create(s.ctx, input, claims)

		assert.EqualError(t, err, "user already has an active reservation for this book")
		assert.Equal(t, entity.Reservation{}, result)
		s.repo.AssertNotCalled(t, "Create")
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Falha ao consultar reserva ativa", func(t *testing.T) {
		s := setupUseCase(t)
		expectedError := errors.New("database error")
		input := entity.Reservation{UserId: 1, BookId: 2}
		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		s.sqlMock.ExpectBegin()
		s.repo.On("GetActiveByUserAndBook", mock.Anything, mock.Anything, 1, 2).
			Return(s.ctx, expectedError, entity.Reservation{}).Once()
		s.sqlMock.ExpectRollback()

		_, err, _ := s.useCase.Create(s.ctx, input, claims)

		assert.ErrorIs(t, err, expectedError)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Falha ao calcular posição", func(t *testing.T) {
		s := setupUseCase(t)
		expectedError := errors.New("position error")
		input := entity.Reservation{UserId: 1, BookId: 2}
		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		s.sqlMock.ExpectBegin()
		s.repo.On("GetActiveByUserAndBook", mock.Anything, mock.Anything, 1, 2).
			Return(s.ctx, sql.ErrNoRows, entity.Reservation{}).Once()
		s.repo.On("GetNextPosition", mock.Anything, mock.Anything, 2).
			Return(s.ctx, expectedError, 0).Once()
		s.sqlMock.ExpectRollback()

		_, err, _ := s.useCase.Create(s.ctx, input, claims)

		assert.ErrorIs(t, err, expectedError)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Falha ao criar reserva", func(t *testing.T) {
		s := setupUseCase(t)
		expectedError := errors.New("create error")
		input := entity.Reservation{UserId: 1, BookId: 2}
		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		s.sqlMock.ExpectBegin()
		s.repo.On("GetActiveByUserAndBook", mock.Anything, mock.Anything, 1, 2).
			Return(s.ctx, sql.ErrNoRows, entity.Reservation{}).Once()
		s.repo.On("GetNextPosition", mock.Anything, mock.Anything, 2).
			Return(s.ctx, nil, 1).Once()
		s.repo.On("Create", mock.Anything, mock.Anything, mock.Anything).
			Return(s.ctx, expectedError, entity.Reservation{}).Once()
		s.sqlMock.ExpectRollback()

		_, err, _ := s.useCase.Create(s.ctx, input, claims)

		assert.ErrorIs(t, err, expectedError)
		assert.NoError(t, s.sqlMock.ExpectationsWereMet())
	})
}
