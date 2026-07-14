package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type ListRepo struct{}

func NewListRepo() ListRepo {
	return ListRepo{}
}

func (r *ListRepo) Execute(ctx context.Context, tx *sqlx.Tx, input entity.BookFilters) (context.Context, error, entity.BookList) {
	var books []*entity.Book
	err := tx.SelectContext(ctx, &books, listSql, input.Offset, input.Limit)
	if err != nil {
		return ctx, err, entity.BookList{}
	}

	return ctx, nil, entity.BookList{
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
