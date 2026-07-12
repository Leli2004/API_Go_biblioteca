package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CreateRepo struct {
	db *sqlx.DB
}

func NewCreateRepo(db *sqlx.DB) CreateRepo {
	return CreateRepo{db: db}
}

func (r *CreateRepo) Execute(input entity.Book) (error, entity.Book) {
	var book entity.Book
	err := r.db.Get(&book, createSql, input.PublisherId, input.Title, input.PublicationYear, input.Description)
	if err != nil {
		return err, entity.Book{}
	}

	return nil, book
}

var createSql = `
	INSERT INTO biblioteca.books (publisher_id, title, publication_year, description)
	VALUES ($1, $2, $3, $4)
	RETURNING id, publisher_id, title, publication_year, description, created_at, updated_at
`
