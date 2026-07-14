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

func (r *DeleteRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.Author) {
	var author entity.Author
	err := tx.GetContext(ctx, &author, deleteSql, id)
	if err != nil {
		return ctx, err, entity.Author{}
	}

	return ctx, nil, author
}

var deleteSql = `
		DELETE FROM biblioteca.authors
		WHERE id = $1
		RETURNING id, name, created_at, updated_at
	`
