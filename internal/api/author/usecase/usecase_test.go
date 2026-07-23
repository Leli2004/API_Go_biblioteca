package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	authorMock "github.com/Leli2004/API_Go_biblioteca/internal/api/author/mocks"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type useCaseSetup struct {
	uc       *AuthorUC
	repo     *authorMock.Repository
	sqlMock  sqlmock.Sqlmock
	database *sqlx.DB
	ctx      context.Context
}

func setup(t *testing.T) useCaseSetup {
	sqlDB, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)

	db := sqlx.NewDb(sqlDB, "sqlmock")
	repo := authorMock.NewRepository(t)

	t.Cleanup(func() {
		_ = db.Close()
	})

	ctx := context.Background()

	return useCaseSetup{
		uc:       NewUseCase(db, repo),
		repo:     repo,
		sqlMock:  sqlMock,
		database: db,
		ctx:      ctx,
	}
}

func Test_Get_UseCase(t *testing.T) {

	t.Run("Happy Path - Retorna autor", func(t *testing.T) {
		setup := setup(t)

		expected := entity.Author{
			Id:   1,
			Name: "Machado de Assis",
		}

		setup.sqlMock.ExpectBegin()

		setup.repo.EXPECT().Get(mock.Anything, mock.Anything, 1).Return(context.Background(), nil, expected)

		setup.sqlMock.ExpectCommit()

		returnedCtx, err, result := setup.uc.Get(setup.ctx, 1)

		assert.NoError(t, err)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, expected, result)
	})

	t.Run("Error Path - Repository retorna erro", func(t *testing.T) {
		setup := setup(t)

		expectedError := errors.New("erro ao buscar autor")

		setup.sqlMock.ExpectBegin()

		setup.repo.
			EXPECT().
			Get(mock.Anything, mock.Anything, 1).
			Return(setup.ctx, expectedError, entity.Author{}).
			Once()

		setup.sqlMock.ExpectRollback()

		returnedCtx, err, result := setup.uc.Get(setup.ctx, 1)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Author{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Erro ao abrir transação", func(t *testing.T) {
		setup := setup(t)

		expectedError := errors.New("erro ao abrir transação")

		setup.sqlMock.
			ExpectBegin().
			WillReturnError(expectedError)

		returnedCtx, err, result := setup.uc.Get(setup.ctx, 1)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Author{}, result)
		setup.repo.AssertNotCalled(t, "Get")
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}

func Test_List_UseCase(t *testing.T) {
	t.Run("Happy Path - Retorna lista de autores", func(t *testing.T) {
		setup := setup(t)

		input := entity.AuthorFilters{}
		expectedInput := input
		expectedInput.SetDefault()

		expected := entity.AuthorList{
			Offset: expectedInput.Offset,
			Limit:  expectedInput.Limit,
			Data: []*entity.Author{
				{
					Id:   1,
					Name: "Machado de Assis",
				},
				{
					Id:   2,
					Name: "Clarice Lispector",
				},
			},
		}

		setup.sqlMock.ExpectBegin()

		setup.repo.
			EXPECT().
			List(mock.Anything, mock.Anything, expectedInput).
			Return(setup.ctx, nil, expected).
			Once()

		setup.sqlMock.ExpectCommit()

		returnedCtx, err, result := setup.uc.List(setup.ctx, input)

		assert.NoError(t, err)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, expected, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Happy Path - Mantém filtros informados", func(t *testing.T) {
		setup := setup(t)

		input := entity.AuthorFilters{
			Offset: 10,
			Limit:  5,
		}

		expectedInput := input
		expectedInput.SetDefault()

		expected := entity.AuthorList{
			Offset: expectedInput.Offset,
			Limit:  expectedInput.Limit,
			Data:   []*entity.Author{},
		}

		setup.sqlMock.ExpectBegin()

		setup.repo.EXPECT().List(mock.Anything, mock.Anything, expectedInput).
			Return(setup.ctx, nil, expected).
			Once()

		setup.sqlMock.ExpectCommit()

		returnedCtx, err, result := setup.uc.List(setup.ctx, input)

		assert.NoError(t, err)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, expected, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Repository retorna erro", func(t *testing.T) {
		setup := setup(t)

		input := entity.AuthorFilters{}
		expectedInput := input
		expectedInput.SetDefault()

		expectedError := errors.New("erro ao listar autores")

		setup.sqlMock.ExpectBegin()

		setup.repo.
			EXPECT().
			List(mock.Anything, mock.Anything, expectedInput).
			Return(setup.ctx, expectedError, entity.AuthorList{}).
			Once()

		setup.sqlMock.ExpectRollback()

		returnedCtx, err, result := setup.uc.List(setup.ctx, input)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, entity.AuthorList{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Erro ao abrir transação", func(t *testing.T) {
		setup := setup(t)

		expectedError := errors.New("erro ao abrir transação")

		setup.sqlMock.
			ExpectBegin().
			WillReturnError(expectedError)

		returnedCtx, err, result := setup.uc.List(
			setup.ctx,
			entity.AuthorFilters{},
		)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.AuthorList{}, result)
		setup.repo.AssertNotCalled(t, "List")
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}

func Test_Create_UseCase(t *testing.T) {
	t.Run("Happy Path - Cria autor", func(t *testing.T) {
		setup := setup(t)

		input := entity.Author{
			Name: "Conceição Evaristo",
		}

		expected := entity.Author{
			Id:   1,
			Name: "Conceição Evaristo",
		}

		setup.sqlMock.ExpectBegin()

		setup.repo.
			EXPECT().
			Create(mock.Anything, mock.Anything, input).
			Return(setup.ctx, nil, expected).
			Once()

		setup.sqlMock.ExpectCommit()

		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		returnedCtx, err, result := setup.uc.Create(setup.ctx, input, claims)

		assert.NoError(t, err)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, expected, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Dados do autor inválidos", func(t *testing.T) {
		setup := setup(t)

		input := entity.Author{
			Name: "",
		}

		setup.sqlMock.ExpectBegin()
		setup.sqlMock.ExpectRollback()

		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		returnedCtx, err, result := setup.uc.Create(setup.ctx, input, claims)

		assert.Error(t, err)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Author{}, result)
		setup.repo.AssertNotCalled(t, "Create")
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Repository retorna erro", func(t *testing.T) {
		setup := setup(t)

		input := entity.Author{
			Name: "Conceição Evaristo",
		}

		expectedError := errors.New("erro ao criar autor")

		setup.sqlMock.ExpectBegin()

		setup.repo.
			EXPECT().
			Create(mock.Anything, mock.Anything, input).
			Return(setup.ctx, expectedError, entity.Author{}).
			Once()

		setup.sqlMock.ExpectRollback()

		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		returnedCtx, err, result := setup.uc.Create(setup.ctx, input, claims)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, entity.Author{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Erro ao abrir transação", func(t *testing.T) {
		setup := setup(t)

		input := entity.Author{
			Name: "Conceição Evaristo",
		}

		expectedError := errors.New("erro ao abrir transação")

		setup.sqlMock.
			ExpectBegin().
			WillReturnError(expectedError)

		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		returnedCtx, err, result := setup.uc.Create(setup.ctx, input, claims)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Author{}, result)
		setup.repo.AssertNotCalled(t, "Create")
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}

func Test_Update_UseCase(t *testing.T) {
	t.Run("Happy Path - Atualiza autor", func(t *testing.T) {
		setup := setup(t)

		id := 1

		input := entity.Author{
			Name: "Machado de Assis Atualizado",
		}

		expected := entity.Author{
			Id:   id,
			Name: input.Name,
		}

		setup.sqlMock.ExpectBegin()

		setup.repo.
			EXPECT().
			Update(mock.Anything, mock.Anything, id, input).
			Return(setup.ctx, nil, expected).
			Once()

		setup.sqlMock.ExpectCommit()

		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		returnedCtx, err, result := setup.uc.Update(setup.ctx, id, input, claims)

		assert.NoError(t, err)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, expected, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Dados do autor inválidos", func(t *testing.T) {
		setup := setup(t)

		id := 1

		input := entity.Author{
			Name: "",
		}

		setup.sqlMock.ExpectBegin()
		setup.sqlMock.ExpectRollback()

		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		returnedCtx, err, result := setup.uc.Update(setup.ctx, id, input, claims)

		assert.Error(t, err)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Author{}, result)
		setup.repo.AssertNotCalled(t, "Update")
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Repository retorna erro", func(t *testing.T) {
		setup := setup(t)

		id := 1

		input := entity.Author{
			Name: "Machado de Assis Atualizado",
		}

		expectedError := errors.New("erro ao atualizar autor")

		setup.sqlMock.ExpectBegin()

		setup.repo.
			EXPECT().
			Update(mock.Anything, mock.Anything, id, input).
			Return(setup.ctx, expectedError, entity.Author{}).
			Once()

		setup.sqlMock.ExpectRollback()

		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		returnedCtx, err, result := setup.uc.Update(setup.ctx, id, input, claims)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, entity.Author{}, result)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Erro ao abrir transação", func(t *testing.T) {
		setup := setup(t)

		id := 1

		input := entity.Author{
			Name: "Machado de Assis Atualizado",
		}

		expectedError := errors.New("erro ao abrir transação")

		setup.sqlMock.
			ExpectBegin().
			WillReturnError(expectedError)

		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		returnedCtx, err, result := setup.uc.Update(setup.ctx, id, input, claims)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.Equal(t, entity.Author{}, result)
		setup.repo.AssertNotCalled(t, "Update")
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}

func Test_Delete_UseCase(t *testing.T) {
	t.Run("Happy Path - Exclui autor", func(t *testing.T) {
		setup := setup(t)

		id := 1

		setup.sqlMock.ExpectBegin()

		setup.repo.
			EXPECT().
			Delete(mock.Anything, mock.Anything, id).
			Return(setup.ctx, nil).
			Once()

		setup.sqlMock.ExpectCommit()

		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		returnedCtx, err := setup.uc.Delete(setup.ctx, id, claims)

		assert.NoError(t, err)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Repository retorna erro", func(t *testing.T) {
		setup := setup(t)

		id := 1
		expectedError := errors.New("erro ao excluir autor")

		setup.sqlMock.ExpectBegin()

		setup.repo.
			EXPECT().
			Delete(mock.Anything, mock.Anything, id).
			Return(setup.ctx, expectedError).
			Once()

		setup.sqlMock.ExpectRollback()

		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		returnedCtx, err := setup.uc.Delete(setup.ctx, id, claims)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.NotNil(t, returnedCtx)
		assert.Equal(t, setup.ctx, returnedCtx)
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})

	t.Run("Error Path - Erro ao abrir transação", func(t *testing.T) {
		setup := setup(t)

		id := 1
		expectedError := errors.New("erro ao abrir transação")

		setup.sqlMock.
			ExpectBegin().
			WillReturnError(expectedError)

		claims := &entity.AuthClaims{Role: entity.RoleAdmin}

		returnedCtx, err := setup.uc.Delete(setup.ctx, id, claims)

		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, setup.ctx, returnedCtx)
		setup.repo.AssertNotCalled(t, "Delete")
		assert.NoError(t, setup.sqlMock.ExpectationsWereMet())
	})
}
