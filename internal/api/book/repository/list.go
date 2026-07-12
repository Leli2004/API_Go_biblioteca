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

func (r *ListRepo) Execute(input entity.BookFilters) (error, entity.BookList) {
	var books []*entity.Book
	err := r.db.Select(&books, listSql, input.Offset, input.Limit)
	if err != nil {
		return err, entity.BookList{}
	}

	return nil, entity.BookList{
		Offset: input.Offset,
		Limit:  helpers.GetMin(input.Limit, len(books)),
		Data:   books,
	}
}

var listSql = `
	SELECT
		id,
		publisher_id,
		title,
		publication_year,
		description,
		created_at,
		updated_at
	FROM biblioteca.books
	ORDER BY id DESC
	OFFSET $1 LIMIT $2
`
