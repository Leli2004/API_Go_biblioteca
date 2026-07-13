package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetRepo struct {
	db *sqlx.DB
}

func NewGetRepo(db *sqlx.DB) GetRepo {
	return GetRepo{db: db}
}

func (r *GetRepo) Execute(id int) (error, entity.Loan) {
	var ln entity.Loan
	err := r.db.Get(&ln, `SELECT id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at FROM biblioteca.loans WHERE id = $1`, id)
	if err != nil {
		return err, entity.Loan{}
	}
	return nil, ln
}
