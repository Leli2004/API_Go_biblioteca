package repository

import (
	"context"
	"fmt"

	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CreateRepo struct{}

func NewCreateRepo() CreateRepo {
	return CreateRepo{}
}

func (r *CreateRepo) Execute(ctx context.Context, tx *sqlx.Tx, input entity.BookCopy) (context.Context, error, entity.BookCopy) {
	var existingId int
	err := tx.GetContext(ctx, &existingId, checkBarcodeSql, input.Barcode)
	if err == nil {
		return ctx, fmt.Errorf("barcode '%s' já existe, não é possível duplicar", input.Barcode), entity.BookCopy{}
	}

	var copy entity.BookCopy
	err = tx.GetContext(ctx, &copy, createSql, input.BookId, input.Barcode, input.Status)
	if err != nil {
		return ctx, err, entity.BookCopy{}
	}

	return ctx, nil, copy
}

var checkBarcodeSql = `
	SELECT id FROM biblioteca.book_copies WHERE barcode = $1
`

var createSql = `
	INSERT INTO biblioteca.book_copies (book_id, barcode, status)
	VALUES ($1, $2, $3)
	RETURNING id, book_id, barcode, status, created_at, updated_at
`
