package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type ListRepo struct {
	db *sqlx.DB
}

func NewListRepo(db *sqlx.DB) ListRepo {
	return ListRepo{
		db: db,
	}
}

func (r *ListRepo) Execute(input entity.AuthorFilters) (error, entity.AuthorList) {
	var authors []*entity.Author

	err := r.db.Select(&authors, listSql, input.Offset, input.Limit)
	if err != nil {
		return err, entity.AuthorList{}
	}

	return nil, entity.AuthorList{
		Offset: input.Offset,
		Limit:  input.Limit,
		Data:   authors,
	}
}

var listSql = `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM biblioteca.authors
		OFFSET $1
		LIMIT $2
	`
