package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetRepo struct {
	sublist SublistRepo
}

func NewGetRepo(sublist SublistRepo) GetRepo { return GetRepo{sublist: sublist} }

func (r *GetRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Book) {
	var book entity.Book
	err := tx.GetContext(ctx, &book, getSql, id)
	if err != nil {
		return ctx, err, entity.Book{}
	}

	ctx, err, book.Authors = r.sublist.Authors(ctx, tx, book.Id)
	if err != nil {
		return ctx, err, entity.Book{}
	}

	ctx, err, book.Genres = r.sublist.Genres(ctx, tx, book.Id)
	if err != nil {
		return ctx, err, entity.Book{}
	}

	ctx, err, book.Copies = r.sublist.Copies(ctx, tx, book.Id)
	if err != nil {
		return ctx, err, entity.Book{}
	}

	return ctx, nil, book
}

var getSql = `
	SELECT
		id,
		publisher_id,
		title,
		publication_year,
		description,
		created_at,
		updated_at
	FROM biblioteca.books
	WHERE id = $1
`
