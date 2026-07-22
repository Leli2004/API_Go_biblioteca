package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type DeleteRepo struct{}

func NewDeleteRepo() DeleteRepo {
	return DeleteRepo{}
}

func (r *DeleteRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Loan) {
	var ln entity.Loan
	err := tx.GetContext(ctx, &ln, deleteSql, id)
	if err != nil {
		return ctx, err, entity.Loan{}
	}
	return ctx, nil, ln
}

var deleteSql = `
    DELETE FROM biblioteca.loans
    WHERE id = $1
    RETURNING id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at
`
