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

func (r *DeleteRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.User) {
	var user entity.User
	err := tx.GetContext(ctx, &user, deleteSql, id)
	if err != nil {
		return ctx, err, entity.User{}
	}

	return ctx, nil, user
}

var deleteSql = `
	DELETE FROM biblioteca.users
	WHERE id = $1
	RETURNING id, name, email, password_hash, phone, role, active, created_at, updated_at
`
