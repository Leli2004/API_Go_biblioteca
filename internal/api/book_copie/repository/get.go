package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetRepo struct{}

func NewGetRepo() GetRepo {
	return GetRepo{}
}

func (r *GetRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.BookCopy) {
	var copy entity.BookCopy
	err := tx.GetContext(ctx, &copy, getSql, id)
	if err != nil {
		return ctx, err, entity.BookCopy{}
	}

	return ctx, nil, copy
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
