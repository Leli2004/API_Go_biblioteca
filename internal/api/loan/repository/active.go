package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type ActiveRepo struct {
	db *sqlx.DB
}

func NewActiveRepo(db *sqlx.DB) ActiveRepo {
	return ActiveRepo{db: db}
}

func (r *ActiveRepo) Execute(userId int, bookCopyId int) (error, entity.Loan) {
	var ln entity.Loan
	err := r.db.Get(&ln, `SELECT id, user_id, book_copy_id, loan_date, due_date, returned_at, status, created_at, updated_at FROM biblioteca.loans WHERE user_id = $1 AND book_copy_id = $2 AND status = 'active' LIMIT 1`, userId, bookCopyId)
	if err != nil {
		return err, entity.Loan{}
	}
	return nil, ln
}
