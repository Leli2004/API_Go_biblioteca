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

func (r *UpdateRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int, input entity.Publisher) (context.Context, error, entity.Publisher) {
	var publisher entity.Publisher
	err := tx.GetContext(ctx, &publisher, updateSql, input.Name, input.Website, id)
	if err != nil {
		return ctx, err, entity.Publisher{}
	}

	return ctx, nil, publisher
}

var updateSql = `
	UPDATE biblioteca.publishers
	SET name = $1,
	    website = $2,
	    updated_at = NOW()
	WHERE id = $3
	RETURNING id, name, website, created_at, updated_at
`
