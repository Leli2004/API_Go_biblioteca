package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type DeleteRepo struct{}

func NewDeleteRepo() DeleteRepo {
	return DeleteRepo{}
}

func (r *DeleteRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, int) {
	var deletedId int
	err := tx.GetContext(ctx, &deletedId, deleteSql, id)
	if err != nil {
		return ctx, err, 0
	}

	return ctx, nil, deletedId
}

var deleteSql = `
	DELETE FROM biblioteca.book_copies
	WHERE id = $1
	RETURNING id
`
