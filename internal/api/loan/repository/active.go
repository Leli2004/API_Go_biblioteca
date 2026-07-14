package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type ActiveRepo struct{}

func NewActiveRepo() ActiveRepo {
	return ActiveRepo{}
}

func (r *ActiveRepo) Execute(ctx context.Context, tx *sqlx.Tx, userId int, bookCopyId int) (context.Context, error, entity.Loan) {
	var ln entity.Loan
	err := tx.GetContext(ctx, &ln, `SELECT id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at FROM biblioteca.loans WHERE user_id = $1 AND book_copy_id = $2 AND status = 'active' LIMIT 1`, userId, bookCopyId)
	if err != nil {
		return ctx, err, entity.Loan{}
	}
	return ctx, nil, ln
}
