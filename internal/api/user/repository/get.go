package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetRepo struct {
	db *sqlx.DB
}

func NewGetRepo(db *sqlx.DB) GetRepo {
	return GetRepo{db: db}
}

func (r *GetRepo) Execute(id int) (error, entity.User) {
	var user entity.User
	err := r.db.Get(&user, getSql, id)
	if err != nil {
		return err, entity.User{}
	}

	return nil, user
}

var getSql = `
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
	WHERE id = $1
`
