package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetRepo struct {
	db *sqlx.DB
}

func NewGetRepo(db *sqlx.DB) GetRepo {
	return GetRepo{db: db}
}

func (r *GetRepo) Execute(id int) (error, entity.Book) {
	var book entity.Book
	err := r.db.Get(&book, getSql, id)
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
