package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type DeleteRepo struct {
	db *sqlx.DB
}

func NewDeleteRepo(db *sqlx.DB) DeleteRepo {
	return DeleteRepo{db: db}
}

func (r *DeleteRepo) Execute(id int) (error, entity.Book) {
	var book entity.Book
	err := r.db.Get(&book, deleteSql, id)
	if err != nil {
		return err, entity.Book{}
	}

	return nil, book
}

var deleteSql = `
	DELETE FROM biblioteca.books
	WHERE id = $1
	RETURNING id, publisher_id, title, publication_year, description, created_at, updated_at
`
