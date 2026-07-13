package repository

import (
	"fmt"

	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CreateRepo struct {
	db *sqlx.DB
}

func NewCreateRepo(db *sqlx.DB) CreateRepo {
	return CreateRepo{db: db}
}

func (r *CreateRepo) Execute(input entity.Loan) (error, entity.Loan) {
	var created entity.Loan
	err := r.db.Get(&created, createAndUpdateSql, input.UserId, input.BookCopyId, input.LoanDate, input.DueDate, input.Status)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return fmt.Errorf("book_copy is not available"), entity.Loan{}
		}
		return err, entity.Loan{}
	}
	return nil, created
}

var createAndUpdateSql = `
WITH check_copy AS (
	SELECT status FROM biblioteca.book_copies WHERE id = $2
), ins AS (
	INSERT INTO biblioteca.loans (user_id, book_copy_id, loan_date, due_date, status)
	SELECT $1, $2, $3::timestamptz, $4::timestamptz, $5
	WHERE (SELECT status FROM check_copy) = 'available'
	RETURNING id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at
), upd AS (
	UPDATE biblioteca.book_copies SET status = 'loaned', updated_at = NOW()
	WHERE id = $2 AND (SELECT status FROM check_copy) = 'available'
	RETURNING id
)
SELECT id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at FROM ins;
`
