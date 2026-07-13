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
	tx, err := r.db.Beginx()
	if err != nil {
		return err, entity.Book{}
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var book entity.Book
	err = tx.Get(&book, updateSql, input.PublisherId, input.Title, input.PublicationYear, input.Description, id)
	if err != nil {
		return err, entity.Book{}
	}

	err = deleteBookAuthors(tx, id)
	if err != nil {
		return err, entity.Book{}
	}

	err = deleteBookGenres(tx, id)
	if err != nil {
		return err, entity.Book{}
	}

	err = insertBookAuthors(tx, id, input.AuthorIds)
	if err != nil {
		return err, entity.Book{}
	}

	err = insertBookGenres(tx, id, input.GenreIds)
	if err != nil {
		return err, entity.Book{}
	}

	err = tx.Commit()
	if err != nil {
		return err, entity.Book{}
	}

	return nil, book
}

func deleteBookAuthors(tx *sqlx.Tx, bookId int) error {
	_, err := tx.Exec(deleteAuthorsSql, bookId)
	return err
}

func deleteBookGenres(tx *sqlx.Tx, bookId int) error {
	_, err := tx.Exec(deleteGenresSql, bookId)
	return err
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

var deleteAuthorsSql = `
	DELETE FROM biblioteca.book_authors
	WHERE book_id = $1
`

var deleteGenresSql = `
	DELETE FROM biblioteca.book_genres
	WHERE book_id = $1
`
