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

func (r *CreateRepo) Execute(ctx context.Context, tx *sqlx.Tx, input entity.Publisher) (context.Context, error, entity.Publisher) {
	var publisher entity.Publisher
	err := tx.GetContext(ctx, &publisher, createSql, input.Name, input.Website)
	if err != nil {
		return ctx, err, entity.Publisher{}
	}

	return ctx, nil, publisher
}

var createSql = `
	INSERT INTO biblioteca.publishers (name, website)
	VALUES ($1, $2)
	RETURNING id, name, website, created_at, updated_at
`
