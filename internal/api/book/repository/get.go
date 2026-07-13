package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetRepo struct {
	db      *sqlx.DB
	sublist SublistRepo
}

func NewGetRepo(db *sqlx.DB, sublist SublistRepo) GetRepo {
	return GetRepo{db: db, sublist: sublist}
}

func (r *GetRepo) Execute(id int) (error, entity.Book) {
	var book entity.Book
	err := r.db.Get(&book, getSql, id)
	if err != nil {
		return err, entity.Book{}
	}

	err, book.Authors = r.sublist.Authors(book.Id)
	if err != nil {
		return err, entity.Book{}
	}

	err, book.Genres = r.sublist.Genres(book.Id)
	if err != nil {
		return err, entity.Book{}
	}

	err, book.Copies = r.sublist.Copies(book.Id)
	if err != nil {
		return err, entity.Book{}
	}

	return nil, book
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
