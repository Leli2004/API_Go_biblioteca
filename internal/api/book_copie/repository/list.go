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

func (r *ListRepo) Execute(ctx context.Context, tx *sqlx.Tx, input entity.BookCopyFilters) (context.Context, error, entity.BookCopyList) {
	var copies []*entity.BookCopy
	err := tx.SelectContext(ctx, &copies, listSql, input.Offset, input.Limit)
	if err != nil {
		return ctx, err, entity.BookCopyList{}
	}

	return ctx, nil, entity.BookCopyList{
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
