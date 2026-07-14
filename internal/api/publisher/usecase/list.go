package usecase

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/api/publisher"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type ListUC struct {
	db   *sqlx.DB
	repo publisher.Repository
}

func NewListUC(db *sqlx.DB, repo publisher.Repository) ListUC {
	return ListUC{db: db, repo: repo}
}

func (u *ListUC) Execute(ctx context.Context, input entity.PublisherFilters) (returnedCtx context.Context, err error, result entity.PublisherList) {
	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	input.SetDefault()
	return u.repo.List(ctx, tx, input)
}
