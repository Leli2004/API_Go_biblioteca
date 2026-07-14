package repository

import (
	"context"
	"fmt"

	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UpdateRepo struct{}

func NewUpdateRepo() UpdateRepo {
	return UpdateRepo{}
}

func (r *UpdateRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int, input entity.BookCopy) (context.Context, error, entity.BookCopy) {
	var existingId int
	err := tx.GetContext(ctx, &existingId, checkBarcodeSql, input.Barcode)
	if err == nil && existingId != id {
		return ctx, fmt.Errorf("barcode '%s' já existe, não é possível duplicar", input.Barcode), entity.BookCopy{}
	}

	var copy entity.BookCopy
	err = tx.GetContext(ctx, &copy, updateSql, input.BookId, input.Barcode, input.Status, id)
	if err != nil {
		return ctx, err, entity.BookCopy{}
	}

	return ctx, nil, copy
}

var updateSql = `
	UPDATE biblioteca.book_copies
	SET book_id = $1,
		barcode = $2,
		status = $3,
		updated_at = NOW()
	WHERE id = $4
	RETURNING id, book_id, barcode, status, created_at, updated_at
`
