package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UpdateRepo struct{}

func NewUpdateRepo() UpdateRepo {
	return UpdateRepo{}
}

func (r *UpdateRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int, input entity.Book) (context.Context, error, entity.Book) {
	var book entity.Book
	if err := tx.GetContext(ctx, &book, updateSql, input.PublisherId, input.Title, input.PublicationYear, input.Description, id); err != nil {
		return ctx, err, entity.Book{}
	}
	if err := deleteBookAuthors(ctx, tx, book.Id); err != nil {
		return ctx, err, entity.Book{}
	}
	if err := deleteBookGenres(ctx, tx, book.Id); err != nil {
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

func deleteBookAuthors(ctx context.Context, tx *sqlx.Tx, bookId int) error {
	_, err := tx.ExecContext(ctx, deleteAuthorsSql, bookId)
	return err
}
func deleteBookGenres(ctx context.Context, tx *sqlx.Tx, bookId int) error {
	_, err := tx.ExecContext(ctx, deleteGenresSql, bookId)
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
