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

func (r *ListRepo) Execute(ctx context.Context, tx *sqlx.Tx, input entity.AuthorFilters) (context.Context, error, entity.AuthorList) {
	var authors []*entity.Author

	err := tx.SelectContext(ctx, &authors, listSql, input.Offset, input.Limit)
	if err != nil {
		return ctx, err, entity.AuthorList{}
	}

	return ctx, nil, entity.AuthorList{
		Offset: input.Offset,
		Limit:  helpers.GetMin(input.Limit, len(authors)),
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
		ORDER BY id DESC
		OFFSET $1 LIMIT $2
	`
