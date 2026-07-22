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

func (r *DeleteRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Publisher) {
	var publisher entity.Publisher
	err := tx.GetContext(ctx, &publisher, deleteSql, id)
	if err != nil {
		return ctx, err, entity.Publisher{}
	}

	return ctx, nil, publisher
}

var deleteSql = `
	DELETE FROM biblioteca.publishers
	WHERE id = $1
	RETURNING id, name, website, created_at, updated_at
`
