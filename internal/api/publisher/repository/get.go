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

func (r *GetRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Publisher) {
	var publisher entity.Publisher
	err := tx.GetContext(ctx, &publisher, getSql, id)
	if err != nil {
		return ctx, err, entity.Publisher{}
	}

	return ctx, nil, publisher
}

var getSql = `
	SELECT
		id,
		name,
		website,
		created_at,
		updated_at
	FROM biblioteca.publishers
	WHERE id = $1
`
