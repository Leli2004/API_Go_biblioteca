package repository

import (
	"fmt"

	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UpdateRepo struct {
	db *sqlx.DB
}

func NewUpdateRepo(db *sqlx.DB) UpdateRepo {
	return UpdateRepo{db: db}
}

func (r *UpdateRepo) Execute(id int, input entity.BookCopy) (error, entity.BookCopy) {
	var existingId int
	err := r.db.Get(&existingId, checkBarcodeSql, input.Barcode)
	if err == nil && existingId != id {
		return fmt.Errorf("barcode '%s' já existe, não é possível duplicar", input.Barcode), entity.BookCopy{}
	}

	var copy entity.BookCopy
	err = r.db.Get(&copy, updateSql, input.BookId, input.Barcode, input.Status, id)
	if err != nil {
		return err, entity.BookCopy{}
	}

	return nil, copy
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
