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

func (r *GetRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Author) {
	var author entity.Author
	err := tx.GetContext(ctx, &author, getSql, id)
	if err != nil {
		return ctx, err, entity.Author{}
	}

	return ctx, nil, author
}

var getSql = `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM biblioteca.authors
		WHERE id = $1
	`
