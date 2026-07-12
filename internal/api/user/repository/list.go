package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type ListRepo struct {
	db *sqlx.DB
}

func NewListRepo(db *sqlx.DB) ListRepo {
	return ListRepo{db: db}
}

func (r *ListRepo) Execute(input entity.UserFilters) (error, entity.UserList) {
	var users []*entity.User
	err := r.db.Select(&users, listSql, input.Offset, input.Limit)
	if err != nil {
		return err, entity.UserList{}
	}

	return nil, entity.UserList{
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
