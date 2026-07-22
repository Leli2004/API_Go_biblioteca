package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CreateRepo struct{}

func NewCreateRepo() CreateRepo {
	return CreateRepo{}
}

func (r *CreateRepo) Execute(ctx context.Context, tx *sqlx.Tx, input entity.Book) (context.Context, error, entity.Book) {
	var book entity.Book
	if err := tx.GetContext(ctx, &book, createSql, input.PublisherId, input.Title, input.PublicationYear, input.Description); err != nil {
		return ctx, err, entity.Book{}
	}
	if err := insertBookAuthors(ctx, tx, book.Id, input.AuthorIds); err != nil {
		return ctx, err, entity.Book{}
	}
	if err := insertBookGenres(ctx, tx, book.Id, input.GenreIds); err != nil {
		return ctx, err, entity.Book{}
	}
	return ctx, nil, book
}

func insertBookAuthors(ctx context.Context, tx *sqlx.Tx, bookId int, authorIds []int) error {
	for _, authorId := range authorIds {
		if authorId > 0 {
			if _, err := tx.ExecContext(ctx, insertAuthorSql, bookId, authorId); err != nil {
				return err
			}
		}
	}
	return nil
}

func insertBookGenres(ctx context.Context, tx *sqlx.Tx, bookId int, genreIds []int) error {
	for _, genreId := range genreIds {
		if genreId > 0 {
			if _, err := tx.ExecContext(ctx, insertGenreSql, bookId, genreId); err != nil {
				return err
			}
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
