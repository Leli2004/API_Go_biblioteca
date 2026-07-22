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

func (r *UpdateRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int, input entity.Author) (context.Context, error, entity.Author) {
	var author entity.Author
	err := tx.GetContext(ctx, &author, updateSql, input.Name, id)
	if err != nil {
		return ctx, err, entity.Author{}
	}

	return ctx, nil, author
}

var updateSql = `
		UPDATE biblioteca.authors
		SET name = $1,
		updated_at = NOW()
		WHERE id = $2
		RETURNING id, name, created_at, updated_at
	`
