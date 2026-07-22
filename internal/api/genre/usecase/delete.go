package usecase

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/api/genre"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type DeleteUC struct {
	db   *sqlx.DB
	repo genre.Repository
}

func NewDeleteUC(db *sqlx.DB, repo genre.Repository) DeleteUC {
	return DeleteUC{db: db, repo: repo}
}

func (u *DeleteUC) Execute(ctx context.Context, id int) (returnedCtx context.Context, err error) {
	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err
	}
	defer helpers.CloseTransaction(tx, &err)

	return u.repo.Delete(ctx, tx, id)
}
