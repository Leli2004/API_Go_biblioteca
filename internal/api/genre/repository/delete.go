package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type DeleteRepo struct{}

func NewDeleteRepo() DeleteRepo {
	return DeleteRepo{}
}

func (r *DeleteRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Genre) {
	var genre entity.Genre
	err := tx.GetContext(ctx, &genre, deleteSql, id)
	if err != nil {
		return ctx, err, entity.Genre{}
	}

	return ctx, nil, genre
}

var deleteSql = `
		DELETE FROM biblioteca.genres
		WHERE id = $1
		RETURNING id, name, description, created_at, updated_at
	`
