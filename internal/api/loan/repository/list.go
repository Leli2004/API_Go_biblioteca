package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type ListRepo struct{}

func NewListRepo() ListRepo {
	return ListRepo{}
}

func (r *ListRepo) Execute(ctx context.Context, tx *sqlx.Tx, input entity.LoanFilters) (context.Context, error, entity.LoanList) {
	var resp entity.LoanList
	err := tx.SelectContext(ctx, &resp.Data, listSql, input.Limit, input.Offset)
	if err != nil {
		return ctx, err, entity.LoanList{}
	}
	resp.Offset = input.Offset
	resp.Limit = input.Limit
	return ctx, nil, resp
}

var listSql = `
    SELECT id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at
    FROM biblioteca.loans
    ORDER BY id DESC
    LIMIT $1 OFFSET $2
`
