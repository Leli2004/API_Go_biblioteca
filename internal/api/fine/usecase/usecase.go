package usecase

import (
	"context"
	"fmt"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/fine"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type FineUC struct {
	db   *sqlx.DB
	repo fine.Repository
}

func NewUseCase(db *sqlx.DB, repo fine.Repository) *FineUC { return &FineUC{db: db, repo: repo} }

func (u *FineUC) ProcessOverdueLoans(ctx context.Context) (returnedCtx context.Context, err error, count int) {
	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, 0
	}
	defer helpers.CloseTransaction(tx, &err)

	returnedCtx, err, count = u.repo.ProcessOverdueLoans(ctx, tx, helpers.DEFAULT_FINE_AMOUNT, helpers.DEFAULT_FINE_REASON)
	if err != nil {
		return ctx, fmt.Errorf("FineUC.ProcessOverdueLoans: %w", err), 0
	}

	return
}
