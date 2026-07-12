package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
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

func (r *ListRepo) Execute(input entity.GenreFilters) (error, entity.GenreList) {
	var genres []*entity.Genre

	err := r.db.Select(&genres, listSql, input.Offset, input.Limit)
	if err != nil {
		return err, entity.GenreList{}
	}

	return nil, entity.GenreList{
		Offset: input.Offset,
		Limit:  helpers.GetMin(input.Limit, len(genres)),
		Data:   genres,
	}
}

var listSql = `
		SELECT
			id,
			name,
			description,
			created_at,
			updated_at
		FROM biblioteca.genres
		OFFSET $1
		LIMIT $2
	`
