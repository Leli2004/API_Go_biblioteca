package usecase

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/api/book"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type GetUC struct {
	db   *sqlx.DB
	repo book.Repository
}

func NewGetUC(db *sqlx.DB, repo book.Repository) GetUC {
	return GetUC{db: db, repo: repo}
}

func (u *GetUC) Execute(ctx context.Context, id int) (returnedCtx context.Context, err error, result entity.Book) {
	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	return u.repo.Get(ctx, tx, id)
}
