package usecase

import (
	"context"
	"fmt"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/user"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/jmoiron/sqlx"
)

type CreateUC struct {
	db   *sqlx.DB
	repo user.Repository
}

func NewCreateUC(db *sqlx.DB, repo user.Repository) CreateUC {
	return CreateUC{db: db, repo: repo}
}

func (u *CreateUC) Execute(ctx context.Context, input entity.User) (returnedCtx context.Context, err error, result entity.User) {
	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	input.SetDefault()

	err = security.ValidatePassword(input.Password)
	if err != nil {
		return ctx, err, entity.User{}
	}

	input.PasswordHash, err = security.HashPassword(input.Password)
	if err != nil {
		return ctx, err, entity.User{}
	}

	err = input.Validate(true)
	if err != nil {
		return ctx, err, entity.User{}
	}

	ctx, err = u.checkDuplicated(ctx, tx, input)
	if err != nil {
		return ctx, err, entity.User{}
	}

	return u.repo.Create(ctx, tx, input)
}

func (u *CreateUC) checkDuplicated(ctx context.Context, tx *sqlx.Tx, input entity.User) (returnedCtx context.Context, err error) {
	var isDuplicated bool

	returnedCtx, err, isDuplicated = u.repo.UsernameExists(ctx, tx, input.Username, 0)
	if err != nil {
		return returnedCtx, fmt.Errorf("UserUC.Create.UsernameExists: %w", err)
	}

	if isDuplicated {
		return returnedCtx, fmt.Errorf("username already exists")
	}

	returnedCtx, err, isDuplicated = u.repo.EmailExists(ctx, tx, input.Email, 0)
	if err != nil {
		return returnedCtx, fmt.Errorf("UserUC.Create.EmailExists: %w", err)
	}

	if isDuplicated {
		return returnedCtx, fmt.Errorf("email already exists")
	}

	return
}
