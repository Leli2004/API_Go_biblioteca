package usecase

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/book"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type CreateUC struct {
	db   *sqlx.DB
	repo book.Repository
}

func NewCreateUC(db *sqlx.DB, repo book.Repository) CreateUC {
	return CreateUC{db: db, repo: repo}
}

func (u *CreateUC) Execute(ctx context.Context, input entity.Book, claims *entity.AuthClaims) (returnedCtx context.Context, err error, result entity.Book) {
	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	err = input.Validate()
	if err != nil {
		return ctx, err, entity.Book{}
	}
	return u.repo.Create(ctx, tx, input)
}
