package usecase

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type ListUC struct {
	db   *sqlx.DB
	repo loan.Repository
}

func NewListUC(db *sqlx.DB, repo loan.Repository) ListUC {
	return ListUC{db: db, repo: repo}
}

func (u *ListUC) Execute(ctx context.Context, input entity.LoanFilters) (returnedCtx context.Context, err error, result entity.LoanList) {
	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	input.SetDefault()
	return u.repo.List(ctx, tx, input)
}
