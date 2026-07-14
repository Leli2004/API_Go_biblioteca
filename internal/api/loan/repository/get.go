package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetRepo struct{}

func NewGetRepo() GetRepo {
	return GetRepo{}
}

func (r *GetRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Loan) {
	var ln entity.Loan
	err := tx.GetContext(ctx, &ln, `SELECT id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at FROM biblioteca.loans WHERE id = $1`, id)
	if err != nil {
		return ctx, err, entity.Loan{}
	}
	return ctx, nil, ln
}
