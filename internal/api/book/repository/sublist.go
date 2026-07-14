package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type SublistRepo struct{}

func NewSublistRepo() SublistRepo { return SublistRepo{} }

func (r *SublistRepo) Authors(ctx context.Context, tx *sqlx.Tx, bookId int) (context.Context, error, []*entity.Author) {
	var authors []*entity.Author
	err := tx.SelectContext(ctx, &authors, authorsSql, bookId)
	if err != nil {
		return ctx, err, nil
	}

	return ctx, nil, authors
}

func (r *SublistRepo) Genres(ctx context.Context, tx *sqlx.Tx, bookId int) (context.Context, error, []*entity.Genre) {
	var genres []*entity.Genre
	err := tx.SelectContext(ctx, &genres, genresSql, bookId)
	if err != nil {
		return ctx, err, nil
	}

	return ctx, nil, genres
}

func (r *SublistRepo) Copies(ctx context.Context, tx *sqlx.Tx, bookId int) (context.Context, error, []*entity.BookCopy) {
	var copies []*entity.BookCopy
	err := tx.SelectContext(ctx, &copies, copiesSql, bookId)
	if err != nil {
		return ctx, err, nil
	}

	return ctx, nil, copies
}

var authorsSql = `
SELECT
	a.id,
	a.name,
	a.created_at,
	a.updated_at
FROM biblioteca.authors a
JOIN biblioteca.book_authors ba ON ba.author_id = a.id
WHERE ba.book_id = $1
ORDER BY a.id
`

var genresSql = `
SELECT
	g.id,
	g.name,
	g.description,
	g.created_at,
	g.updated_at
FROM biblioteca.genres g
JOIN biblioteca.book_genres bg ON bg.genre_id = g.id
WHERE bg.book_id = $1
ORDER BY g.id
`

var copiesSql = `
SELECT
	id,
	book_id,
	barcode,
	status,
	created_at,
	updated_at
FROM biblioteca.book_copies
WHERE book_id = $1
ORDER BY id
`
