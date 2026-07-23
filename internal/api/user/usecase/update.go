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

type UpdateUC struct {
	db   *sqlx.DB
	repo user.Repository
}

func NewUpdateUC(db *sqlx.DB, repo user.Repository) UpdateUC {
	return UpdateUC{db: db, repo: repo}
}

func (u *UpdateUC) Execute(ctx context.Context, id int, input entity.User, claims *entity.AuthClaims) (returnedCtx context.Context, err error, result entity.User) {
	if err := security.ValidateRoles(claims, entity.RoleAdmin); err != nil {
		return ctx, err, result
	}

	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	input.SetDefault()

	if input.Password != "" {
		err = security.ValidatePassword(input.Password)
		if err != nil {
			return ctx, err, entity.User{}
		}

		input.PasswordHash, err = security.HashPassword(input.Password)
		if err != nil {
			return ctx, err, entity.User{}
		}
	} else {
		var result entity.User
		returnedCtx, err, result = u.repo.Get(ctx, tx, id)
		if err != nil {
			return ctx, err, entity.User{}
		}

		input.PasswordHash = result.PasswordHash
	}

	err = input.Validate(false)
	if err != nil {
		return ctx, err, entity.User{}
	}

	ctx, err = u.checkDuplicated(ctx, tx, id, input)
	if err != nil {
		return ctx, err, entity.User{}
	}

	return u.repo.Update(ctx, tx, id, input)
}

func (u *UpdateUC) checkDuplicated(ctx context.Context, tx *sqlx.Tx, id int, input entity.User) (returnedCtx context.Context, err error) {
	var isDuplicated bool

	returnedCtx, err, isDuplicated = u.repo.UsernameExists(ctx, tx, input.Username, id)
	if err != nil {
		return returnedCtx, fmt.Errorf("UserUC.Create.UsernameExists: %w", err)
	}

	if isDuplicated {
		return returnedCtx, fmt.Errorf("username already exists")
	}

	returnedCtx, err, isDuplicated = u.repo.EmailExists(ctx, tx, input.Email, id)
	if err != nil {
		return returnedCtx, fmt.Errorf("UserUC.Create.EmailExists: %w", err)
	}

	if isDuplicated {
		return returnedCtx, fmt.Errorf("email already exists")
	}

	return
}
