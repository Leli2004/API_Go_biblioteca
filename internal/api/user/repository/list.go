package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type ListRepo struct{}

func NewListRepo() ListRepo {
	return ListRepo{}
}

func (r *ListRepo) Execute(ctx context.Context, tx *sqlx.Tx, input entity.UserFilters) (context.Context, error, entity.UserList) {
	var users []*entity.User
	err := tx.SelectContext(ctx, &users, listSql, input.Offset, input.Limit)
	if err != nil {
		return ctx, err, entity.UserList{}
	}

	return ctx, nil, entity.UserList{
		Offset: input.Offset,
		Limit:  helpers.GetMin(input.Limit, len(users)),
		Data:   users,
	}
}

var listSql = `
	SELECT
		id,
		name,
		email,
		password_hash,
		phone,
		role,
		active,
		created_at,
		updated_at
	FROM biblioteca.users
	ORDER BY id DESC
	OFFSET $1 LIMIT $2
`
