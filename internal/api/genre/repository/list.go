package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type ListRepo struct{}

func NewListRepo() ListRepo {
	return ListRepo{}
}

func (r *ListRepo) Execute(ctx context.Context, tx *sqlx.Tx, input entity.GenreFilters) (context.Context, error, entity.GenreList) {
	var genres []*entity.Genre

	err := tx.SelectContext(ctx, &genres, listSql, input.Offset, input.Limit)
	if err != nil {
		return ctx, err, entity.GenreList{}
	}

	return ctx, nil, entity.GenreList{
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
