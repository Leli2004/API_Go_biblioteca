package usecase

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/api/book"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type UpdateUC struct {
	db   *sqlx.DB
	repo book.Repository
}

func NewUpdateUC(db *sqlx.DB, repo book.Repository) UpdateUC {
	return UpdateUC{db: db, repo: repo}
}

func (u *UpdateUC) Execute(ctx context.Context, id int, input entity.Book) (returnedCtx context.Context, err error, result entity.Book) {
	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	err = input.Validate()
	if err != nil {
		return ctx, err, entity.Book{}
	}
	return u.repo.Update(ctx, tx, id, input)
}
