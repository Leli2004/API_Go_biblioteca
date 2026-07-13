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

func (r *GetRepo) Execute(id int) (error, entity.BookCopy) {
	var copy entity.BookCopy
	err := r.db.Get(&copy, getSql, id)
	if err != nil {
		return err, entity.BookCopy{}
	}

	return nil, copy
}

var getSql = `
	SELECT
		id,
		book_id,
		barcode,
		status,
		created_at,
		updated_at
	FROM biblioteca.book_copies
	WHERE id = $1
`
