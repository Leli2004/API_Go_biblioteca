package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type DeleteRepo struct {
	db *sqlx.DB
}

func NewDeleteRepo(db *sqlx.DB) DeleteRepo {
	return DeleteRepo{db: db}
}

func (r *DeleteRepo) Execute(id int) (error, entity.User) {
	var user entity.User
	err := r.db.Get(&user, deleteSql, id)
	if err != nil {
		return err, entity.User{}
	}

	return nil, user
}

var deleteSql = `
	DELETE FROM biblioteca.users
	WHERE id = $1
	RETURNING id, name, email, password_hash, phone, role, active, created_at, updated_at
`
