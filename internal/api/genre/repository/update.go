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

func (r *UpdateRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int, input entity.Genre) (context.Context, error, entity.Genre) {
	var genre entity.Genre
	err := tx.GetContext(ctx, &genre, updateSql, input.Name, input.Description, id)
	if err != nil {
		return ctx, err, entity.Genre{}
	}

	return ctx, nil, genre
}

var updateSql = `
		UPDATE biblioteca.genres
		SET name = $1,
		    description = $2,
		    updated_at = NOW()
		WHERE id = $3
		RETURNING id, name, description, created_at, updated_at
	`
