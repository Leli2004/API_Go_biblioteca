package usecase

import (
	"context"
	"errors"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type CreateUC struct {
	db   *sqlx.DB
	repo loan.Repository
}

func NewCreateUC(db *sqlx.DB, repo loan.Repository) CreateUC {
	return CreateUC{db: db, repo: repo}
}

func (u *CreateUC) Execute(ctx context.Context, input entity.Loan) (returnedCtx context.Context, err error, result entity.Loan) {
	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	input.SetDefault()

	if err := input.Validate(); err != nil {
		return ctx, err, entity.Loan{}
	}

	ctx, errChk, existing := u.repo.GetActiveByUserAndBookCopy(ctx, tx, input.UserId, input.BookCopyId)
	if errChk == nil && existing.Id != 0 {
		return ctx, errors.New("book_copy is already loaned"), entity.Loan{}
	}

	ctx, err, created := u.repo.Create(ctx, tx, input)
	if err != nil {
		return ctx, err, entity.Loan{}
	}

	return ctx, nil, created
}
