package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UpdateRepo struct {
	db *sqlx.DB
}

func NewUpdateRepo(db *sqlx.DB) UpdateRepo {
	return UpdateRepo{db: db}
}

func (r *UpdateRepo) Execute(id int, input entity.Book) (error, entity.Book) {
	var book entity.Book
	err := r.db.Get(&book, updateSql, input.PublisherId, input.Title, input.PublicationYear, input.Description, id)
	if err != nil {
		return err, entity.Book{}
	}

	return nil, book
}

var updateSql = `
	UPDATE biblioteca.books
	SET publisher_id = $1,
	    title = $2,
	    publication_year = $3,
	    description = $4,
	    updated_at = NOW()
	WHERE id = $5
	RETURNING id, publisher_id, title, publication_year, description, created_at, updated_at
`
