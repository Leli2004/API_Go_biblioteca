package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UpdateRepo struct{}

func NewUpdateRepo() UpdateRepo {
	return UpdateRepo{}
}

func (r *UpdateRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int, input entity.Loan) (context.Context, error, entity.Loan) {
	var updated entity.Loan
	err := tx.GetContext(ctx, &updated, updateAndMaybeRestoreSql, input.ReturnedAt, input.Status, id)
	if err != nil {
		return ctx, err, entity.Loan{}
	}
	return ctx, nil, updated
}

var updateAndMaybeRestoreSql = `
WITH upd AS (
	UPDATE biblioteca.loans
	SET returned_at = $1, status = $2, updated_at = NOW()
	WHERE id = $3
	RETURNING id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at
), upd_copy AS (
	UPDATE biblioteca.book_copies SET status = 'available', updated_at = NOW()
	WHERE id = (SELECT book_copy_id FROM upd) AND $2 = 'returned'
	RETURNING id
)
SELECT id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at FROM upd;
`
