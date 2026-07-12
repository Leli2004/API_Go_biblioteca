package repository

import (
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UpdateRepo struct {
	db *sqlx.DB
}

func NewUpdateRepo(db *sqlx.DB) UpdateRepo {
	return UpdateRepo{db: db}
}

func (r *UpdateRepo) Execute(id int, input entity.User) (error, entity.User) {
	var user entity.User
	err := r.db.Get(&user, updateSql, input.Name, input.Email, input.PasswordHash, input.Phone, input.Role, input.Active, id)
	if err != nil {
		return err, entity.User{}
	}

	return nil, user
}

var updateSql = `
	UPDATE biblioteca.users
	SET name = $1,
	    email = $2,
	    password_hash = $3,
	    phone = $4,
	    role = COALESCE(NULLIF($5, ''), role),
	    active = COALESCE($6, active),
	    updated_at = NOW()
	WHERE id = $7
	RETURNING id, name, email, password_hash, phone, role, active, created_at, updated_at
`
