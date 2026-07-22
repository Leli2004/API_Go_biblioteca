package usecase

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	userMocks "github.com/Leli2004/API_Go_biblioteca/internal/api/user/mocks"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type authUseCaseSetup struct {
	useCase      *UseCase
	userRepo     *userMocks.Repository
	sqlMock      sqlmock.Sqlmock
	db           *sqlx.DB
	ctx          context.Context
	tokenManager *security.TokenManager
}

func setupAuthUseCase(t *testing.T) authUseCaseSetup {
	sqlDB, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)

	db := sqlx.NewDb(sqlDB, "sqlmock")
	userRepo := userMocks.NewRepository(t)
	tokenManager, err := security.NewTokenManager(
		"test-secret-with-sufficient-length",
		"biblioteca-api",
		24*time.Hour,
	)
	assert.NoError(t, err)

	t.Cleanup(func() { _ = db.Close() })

	return authUseCaseSetup{
		useCase:      NewUseCase(db, userRepo, tokenManager).(*UseCase),
		userRepo:     userRepo,
		sqlMock:      sqlMock,
		db:           db,
		ctx:          context.Background(),
		tokenManager: tokenManager,
	}
}

func boolPointer(value bool) *bool {
	return &value
}

func TestLoginUseCase(t *testing.T) {
	t.Run("Happy Path - gera JWT", func(t *testing.T) {
		setup := setupAuthUseCase(t)
		passwordHash, err := security.HashPassword("senha123")
		assert.NoError(t, err)

		input := entity.LoginRequest{
			Username: " admin ",
			Password: "senha123",
		}
		foundUser := entity.User{
			Id:           1,
			Name:         "Administrador",
			Username:     "admin",
			Email:        "admin@biblioteca.com",
			PasswordHash: passwordHash,
			Role:         entity.RoleAdmin,
			Active:       boolPointer(true),
		}

		setup.sqlMock.ExpectBegin()
		setup.userRepo.On(
			"GetByUsername",
			mock.Anything,
			mock.Anything,
			"admin",
		).Return(setup.ctx, nil, foundUser).Once()
		setup.sqlMock.ExpectCommit()

		returnedCtx, err, result := setup.useCase.Login(setup.ctx, input)

		assert.NoError(t, err)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.NotEmpty(t, result.Token)
		assert.Equal(t, security.TokenTypeBearer, result.TokenType)
		assert.Equal(t, int64(86400), result.ExpiresIn)
		assert.Equal(t, foundUser.Id, result.User.Id)
		assert.Equal(t, foundUser.Username, result.User.Username)

		claims, err := setup.tokenManager.Parse(result.Token)
		assert.NoError(t, err)
		assert.Equal(t, foundUser.Id, claims.UserId)
	})

	t.Run("Error Path - input invalido", func(t *testing.T) {
		setup := setupAuthUseCase(t)

		_, err, result := setup.useCase.Login(
			setup.ctx,
			entity.LoginRequest{},
		)

		assert.Error(t, err)
		assert.Equal(t, entity.LoginResponse{}, result)
		setup.userRepo.AssertNotCalled(t, "GetByUsername")
	})

	t.Run("Error Path - username inexistente", func(t *testing.T) {
		setup := setupAuthUseCase(t)
		input := entity.LoginRequest{
			Username: "nao-existe",
			Password: "senha123",
		}

		setup.sqlMock.ExpectBegin()
		setup.userRepo.On(
			"GetByUsername",
			mock.Anything,
			mock.Anything,
			input.Username,
		).Return(setup.ctx, sql.ErrNoRows, entity.User{}).Once()
		setup.sqlMock.ExpectRollback()

		_, err, result := setup.useCase.Login(setup.ctx, input)

		assert.ErrorIs(t, err, security.ErrInvalidCredentials)
		assert.Equal(t, entity.LoginResponse{}, result)
	})

	t.Run("Error Path - senha incorreta", func(t *testing.T) {
		setup := setupAuthUseCase(t)
		passwordHash, err := security.HashPassword("senha-correta")
		assert.NoError(t, err)

		input := entity.LoginRequest{
			Username: "admin",
			Password: "senha-errada",
		}
		foundUser := entity.User{
			Id:           1,
			Username:     "admin",
			PasswordHash: passwordHash,
			Role:         entity.RoleAdmin,
			Active:       boolPointer(true),
		}

		setup.sqlMock.ExpectBegin()
		setup.userRepo.On(
			"GetByUsername",
			mock.Anything,
			mock.Anything,
			input.Username,
		).Return(setup.ctx, nil, foundUser).Once()
		setup.sqlMock.ExpectRollback()

		_, err, result := setup.useCase.Login(setup.ctx, input)

		assert.ErrorIs(t, err, security.ErrInvalidCredentials)
		assert.Equal(t, entity.LoginResponse{}, result)
	})

	t.Run("Error Path - usuario inativo", func(t *testing.T) {
		setup := setupAuthUseCase(t)
		input := entity.LoginRequest{
			Username: "admin",
			Password: "senha123",
		}
		foundUser := entity.User{
			Id:       1,
			Username: "admin",
			Role:     entity.RoleAdmin,
			Active:   boolPointer(false),
		}

		setup.sqlMock.ExpectBegin()
		setup.userRepo.On(
			"GetByUsername",
			mock.Anything,
			mock.Anything,
			input.Username,
		).Return(setup.ctx, nil, foundUser).Once()
		setup.sqlMock.ExpectRollback()

		_, err, result := setup.useCase.Login(setup.ctx, input)

		assert.ErrorIs(t, err, security.ErrInvalidCredentials)
		assert.Equal(t, entity.LoginResponse{}, result)
	})

	t.Run("Error Path - repository retorna erro", func(t *testing.T) {
		setup := setupAuthUseCase(t)
		expectedError := errors.New("repository error")
		input := entity.LoginRequest{
			Username: "admin",
			Password: "senha123",
		}

		setup.sqlMock.ExpectBegin()
		setup.userRepo.On(
			"GetByUsername",
			mock.Anything,
			mock.Anything,
			input.Username,
		).Return(setup.ctx, expectedError, entity.User{}).Once()
		setup.sqlMock.ExpectRollback()

		_, err, result := setup.useCase.Login(setup.ctx, input)

		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, entity.LoginResponse{}, result)
	})
}

