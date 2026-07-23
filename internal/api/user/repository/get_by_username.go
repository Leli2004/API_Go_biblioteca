package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetByUsernameRepo struct{}

func NewGetByUsernameRepo() GetByUsernameRepo {
	return GetByUsernameRepo{}
}

func (r *GetByUsernameRepo) Execute(ctx context.Context, tx *sqlx.Tx, username string) (context.Context, error, entity.User) {
	var user entity.User
	err := tx.GetContext(ctx, &user, getByUsernameSql, username)
	if err != nil {
		return ctx, err, entity.User{}
	}

	return ctx, nil, user
}

var getByUsernameSql = `
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
	WHERE username = $1
	LIMIT 1
`
