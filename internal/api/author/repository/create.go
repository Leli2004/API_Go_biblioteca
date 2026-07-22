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

func (r *CreateRepo) Execute(ctx context.Context, tx *sqlx.Tx, input entity.Author) (context.Context, error, entity.Author) {
	var author entity.Author
	err := tx.GetContext(ctx, &author, createSql, input.Name)
	if err != nil {
		return ctx, err, entity.Author{}
	}

	return ctx, nil, author
}

var createSql = `
		INSERT INTO biblioteca.authors (name)
		VALUES ($1)
		RETURNING id, name, created_at, updated_at
	`
