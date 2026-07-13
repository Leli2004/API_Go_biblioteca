package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type ListRepo struct {
	db *sqlx.DB
}

func NewListRepo(db *sqlx.DB) ListRepo {
	return ListRepo{db: db}
}

func (r *ListRepo) Execute(input entity.BookCopyFilters) (error, entity.BookCopyList) {
	var copies []*entity.BookCopy
	err := r.db.Select(&copies, listSql, input.Offset, input.Limit)
	if err != nil {
		return err, entity.BookCopyList{}
	}

	return nil, entity.BookCopyList{
		Offset: input.Offset,
		Limit:  helpers.GetMin(input.Limit, len(copies)),
		Data:   copies,
	}
}

var listSql = `
	SELECT
		id,
		book_id,
		barcode,
		status,
		created_at,
		updated_at
	FROM biblioteca.book_copies
	ORDER BY id DESC
	OFFSET $1 LIMIT $2
`
