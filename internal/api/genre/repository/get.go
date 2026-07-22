package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetRepo struct{}

func NewGetRepo() GetRepo {
	return GetRepo{}
}

func (r *GetRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Genre) {
	var genre entity.Genre
	err := tx.GetContext(ctx, &genre, getSql, id)
	if err != nil {
		return ctx, err, entity.Genre{}
	}

	return ctx, nil, genre
}

var getSql = `
		SELECT
			id,
			name,
			description,
			created_at,
			updated_at
		FROM biblioteca.genres
		WHERE id = $1
	`
