package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type DeleteRepo struct {
	db *sqlx.DB
}

func NewDeleteRepo(db *sqlx.DB) DeleteRepo {
	return DeleteRepo{db: db}
}

func (r *DeleteRepo) Execute(id int) (error, entity.Loan) {
	var ln entity.Loan
	err := r.db.Get(&ln, deleteSql, id)
	if err != nil {
		return err, entity.Loan{}
	}
	return nil, ln
}

var deleteSql = `
    DELETE FROM biblioteca.loans
    WHERE id = $1
    RETURNING id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at
`