func TestMeUseCase(t *testing.T) {
	
	t.Run("Happy Path - retorna usuario autenticado", func(t *testing.T) {
		setup := setupAuthUseCase(t)
		foundUser := entity.User{
			Id:       1,
			Name:     "Administrador",
			Username: "admin",
			Email:    "admin@biblioteca.com",
			Role:     entity.RoleAdmin,
			Active:   boolPointer(true),
		}

		setup.sqlMock.ExpectBegin()
		setup.userRepo.On(
			"Get",
			mock.Anything,
			mock.Anything,
			foundUser.Id,
		).Return(setup.ctx, nil, foundUser).Once()
		setup.sqlMock.ExpectCommit()

		returnedCtx, err, result := setup.useCase.Me(
			setup.ctx,
			foundUser.Id,
		)

		assert.NoError(t, err)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, foundUser.Id, result.Id)
		assert.Equal(t, foundUser.Username, result.Username)
	})

	t.Run("Error Path - ID invalido", func(t *testing.T) {
		setup := setupAuthUseCase(t)

		_, err, result := setup.useCase.Me(setup.ctx, 0)

		assert.Error(t, err)
		assert.Equal(t, entity.AuthUser{}, result)
		setup.userRepo.AssertNotCalled(t, "Get")
	})

	t.Run("Error Path - usuario inativo", func(t *testing.T) {
		setup := setupAuthUseCase(t)
		foundUser := entity.User{
			Id:       1,
			Username: "admin",
			Role:     entity.RoleAdmin,
			Active:   boolPointer(false),
		}

		setup.sqlMock.ExpectBegin()
		setup.userRepo.On(
			"Get",
			mock.Anything,
			mock.Anything,
			foundUser.Id,
		).Return(setup.ctx, nil, foundUser).Once()
		setup.sqlMock.ExpectRollback()

		_, err, result := setup.useCase.Me(setup.ctx, foundUser.Id)

		assert.Error(t, err)
		assert.Equal(t, entity.AuthUser{}, result)
	})
}
