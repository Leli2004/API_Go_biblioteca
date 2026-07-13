package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UpdateRepo struct {
	db *sqlx.DB
}

func NewUpdateRepo(db *sqlx.DB) UpdateRepo {
	return UpdateRepo{db: db}
}

func (r *UpdateRepo) Execute(id int, input entity.Loan) (error, entity.Loan) {
	var updated entity.Loan
	err := r.db.Get(&updated, updateAndMaybeRestoreSql, input.ReturnedAt, input.Status, id)
	if err != nil {
		return err, entity.Loan{}
	}
	return nil, updated
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
