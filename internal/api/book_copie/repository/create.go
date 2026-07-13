package repository

import (
	"fmt"

	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CreateRepo struct {
	db *sqlx.DB
}

func NewCreateRepo(db *sqlx.DB) CreateRepo {
	return CreateRepo{db: db}
}

func (r *CreateRepo) Execute(input entity.BookCopy) (error, entity.BookCopy) {
	var existingId int
	err := r.db.Get(&existingId, checkBarcodeSql, input.Barcode)
	if err == nil {
		return fmt.Errorf("barcode '%s' já existe, não é possível duplicar", input.Barcode), entity.BookCopy{}
	}

	var copy entity.BookCopy
	err = r.db.Get(&copy, createSql, input.BookId, input.Barcode, input.Status)
	if err != nil {
		return err, entity.BookCopy{}
	}

	return nil, copy
}

var checkBarcodeSql = `
	SELECT id FROM biblioteca.book_copies WHERE barcode = $1
`

var createSql = `
	INSERT INTO biblioteca.book_copies (book_id, barcode, status)
	VALUES ($1, $2, $3)
	RETURNING id, book_id, barcode, status, created_at, updated_at
`
