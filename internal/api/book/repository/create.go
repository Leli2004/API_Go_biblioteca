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
	err = tx.Get(&book, createSql, input.PublisherId, input.Title, input.PublicationYear, input.Description)
	if err != nil {
		return err, entity.Book{}
	}

	err = insertBookAuthors(tx, book.Id, input.AuthorIds)
	if err != nil {
		return err, entity.Book{}
	}

	err = insertBookGenres(tx, book.Id, input.GenreIds)
	if err != nil {
		return err, entity.Book{}
	}

	err = tx.Commit()
	if err != nil {
		return err, entity.Book{}
	}

	return nil, book
}

func insertBookAuthors(tx *sqlx.Tx, bookId int, authorIds []int) error {
	for _, authorId := range authorIds {
		if authorId <= 0 {
			continue
		}
		_, err := tx.Exec(insertAuthorSql, bookId, authorId)
		if err != nil {
			return err
		}
	}
	return nil
}

func insertBookGenres(tx *sqlx.Tx, bookId int, genreIds []int) error {
	for _, genreId := range genreIds {
		if genreId <= 0 {
			continue
		}
		_, err := tx.Exec(insertGenreSql, bookId, genreId)
		if err != nil {
			return err
		}
	}
	return nil
}

var createSql = `
	INSERT INTO biblioteca.books (publisher_id, title, publication_year, description)
	VALUES ($1, $2, $3, $4)
	RETURNING id, publisher_id, title, publication_year, description, created_at, updated_at
`

var insertAuthorSql = `
	INSERT INTO biblioteca.book_authors (book_id, author_id)
	VALUES ($1, $2)
`

var insertGenreSql = `
	INSERT INTO biblioteca.book_genres (book_id, genre_id)
	VALUES ($1, $2)
`
