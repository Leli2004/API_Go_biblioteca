package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type ListRepo struct {
	db *sqlx.DB
}

func NewListRepo(db *sqlx.DB) ListRepo {
	return ListRepo{db: db}
}

func (r *ListRepo) Execute(input entity.LoanFilters) (error, entity.LoanList) {
	var resp entity.LoanList
	err := r.db.Select(&resp.Data, listSql, input.Limit, input.Offset)
	if err != nil {
		return err, entity.LoanList{}
	}
	resp.Offset = input.Offset
	resp.Limit = input.Limit
	return nil, resp
}

var listSql = `
    SELECT id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at
    FROM biblioteca.loans
    ORDER BY id DESC
    LIMIT $1 OFFSET $2
`
