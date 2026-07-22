package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CreateRepo struct{}

func NewCreateRepo() CreateRepo {
	return CreateRepo{}
}

func (r *CreateRepo) Execute(ctx context.Context, tx *sqlx.Tx, input entity.Genre) (context.Context, error, entity.Genre) {
	var genre entity.Genre
	err := tx.GetContext(ctx, &genre, createSql, input.Name, input.Description)
	if err != nil {
		return ctx, err, entity.Genre{}
	}

	return ctx, nil, genre
}

var createSql = `
		INSERT INTO biblioteca.genres (name, description)
		VALUES ($1, $2)
		RETURNING id, name, description, created_at, updated_at
	`
