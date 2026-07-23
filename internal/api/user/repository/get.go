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

func (r *GetRepo) Execute(ctx context.Context, tx *sqlx.Tx, id int) (context.Context, error, entity.User) {
	var user entity.User
	err := tx.GetContext(ctx, &user, getSql, id)
	if err != nil {
		return ctx, err, entity.User{}
	}

	return ctx, nil, user
}

var getSql = `
	SELECT
		id,
		name,
		email,
		username,
		password_hash,
		phone,
		role,
		active,
		created_at,
		updated_at
	FROM biblioteca.users
	WHERE id = $1
`
