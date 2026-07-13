package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type SublistRepo struct {
	db *sqlx.DB
}

func NewSublistRepo(db *sqlx.DB) SublistRepo {
	return SublistRepo{db: db}
}

func (r *SublistRepo) Authors(bookId int) (error, []*entity.Author) {
	var authors []*entity.Author
	err := r.db.Select(&authors, authorsSql, bookId)
	if err != nil {
		return err, nil
	}

	return nil, authors
}

func (r *SublistRepo) Genres(bookId int) (error, []*entity.Genre) {
	var genres []*entity.Genre
	err := r.db.Select(&genres, genresSql, bookId)
	if err != nil {
		return err, nil
	}

	return nil, genres
}

func (r *SublistRepo) Copies(bookId int) (error, []*entity.BookCopy) {
	var copies []*entity.BookCopy
	err := r.db.Select(&copies, copiesSql, bookId)
	if err != nil {
		return err, nil
	}

	return nil, copies
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
