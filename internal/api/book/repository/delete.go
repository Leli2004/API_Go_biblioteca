package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type DeleteRepo struct{}

func NewDeleteRepo() DeleteRepo {
	return DeleteRepo{}
}

func (r *DeleteRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Book) {
	var book entity.Book
	err := tx.GetContext(ctx, &book, deleteSql, id)
	if err != nil {
		return ctx, err, entity.Book{}
	}

	return ctx, nil, book
}

var deleteSql = `
	DELETE FROM biblioteca.books
	WHERE id = $1
	RETURNING id, publisher_id, title, publication_year, description, created_at, updated_at
`
